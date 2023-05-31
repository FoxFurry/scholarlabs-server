package server

import (
	"net/http"

	"github.com/FoxFurry/scholarlabs/services/course/proto"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/httperr"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ScholarLabs) CreateCourse(ctx *gin.Context) {
	var (
		err error
		c   models.Course
	)

	if err = ctx.BindJSON(&c); err != nil {
		s.lg.WithError(err).Error("failed to bind request")
		httperr.Handle(ctx, httperr.New("bad request", 400))
		return
	}

	c.AuthorUUID, err = s.getUUIDFromContext(ctx)
	if err != nil {
		s.lg.WithError(err).Error("failed to get uuid from request")
		httperr.Handle(ctx, httperr.New("user uuid missing from the request", http.StatusUnauthorized))
		return
	}

	response, err := s.courseService.CreateCourse(ctx, &proto.CreateCourseRequest{
		AuthorUUID:       c.AuthorUUID,
		Title:            c.Title,
		ShortDescription: c.ShortDescription,
		Description:      c.Description,
		Thumbnail:        []byte(c.Thumbnail),
		Background:       []byte(c.Background),
	})
	if err != nil {
		s.lg.WithError(err).Error("failed to create course")
		httperr.Handle(ctx, httperr.New("abobus", http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"uuid": response.UUID,
	})
}

func (s *ScholarLabs) Enroll(ctx *gin.Context) {
	courseUUID := ctx.Param("courseUUID")
	if courseUUID == "" {
		s.lg.Warning("empty course uuid in the request")
		httperr.Handle(ctx, httperr.New("missing course uuid", http.StatusBadRequest))
		return
	}

	userUUID, err := s.getUUIDFromContext(ctx)
	if err != nil {
		s.lg.WithError(err).Error("failed to get uuid from request")
		httperr.Handle(ctx, httperr.New("user uuid missing from the request", http.StatusUnauthorized))
		return
	}

	_, err = s.courseService.Enroll(ctx, &proto.EnrollRequest{
		CourseUUID: courseUUID,
		UserUUID:   userUUID,
	})
	if err != nil {
		s.lg.WithError(err).Error("failed to enroll user")
		httperr.Handle(ctx, httperr.New("something went wrong", http.StatusInternalServerError))
		return
	}

	ctx.Status(http.StatusOK)
}

func (s *ScholarLabs) Unenroll(ctx *gin.Context) {
	courseUUID := ctx.Param("courseUUID")
	if courseUUID == "" {
		s.lg.Warning("empty course uuid in the request")
		httperr.Handle(ctx, httperr.New("missing course uuid", http.StatusBadRequest))
		return
	}

	userUUID, err := s.getUUIDFromContext(ctx)
	if err != nil {
		s.lg.WithError(err).Error("failed to get uuid from request")
		httperr.Handle(ctx, httperr.New("user uuid missing from the request", http.StatusUnauthorized))
		return
	}

	_, err = s.courseService.Unenroll(ctx, &proto.UnenrollRequest{
		CourseUUID: courseUUID,
		UserUUID:   userUUID,
	})
	if err != nil {
		s.lg.WithError(err).Error("failed to unenroll user")
		httperr.Handle(ctx, httperr.New("something went wrong", http.StatusInternalServerError))
		return
	}

	ctx.Status(http.StatusOK)
}

func (s *ScholarLabs) GetEnrolledCoursesForUser(ctx *gin.Context) {
	userUUID, err := s.getUUIDFromContext(ctx)
	if err != nil {
		s.lg.WithError(err).Error("failed to get uuid from request")
		httperr.Handle(ctx, httperr.New("user uuid missing from the request", http.StatusUnauthorized))
		return
	}

	courses, err := s.courseService.GetEnrolledCoursesForUser(ctx, &proto.GetEnrolledCoursesForUserRequest{
		UserUUID: userUUID,
	})
	if err != nil {
		s.lg.WithError(err).Error("failed to get enrolled courses")
		httperr.Handle(ctx, httperr.New("something went wrong", http.StatusInternalServerError))
		return
	}

	ctx.JSON(200, protoToModelCoursesMetadata(courses.Courses))
}

func (s *ScholarLabs) GetAllPublicCourses(ctx *gin.Context) {
	courses, err := s.courseService.GetAllPublicCourses(ctx, &emptypb.Empty{})
	if err != nil {
		s.lg.WithError(err).Error("failed to get public courses")
		httperr.Handle(ctx, httperr.New("something went wrong", http.StatusInternalServerError))
		return
	}

	ctx.JSON(200, protoToModelCoursesMetadata(courses.Courses))
}

func (s *ScholarLabs) GetCourseToC(ctx *gin.Context) {
	courseUUID := ctx.Param("courseUUID")
	if courseUUID == "" {
		s.lg.Warning("empty course uuid in the request")
		httperr.Handle(ctx, httperr.New("missing course uuid", http.StatusBadRequest))
		return
	}

	toc, err := s.courseService.GetCourseToC(ctx, &proto.GetCourseToCRequest{
		CourseUUID: courseUUID,
	})
	if err != nil {
		s.lg.WithError(err).Error("failed to get toc")
		httperr.Handle(ctx, httperr.New("something went wrong", http.StatusInternalServerError))
		return
	}

	ctx.JSON(200, protoToModelToC(toc.Pages))
}

func (s *ScholarLabs) GetCourseDescription(ctx *gin.Context) {
	courseUUID := ctx.Param("courseUUID")
	if courseUUID == "" {
		s.lg.Warning("empty course uuid in the request")
		httperr.Handle(ctx, httperr.New("missing course uuid", http.StatusBadRequest))
		return
	}

	description, err := s.courseService.GetCourseSummary(ctx, &proto.GetCourseSummaryRequest{
		CourseUUID: courseUUID,
	})
	if err != nil {
		s.lg.WithError(err).Error("failed to get course description")
		httperr.Handle(ctx, httperr.New("something went wrong", http.StatusInternalServerError))
		return
	}

	ctx.JSON(200, protoToModelCourse(description.Course))
}

func protoToModelToC(protoToC []*proto.PageMetadata) []models.PageMetadata {
	var modelToC []models.PageMetadata

	for _, protoMetadata := range protoToC {
		modelToC = append(modelToC, models.PageMetadata{
			ID:    protoMetadata.GetID(),
			Title: protoMetadata.GetTitle(),
			Type:  protoMetadata.GetType().String(),
		})
	}

	return modelToC
}

func protoToModelCoursesMetadata(protoCourses []*proto.CourseMetadata) []models.Course {
	var modelCoursesMD []models.Course

	for _, protoCourseMD := range protoCourses {
		modelCoursesMD = append(modelCoursesMD, protoToModelCourseMetadata(protoCourseMD))
	}

	return modelCoursesMD
}

func protoToModelCourseMetadata(protoCourse *proto.CourseMetadata) models.Course {
	return models.Course{
		UUID:             protoCourse.UUID,
		AuthorUUID:       protoCourse.AuthorUUID,
		Title:            protoCourse.Title,
		ShortDescription: protoCourse.ShortDescription,
	}
}

func protoToModelCourse(protoCourse *proto.Course) models.Course {
	return models.Course{
		UUID:             protoCourse.Metadata.UUID,
		AuthorUUID:       protoCourse.Metadata.AuthorUUID,
		Title:            protoCourse.Metadata.Title,
		ShortDescription: protoCourse.Metadata.ShortDescription,
		Thumbnail:        protoCourse.Metadata.Thumbnail,
		Description:      protoCourse.Description,
	}
}
