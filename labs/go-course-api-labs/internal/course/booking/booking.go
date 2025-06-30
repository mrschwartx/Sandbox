package booking

import (
	"context"
	"database/sql"
	"time"

	"example.com/internal/course/catalog"
	"example.com/pkg/helper"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	v1 "example.com/pkg/api/course/v1"
)

type Status int

func (s Status) ApiV1() v1.Status {
	switch s {
	case StatusCreated:
		return v1.Status_CREATED
	case StatusReserved:
		return v1.Status_RESERVED
	case StatusCompleted:
		return v1.Status_COMPLETED
	case StatusFailed:
		return v1.Status_FAILED
	case StatusExpired:
		return v1.Status_EXPIRED
	default:
		return v1.Status_BOOKING_UNSPECIFIED
	}
}

const (
	StatusUnknown Status = iota
	StatusCreated
	StatusReserved
	StatusCompleted
	StatusFailed
	StatusExpired
)

type builder struct {
	b *Booking
}

func (b *builder) WithCustomer(name string, email string, phone string) *builder {
	b.b.Customer = Customer{
		Name:  name,
		Email: email,
		Phone: sql.NullString{Valid: true, String: phone},
	}
	return b
}

func (b *builder) Build() *Booking {
	return b.b
}

func For(c *catalog.Course, b *catalog.Batch) *builder {
	booking := &Booking{
		ID:        uuid.New(),
		Course:    c,
		Batch:     b,
		Price:     b.Price,
		Currency:  b.Currency,
		Status:    StatusCreated,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return &builder{
		b: booking,
	}
}

type Booking struct {
	ID            uuid.UUID       `json:"id"`
	Course        *catalog.Course `json:"course"`
	Batch         *catalog.Batch  `json:"batch"`
	NumTickets    int64           `json:"num_tickets"`
	Price         float64         `json:"price"`
	Currency      string          `json:"currency"`
	Status        Status          `json:"status"`
	ReservedAt    sql.NullTime    `json:"reserved_at"`
	ExpiredAt     sql.NullTime    `json:"expired_at"`
	PaidAt        sql.NullTime    `json:"paid_at"`
	FailedAt      sql.NullTime    `json:"failed_at"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	DeletedAt     sql.NullTime    `json:"deleted_at"`
	PaymentType   sql.NullString  `json:"payment_type"`
	InvoiceNumber sql.NullString  `json:"invoice_number"`
	Version       int64           `json:"version"`
	Customer      Customer        `json:"customer"`
}

func (b *Booking) CompletePayment(ctx context.Context, paidAt time.Time) error {
	b.Status = StatusCompleted
	b.PaidAt = sql.NullTime{
		Time:  paidAt,
		Valid: true,
	}
	b.FailedAt = sql.NullTime{}
	return nil
}

func (b *Booking) FailPayment(ctx context.Context, failedAt time.Time) error {
	b.Status = StatusFailed
	b.FailedAt = sql.NullTime{
		Time:  failedAt,
		Valid: true,
	}
	b.PaidAt = sql.NullTime{}
	return nil
}

// UpdatePayment will update the payment type and reset the payment id created before
func (b *Booking) UpdatePayment(ctx context.Context, paymentType string) error {
	b.PaymentType = sql.NullString{Valid: true, String: paymentType}
	b.InvoiceNumber = sql.NullString{}
	return nil
}

const (
	bookingHoldDuration = 10 * time.Minute
)

func (b *Booking) Reserve(ctx context.Context, batch *catalog.Batch) error {
	if err := batch.Available(ctx); err != nil {
		return err
	}
	if err := batch.Reserve(ctx); err != nil {
		return err
	}
	now := time.Now()
	b.Status = StatusReserved
	b.ReservedAt = sql.NullTime{
		Time:  now,
		Valid: true,
	}
	b.ExpiredAt = sql.NullTime{
		Time:  now.Add(bookingHoldDuration),
		Valid: true,
	}
	return nil
}

func (b *Booking) Expire(ctx context.Context) error {
	if b.Status == StatusExpired {
		return ErrBookingAlreadyExpired
	}
	if b.Status == StatusCompleted || b.Status == StatusFailed {
		return ErrBookingAlreadyCompleted
	}
	b.Status = StatusExpired
	b.UpdatedAt = time.Now()
	return nil
}

func (b Booking) ApiV1() *v1.Booking {
	var course *v1.Course
	if b.Course != nil {
		course = b.Course.ApiV1()
	}
	var batch *v1.Batch
	if b.Batch != nil {
		batch = b.Batch.ApiV1()
	}

	return &v1.Booking{
		Number:     b.ID.String(),
		Course:     course.GetCourseId(),
		Batch:      batch.GetBatchId(),
		Price:      b.Price,
		Currency:   b.Currency,
		Status:     b.Status.ApiV1(),
		CreatedAt:  timestamppb.New(b.CreatedAt),
		ReservedAt: helper.ProtoFromSQLNullTime(b.ReservedAt),
		PaidAt:     helper.ProtoFromSQLNullTime(b.PaidAt),
		ExpiredAt:  helper.ProtoFromSQLNullTime(b.ExpiredAt),
		FailedAt:   helper.ProtoFromSQLNullTime(b.FailedAt),
		Customer: &v1.Customer{
			Name:        b.Customer.Name,
			Email:       b.Customer.Email,
			PhoneNumber: b.Customer.Phone.String,
		},
		Payment: &v1.Payment{
			InvoiceNumber: b.InvoiceNumber.String,
		},
	}
}

type Customer struct {
	Name  string
	Email string
	Phone sql.NullString
}
