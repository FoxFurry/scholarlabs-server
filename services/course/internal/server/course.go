package server

import (
	"context"

	"github.com/FoxFurry/scholarlabs/services/course/internal/store"
	"github.com/FoxFurry/scholarlabs/services/course/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (p *ScholarLabs) Enroll(ctx context.Context, req *proto.EnrollRequest) (*emptypb.Empty, error) {
	err := p.service.Enroll(ctx, req.UserUUID, req.CourseUUID)
	if err != nil {
		p.lg.WithError(err).WithField("req", req).Error("failed to enroll")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (p *ScholarLabs) Unenroll(ctx context.Context, req *proto.UnenrollRequest) (*emptypb.Empty, error) {
	err := p.service.Unenroll(ctx, req.UserUUID, req.CourseUUID)
	if err != nil {
		p.lg.WithError(err).WithField("req", req).Error("failed to unenroll")
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (p *ScholarLabs) GetEnrolledCoursesForUser(ctx context.Context, req *proto.GetEnrolledCoursesForUserRequest) (*proto.GetEnrolledCoursesForUserResponse, error) {
	courses, err := p.service.GetEnrolledCoursesForUser(ctx, req.UserUUID)
	if err != nil {
		p.lg.WithError(err).WithField("req", req).Error("failed to get enrolled courses for user")
		return nil, err
	}

	return &proto.GetEnrolledCoursesForUserResponse{
		Courses: coursesToCoursesShort(courses),
	}, nil
}

func (p *ScholarLabs) CreateCourse(ctx context.Context, req *proto.CreateCourseRequest) (*emptypb.Empty, error) {
	if _, err := p.service.CreateCourse(ctx, store.Course{
		AuthorUUID:  req.GetCourse().GetMetadata().AuthorUUID,
		Title:       req.GetCourse().GetMetadata().Title,
		Description: req.GetCourse().GetMetadata().Description,
		Text:        req.GetCourse().GetText(),
	}); err != nil {
		p.lg.WithError(err).WithField("req", req).Error("failed to create a course")

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (p *ScholarLabs) GetCourseInfo(ctx context.Context, req *proto.GetCourseInfoRequest) (*proto.GetCourseInfoResponse, error) {
	course, err := p.service.GetCourseInfo(ctx, req.CourseUUID)
	if err != nil {
		p.lg.WithError(err).WithField("req", req).Error("failed to enroll")

		return nil, err
	}

	return &proto.GetCourseInfoResponse{
		Course: courseToFullCourse(*course),
	}, nil
}

func (p *ScholarLabs) GetAllPublicCourses(ctx context.Context, _ *emptypb.Empty) (*proto.GetAllPublicCoursesResponse, error) {
	courses, err := p.service.GetAllPublicCourses(ctx)
	if err != nil {
		p.lg.WithError(err).Error("failed to get all public courses")

		return nil, err
	}

	return &proto.GetAllPublicCoursesResponse{
		Courses: coursesToCoursesShort(courses),
	}, nil
}

func coursesToCoursesShort(courses []store.Course) []*proto.CourseShort {
	var (
		shortCourses []*proto.CourseShort
	)

	for _, course := range courses {
		shortCourses = append(shortCourses, courseToShortCourse(course))
	}

	return shortCourses
}

func courseToShortCourse(course store.Course) *proto.CourseShort {
	var shortCourse = new(proto.CourseShort)

	shortCourse.Title = course.Title
	shortCourse.Description = course.Description
	shortCourse.UUID = course.UUID
	shortCourse.AuthorUUID = course.AuthorUUID
	shortCourse.Thumbnail = course.Thumbnail

	return shortCourse
}

func courseToFullCourse(course store.Course) *proto.CourseFull {
	var fullCourse = new(proto.CourseFull)

	fullCourse.Metadata = courseToShortCourse(course)

	fullCourse.Text = course.Text

	return fullCourse
}
