package catalog

import (
	"database/sql"
	"time"

	v1 "example.com/pkg/api/concert/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ConcertStatus int

const (
	ConcertStatusDraft ConcertStatus = iota
	ConcertStatusPublished
	ConcertStatusArchived
)

type Concert struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Slug        string        `json:"slug"`
	Description string        `json:"description"`
	Status      ConcertStatus `json:"status"`
	AnnouncedAt sql.NullTime  `json:"announced_at,omitempty"`
	Batches     []Batch       `json:"batchs"`

	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at,omitempty"`
}

func (c Concert) ApiV1() *v1.Concert {
	var announcedAt *timestamppb.Timestamp
	if c.AnnouncedAt.Valid {
		announcedAt = timestamppb.New(c.AnnouncedAt.Time)
	}

	return &v1.Concert{
		Name:        c.Slug,
		ConcertId:   c.ID.String(),
		DisplayName: c.Name,
		Description: c.Description,
		AnnouncedAt: announcedAt,
		Batches:     c.batchesPkg(),
	}
}

func (c Concert) batchesPkg() []*v1.Batch {
	var bs []*v1.Batch
	for _, b := range c.Batches {
		bs = append(bs, b.ApiV1())
	}
	return bs
}
