package catalog

import (
	"context"

	"example.com/internal/course/catalog"

	v1 "example.com/pkg/api/course/v1"
)

type Service interface {
	ListCourse(ctx context.Context, req *v1.ListCoursesParam) ([]catalog.Course, string, error)
	GetCourse(ctx context.Context, req *v1.GetCourseParam) (*catalog.Course, error)
}

type Server struct {
	v1.UnimplementedCatalogServiceServer
	service Service
}

func New(s Service) *Server {
	return &Server{
		service: s,
	}
}

func (s Server) ListCourses(ctx context.Context, req *v1.ListCoursesParam) (*v1.ListCoursesResult, error) {
	courses, nextPage, err := s.service.ListCourse(ctx, req)
	if err != nil {
		return nil, err
	}

	var data []*v1.Course
	for _, c := range courses {
		data = append(data, c.ApiV1())
	}

	res := &v1.ListCoursesResult{
		Courses:       data,
		NextPageToken: nextPage,
	}

	return res, nil
}

func (s Server) GetCourse(ctx context.Context, req *v1.GetCourseParam) (*v1.Course, error) {
	course, err := s.service.GetCourse(ctx, req)
	if err != nil {
		return nil, err
	}

	return course.ApiV1(), nil
}
