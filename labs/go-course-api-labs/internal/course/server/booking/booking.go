package booking

import (
	"context"

	"example.com/internal/course/booking"

	v1 "example.com/pkg/api/course/v1"
)

type Service interface {
	CreateBooking(ctx context.Context, req *v1.CreateBookingParam) (*booking.Booking, error)
	ReserveBooking(ctx context.Context, req *v1.ReserveBookingParam) (*booking.Booking, error)
	GetBooking(ctx context.Context, req *v1.GetBookingParam) (*booking.Booking, error)
	ExpireBooking(ctx context.Context, req *v1.ExpireBookingParam) error
	ListBookings(ctx context.Context, req *v1.ListBookingsParam) ([]booking.Booking, string, error)
}

type Server struct {
	v1.UnimplementedBookingServiceServer
	service Service
}

func (s Server) CreateBooking(ctx context.Context, req *v1.CreateBookingParam) (*v1.Booking, error) {
	b, err := s.service.CreateBooking(ctx, req)
	if err != nil {
		return nil, err
	}

	return b.ApiV1(), nil
}

func (s Server) ReserveBooking(ctx context.Context, req *v1.ReserveBookingParam) (*v1.ReserveBookingResult, error) {
	_, err := s.service.ReserveBooking(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.ReserveBookingResult{}, nil
}

func (s Server) GetBooking(ctx context.Context, req *v1.GetBookingParam) (*v1.Booking, error) {
	b, err := s.service.GetBooking(ctx, req)
	if err != nil {
		return nil, err
	}

	return b.ApiV1(), nil
}

func (s Server) ExpireBooking(ctx context.Context, req *v1.ExpireBookingParam) (*v1.ExpireBookingResult, error) {
	err := s.service.ExpireBooking(ctx, req)
	if err != nil {
		return nil, err
	}

	return &v1.ExpireBookingResult{}, nil
}

func (s Server) ListBookings(ctx context.Context, req *v1.ListBookingsParam) (*v1.ListBookingsResult, error) {
	bookings, _, err := s.service.ListBookings(ctx, req)
	if err != nil {
		return nil, err
	}

	var bks []*v1.Booking
	for _, b := range bookings {
		bks = append(bks, b.ApiV1())
	}

	return &v1.ListBookingsResult{
		Bookings: bks,
	}, nil
}
