package catalog

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	v1 "example.com/pkg/api/course/v1"
)

type BatchStatus int

const (
	BatchStatusDraft BatchStatus = iota
	BatchStatusPublished
	BatchStatusArchived
)

type Batch struct {
	ID             uuid.UUID    `json:"id"`
	Name           string       `json:"name"`
	MaxSeats       int32        `json:"max_seats"`
	AvailableSeats int32        `json:"available_seats"`
	Price          float64      `json:"price"`
	Currency       string       `json:"currency"`
	Status         BatchStatus  `json:"status"`
	StartDate      sql.NullTime `json:"start_date"`
	EndDate        sql.NullTime `json:"end_date"`
	Version        int64        `json:"version"`

	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

func (b Batch) ApiV1() *v1.Batch {
	var startDate, endDate *timestamppb.Timestamp
	if b.StartDate.Valid {
		startDate = timestamppb.New(b.StartDate.Time)
	}
	if b.EndDate.Valid {
		endDate = timestamppb.New(b.EndDate.Time)
	}

	return &v1.Batch{
		DisplayName: b.Name,
		Name:        string(b.ID.String()),
		BatchId:     b.ID.String(),
		Price: &v1.Price{
			Value:    b.Price,
			Currency: b.Currency,
		},
		MaxSeats:       b.MaxSeats,
		AvailableSeats: b.AvailableSeats,
		StartDate:      startDate,
		EndDate:        endDate,
	}
}

func (b *Batch) Reserve(ctx context.Context) error {
	if err := b.Available(ctx); err != nil {
		return ErrClassNotAvailableForSale
	}
	if b.AvailableSeats < 1 {
		return ErrNotEnoughSeats
	}
	if b.MaxSeats > 0 {
		b.AvailableSeats -= 1
	}
	return nil
}

func (b *Batch) Available(ctx context.Context) error {
	if b.MaxSeats <= 0 {
		return nil
	}
	if b.AvailableSeats == 0 {
		return ErrClassSoldOut
	}
	if b.EndDate.Valid && time.Now().After(b.EndDate.Time) {
		return ErrClassNotAvailableForSale
	}
	return nil
}

// Allocate increases number of available seats. Only applicable for batch with limited seats.
func (b *Batch) Allocate(ctx context.Context, numSeat int) error {
	if b.MaxSeats > 0 {
		b.AvailableSeats += int32(numSeat)
	}
	return nil
}
