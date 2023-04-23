package server

import (
	"github.com/gin-gonic/gin"
)

func (s *ScholarLabs) CreateCourse(ctx *gin.Context) {
	//var (
	//	err error
	//	c   models.CourseFull
	//)
	//
	//if err = ctx.BindJSON(&c); err != nil {
	//	s.lg.WithError(err).Error("failed to bind request")
	//	httperr.Handle(ctx, httperr.New("bad request", 400))
	//	return
	//}
	//
	//c.AuthorUUID, err = s.getUUIDFromContext(ctx)
	//if err = ctx.BindJSON(&c); err != nil {
	//	s.lg.WithError(err).Error("failed to get uuid from request")
	//	httperr.Handle()
	//}
}

func (s *ScholarLabs) Enroll(ctx *gin.Context) {

}

func (s *ScholarLabs) Unenroll(ctx *gin.Context) {

}

func (s *ScholarLabs) GetEnrolledCoursesForUser(ctx *gin.Context) {

}

func (s *ScholarLabs) GetAllPublicCourses(ctx *gin.Context) {

}
