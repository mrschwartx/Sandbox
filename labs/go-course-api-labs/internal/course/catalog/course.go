package catalog

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	v1 "example.com/pkg/api/course/v1"
)

type CourseStatus int

const (
	CourseStatusDraft CourseStatus = iota
	CourseStatusPublished
	CourseStatusArchived
)

type Course struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Slug        string       `json:"slug"`
	Description string       `json:"description"`
	PublishedAt sql.NullTime `json:"published_at"`
	Batches     []Batch      `json:"batches"`
	Status      CourseStatus `json:"status"`

	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

func (c Course) ApiV1() *v1.Course {
	var publishedAt *timestamppb.Timestamp
	if c.PublishedAt.Valid {
		publishedAt = timestamppb.New(c.PublishedAt.Time)
	}

	return &v1.Course{
		Name:        c.Slug,
		CourseId:    c.ID.String(),
		DisplayName: c.Name,
		Description: c.Description,
		PublishedAt: publishedAt,
		Batches:     c.batchesPkg(),
	}
}

func (c Course) batchesPkg() []*v1.Batch {
	var bs []*v1.Batch
	for _, b := range c.Batches {
		bs = append(bs, b.ApiV1())
	}
	return bs
}
