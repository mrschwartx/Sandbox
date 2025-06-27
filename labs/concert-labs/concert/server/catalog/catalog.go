package catalog

import (
	"context"
	"time"

	v1 "example.com/pkg/api/concert/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	v1.UnimplementedCatalogServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) ListConcerts(ctx context.Context, req *v1.ListConcertParam) (*v1.ListConcertsResult, error) {
	// Dummy data
	concerts := []*v1.Concert{
		{
			Name:        "concerts/rock-festival-2025",
			ConcertId:   "rock-festival-2025",
			DisplayName: "Rock Festival 2025",
			Description: "Annual rock concert featuring top performers",
			AnnouncedAt: timestamppb.New(time.Now().AddDate(0, -1, 0)),
			Performers: []*v1.Performer{
				{Name: "Band A", ImageUrl: "https://example.com/a.jpg", Genres: []string{"rock", "alternative"}},
				{Name: "Band B", ImageUrl: "https://example.com/b.jpg", Genres: []string{"metal"}},
			},
			TicketPrice: &v1.TicketPrice{
				Value:    99.99,
				Currency: "USD",
			},
			Batches: []*v1.Batch{
				{
					Name:             "concerts/rock-festival-2025/shows/night-1",
					BatchId:          "night-1",
					Concert:          "concerts/rock-festival-2025",
					StartTime:        timestamppb.New(time.Now().AddDate(0, 1, 0)),
					EndTime:          timestamppb.New(time.Now().AddDate(0, 1, 0).Add(3 * time.Hour)),
					Venue:            "Stadium A",
					MaxCapacity:      10000,
					AvailableTickets: 9800,
					TicketPrice: &v1.TicketPrice{
						Value:    99.99,
						Currency: "USD",
					},
				},
			},
		},
	}

	return &v1.ListConcertsResult{
		Concerts:      concerts,
		NextPageToken: "", // Implement pagination logic if needed
	}, nil
}

func (s *Server) GetConcert(ctx context.Context, req *v1.GetConcertParam) (*v1.GetConcertResult, error) {
	// Normally, you'd fetch from DB using req.Concert
	// Example: "concerts/rock-festival-2025"
	return &v1.GetConcertResult{
		Concert: &v1.Concert{
			Name:        req.Concert,
			ConcertId:   "rock-festival-2025",
			DisplayName: "Rock Festival 2025",
			Description: "Annual rock concert featuring top performers",
			AnnouncedAt: timestamppb.Now(),
			Performers: []*v1.Performer{
				{Name: "Band A", ImageUrl: "https://example.com/a.jpg", Genres: []string{"rock"}},
			},
			TicketPrice: &v1.TicketPrice{
				Value:    89.99,
				Currency: "USD",
			},
			Batches: []*v1.Batch{
				{
					Name:             req.Concert + "/shows/night-1",
					BatchId:          "night-1",
					Concert:          req.Concert,
					StartTime:        timestamppb.New(time.Now().AddDate(0, 1, 0)),
					EndTime:          timestamppb.New(time.Now().AddDate(0, 1, 0).Add(2 * time.Hour)),
					Venue:            "Arena A",
					MaxCapacity:      5000,
					AvailableTickets: 4500,
					TicketPrice: &v1.TicketPrice{
						Value:    89.99,
						Currency: "USD",
					},
				},
			},
		},
	}, nil
}
