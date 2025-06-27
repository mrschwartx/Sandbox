package catalog

import (
	"context"
	"database/sql"
	"math/rand"
	"time"

	v1 "example.com/pkg/api/concert/v1"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db    *sqlx.DB
	store *Store
}

func NewService(store *Store, db *sqlx.DB) *Service {
	return &Service{
		db:    db,
		store: store,
	}
}

func (s Service) ListConcert(ctx context.Context, req *v1.ListConcertParam) ([]Concert, string, error) {
	opts := []ListOption{
		WithMaxResults(req.GetPageSize()),
		WithNextPage(req.GetPageToken()),
	}
	for _, f := range req.GetListMask().GetPaths() {
		switch f {
		case "concerts.batches":
			opts = append(opts, WithPreload())
		}
	}
	return s.store.FindAllConcert(ctx, opts...)
}

func (s Service) GetCourse(ctx context.Context, req *v1.GetConcertParam) (*Concert, error) {
	return s.store.FindConcertByID(ctx, req.GetConcert())
}

func (s Service) Seed(ctx context.Context) error {
	for i := 0; i < 1000; i++ {
		start := time.Now()
		end := start.AddDate(0, 2, 0)

		// generate random number of batches
		numBatches := rand.Intn(5) + 1
		var batches []Batch
		for j := 0; j < numBatches; j++ {
			tc := Batch{
				ID:               uuid.New(),
				Name:             faker.Name(),
				StartTime:        sql.NullTime{Time: start, Valid: true},
				EndTime:          sql.NullTime{Time: end, Valid: true},
				Venue:            faker.Name() + " Veneu",
				MaxCapacity:      1000,
				AvailableTickets: int32(rand.Intn(50)),
				Price:            float64(rand.Intn(100000)) + 100000,
				Currency:         "IDR",
				Status:           BatchStatusPublished,
				CreatedAt:        time.Now(),
				UpdatedAt:        time.Now(),
			}
			batches = append(batches, tc)
		}

		c := &Concert{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Name:        faker.Name(),
			Slug:        faker.Username(),
			Description: faker.Paragraph(),
			AnnouncedAt: sql.NullTime{
				Time:  start,
				Valid: true,
			},
			Status:  ConcertStatusPublished,
			Batches: batches,
		}

		err := s.store.CreateConcert(ctx, c)
		if err != nil {
			return err
		}
	}
	return nil
}
