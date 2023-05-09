package server

import (
	course "github.com/FoxFurry/scholarlabs/services/course/proto"
	environment "github.com/FoxFurry/scholarlabs/services/environment/proto"
	user "github.com/FoxFurry/scholarlabs/services/user/proto"

	"github.com/FoxFurry/scholarlabs/services/gateway/internal/config"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/scholarlabs"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ScholarLabs struct {
	service            scholarlabs.Service
	gEng               *gin.Engine
	jwt                util.JWTProvider
	cfg                config.Config
	userService        user.UserClient
	courseService      course.CoursesClient
	environmentService environment.EnvironmentClient
	lg                 *logrus.Logger
}

func New(cfg config.Config, logger *logrus.Logger, userSrv user.UserClient, courseSrv course.CoursesClient, environmentSrv environment.EnvironmentClient) (*ScholarLabs, error) {
	gin.SetMode(gin.ReleaseMode)

	ginEngine := gin.Default()

	p := ScholarLabs{
		service:            scholarlabs.New(),
		gEng:               ginEngine,
		jwt:                util.NewJWT(),
		cfg:                cfg,
		userService:        userSrv,
		courseService:      courseSrv,
		environmentService: environmentSrv,
		lg:                 logger,
	}

	withAuth := p.jwtMiddleware(p.cfg.TokenSecret)

	v1 := ginEngine.Group("/v1")
	v1.Use(p.corsMiddleware)
	{
		userPath := v1.Group("/user")
		{
			userPath.POST("/register", p.Register)
			userPath.POST("/login", p.Login)
		}

		coursePath := v1.Group("/course")
		{
			coursePath.GET("/", p.GetAllPublicCourses)
			coursePath.GET("/mycourses", withAuth, p.GetEnrolledCoursesForUser)

			coursePath.POST("/new", withAuth, p.CreateCourse)

			coursePath.POST("/enroll", withAuth, p.Enroll)
			coursePath.POST("/unenroll", withAuth, p.Unenroll)
		}

		environmentsPath := v1.Group("/env")
		{
			environmentsPath.POST("/", withAuth, p.CreateEnvironment)
			environmentsPath.GET("/", withAuth, p.GetEnvironmentsForUser)
		}
	}

	return &p, nil
}

func (p *ScholarLabs) Run() {
	if err := p.gEng.Run(p.cfg.Host); err != nil {
		p.lg.Fatalf("failed to run http server: %v", err)
	}
}
