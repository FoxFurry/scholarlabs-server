package server

import (
	"context"

	"github.com/FoxFurry/scholarlabs/services/course/internal/store"
	"github.com/FoxFurry/scholarlabs/services/course/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

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

func (p *ScholarLabs) CreateCourse(ctx context.Context, req *proto.CreateCourseRequest) (*proto.CreateCourseResponse, error) {
	uuid, err := p.service.CreateCourse(ctx, store.Course{
		AuthorUUID:  req.GetAuthorUUID(),
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
	},
		req.GetThumbnail(),
		req.GetBackground(),
	)
	if err != nil {
		p.lg.WithError(err).WithField("req", req).Error("failed to create a course")
		return nil, err
	}

	return &proto.CreateCourseResponse{
		UUID: uuid,
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

func (p *ScholarLabs) GetCourseSummary(ctx context.Context, req *proto.GetCourseSummaryRequest) (*proto.GetCourseSummaryResponse, error) {
	course, err := p.service.GetCourseInfo(ctx, req.GetCourseUUID())
	if err != nil {
		p.lg.WithError(err).Error("failed to get course summary")
		return nil, err
	}

	return &proto.GetCourseSummaryResponse{
		Course: courseToFullCourse(*course),
	}, nil
}

func coursesToCoursesShort(courses []store.Course) []*proto.CourseMetadata {
	var (
		shortCourses []*proto.CourseMetadata
	)

	for _, course := range courses {
		shortCourses = append(shortCourses, courseToShortCourse(course))
	}

	return shortCourses
}

func courseToShortCourse(course store.Course) *proto.CourseMetadata {
	var shortCourse = new(proto.CourseMetadata)

	shortCourse.Title = course.Title
	shortCourse.ShortDescription = course.ShortDescription
	shortCourse.UUID = course.UUID
	shortCourse.AuthorUUID = course.AuthorUUID
	shortCourse.Thumbnail = course.Thumbnail

	return shortCourse
}

func courseToFullCourse(course store.Course) *proto.Course {
	var fullCourse = new(proto.Course)

	fullCourse.Metadata = courseToShortCourse(course)
	fullCourse.Background = course.Background
	fullCourse.Description = course.Description

	return fullCourse
}
