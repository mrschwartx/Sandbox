package catalog

import (
	"context"
	"database/sql"
	"errors"
	"time"

	v1 "example.com/pkg/api/concert/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BatchStatus int

const (
	BatchStatusDraft BatchStatus = iota
	BatchStatusPublished
	BatchStatusArchived
)

type Batch struct {
	ID               uuid.UUID    `json:"id"`
	ConcertID        uuid.UUID    `json:"concert_id"`
	Name             string       `json:"name"`
	StartTime        sql.NullTime `json:"start_time,omitempty"`
	EndTime          sql.NullTime `json:"end_time,omitempty"`
	Venue            string       `json:"venue"`
	MaxCapacity      int32        `json:"max_capacity"`
	AvailableTickets int32        `json:"available_tickets,omitempty"`
	Price            float64      `json:"price"`
	Currency         string       `json:"currency"`
	Status           BatchStatus  `json:"status"`
	Version          int64        `json:"version"`

	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at,omitempty"`
}

func (b Batch) ApiV1() *v1.Batch {
	var startDate, endDate *timestamppb.Timestamp
	if b.StartTime.Valid {
		startDate = timestamppb.New(b.StartTime.Time)
	}
	if b.EndTime.Valid {
		endDate = timestamppb.New(b.EndTime.Time)
	}

	return &v1.Batch{
		DisplayName: b.Name,
		Name:        string(b.ID.String()), // TODO change with slug
		BatchId:     b.ID.String(),
		TicketPrice: &v1.TicketPrice{
			Value:    b.Price,
			Currency: b.Currency,
		},
		MaxCapacity:      b.MaxCapacity,
		AvailableTickets: b.AvailableTickets,
		StartTime:        startDate,
		EndTime:          endDate,
	}
}

var (
	ErrNotEnoughSeats           = errors.New("no seat available")
	ErrClassSoldOut             = errors.New("class is sold out")
	ErrClassNotAvailableForSale = errors.New("class is not available for sale")
)

func (b *Batch) Reserve(ctx context.Context) error {
	if err := b.Available(ctx); err != nil {
		return ErrClassNotAvailableForSale
	}
	if b.AvailableTickets < 1 {
		return ErrNotEnoughSeats
	}
	if b.MaxCapacity > 0 {
		b.AvailableTickets -= 1
	}
	return nil
}

func (b *Batch) Available(ctx context.Context) error {
	if b.MaxCapacity <= 0 {
		return nil
	}
	if b.AvailableTickets == 0 {
		return ErrClassSoldOut
	}
	if b.EndTime.Valid && time.Now().After(b.EndTime.Time) {
		return ErrClassNotAvailableForSale
	}
	return nil
}

// Allocate increases number of available seats. Only applicable for batch with limited seats.
func (b *Batch) Allocate(ctx context.Context, numSeat int) error {
	if b.MaxCapacity > 0 {
		b.AvailableTickets += int32(numSeat)
	}
	return nil
}
