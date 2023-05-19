package server

import (
	"fmt"

	course "github.com/FoxFurry/scholarlabs/services/course/proto"
	environment "github.com/FoxFurry/scholarlabs/services/environment/proto"
	user "github.com/FoxFurry/scholarlabs/services/user/proto"

	"github.com/FoxFurry/scholarlabs/services/gateway/internal/config"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	errUserUUIDMissing = fmt.Errorf("user uuid missing from context")
)

type ScholarLabs struct {
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
		gEng:               ginEngine,
		jwt:                util.NewJWT(),
		cfg:                cfg,
		userService:        userSrv,
		courseService:      courseSrv,
		environmentService: environmentSrv,
		lg:                 logger,
	}

	withAuth := p.jwtMiddleware(p.cfg.TokenSecret)

	ginEngine.Use(cors.New(cors.Config{
		AllowWebSockets:  true,
		AllowCredentials: true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept"},
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	v1 := ginEngine.Group("/v1")
	{
		userPath := v1.Group("/user")
		{
			userPath.POST("/register", p.Register)
			userPath.POST("/login", p.Login)
		}

		coursePath := v1.Group("/course")
		{
			coursePath.GET("", p.GetAllPublicCourses)
			coursePath.GET("/mycourses", withAuth, p.GetEnrolledCoursesForUser)

			coursePath.POST("/new", withAuth, p.CreateCourse)

			coursePath.POST("/enroll", withAuth, p.Enroll)
			coursePath.POST("/unenroll", withAuth, p.Unenroll)
		}

		environmentsPath := v1.Group("/env")
		{
			environmentsPath.POST("", withAuth, p.CreateEnvironment)
			environmentsPath.GET("", withAuth, p.GetEnvironmentsForUser)
		}

		terminalPath := v1.Group("/terminal")
		{
			terminalPath.GET("", p.BidirectionalTerminal)
		}

		prototypePaths := v1.Group("/prototype")
		{
			prototypePaths.GET("/", withAuth, p.GetPublicPrototypes)
		}
	}

	return &p, nil
}

func (s *ScholarLabs) Run() {
	if err := s.gEng.Run(s.cfg.Host); err != nil {
		s.lg.Fatalf("failed to run http server: %v", err)
	}
}
