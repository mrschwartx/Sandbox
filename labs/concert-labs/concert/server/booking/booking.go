package booking

import (
	"context"
	"time"

	v1 "example.com/pkg/api/concert/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Server struct {
	v1.UnimplementedBookingServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) ListBookings(ctx context.Context, req *v1.ListBookingsParam) (*v1.ListBookingsResult, error) {
	return &v1.ListBookingsResult{
		Bookings: []*v1.Booking{
			dummyBooking("bookings/abc123"),
			dummyBooking("bookings/xyz456"),
		},
		NextPageToken: "",
	}, nil
}

func (s *Server) CreateBooking(ctx context.Context, req *v1.CreateBookingParam) (*v1.CreateBookingResult, error) {
	// In a real app, you'd persist req.Booking to the DB.
	created := req.Booking
	created.Number = "bookings/abc123"
	created.Status = v1.Status_CREATED
	created.CreatedAt = timestamppb.Now()

	return &v1.CreateBookingResult{
		Booking: created,
	}, nil
}

func (s *Server) GetBooking(ctx context.Context, req *v1.GetBookingParam) (*v1.GetBookingResult, error) {
	return &v1.GetBookingResult{
		Booking: dummyBooking(req.Booking),
	}, nil
}

func (s *Server) ReserveBooking(ctx context.Context, req *v1.ReserveBookingParam) (*v1.ReserveBookingResult, error) {
	// Update booking status to RESERVED in database (mocked here)
	return &v1.ReserveBookingResult{}, nil
}

func (s *Server) ExpireBooking(ctx context.Context, req *v1.ExpireBookingParam) (*v1.ExpireBookingResult, error) {
	// Update booking status to EXPIRED in database (mocked here)
	return &v1.ExpireBookingResult{}, nil
}

func dummyBooking(id string) *v1.Booking {
	return &v1.Booking{
		Number:    id,
		Concert:   "concerts/rockfest2025",
		Batch:     "concerts/rockfest2025/batches/first-show",
		Price:     100.00,
		Currency:  "USD",
		Status:    v1.Status_CREATED,
		CreatedAt: timestamppb.New(time.Now().Add(-1 * time.Hour)),
		Customer: &v1.Customer{
			Name:        "Alice",
			Email:       "alice@example.com",
			PhoneNumber: "+621234567890",
			ShippingAddress: &v1.Address{
				StreetAddress: "Jl. Example 123",
				City:          "Jakarta",
				Country:       "Indonesia",
			},
			BillingAddress: &v1.Address{
				StreetAddress: "Jl. Billing 456",
				City:          "Jakarta",
				Country:       "Indonesia",
			},
		},
		Payment: &v1.Payment{
			Method: "credit_card",
		},
	}
}
