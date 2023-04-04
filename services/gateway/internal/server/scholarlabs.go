package server

import (
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/config"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/scholarlabs"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/util"
	"github.com/FoxFurry/scholarlabs/services/user/proto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ScholarLabs struct {
	service     scholarlabs.Service
	gEng        *gin.Engine
	jwt         util.JWTProvider
	cfg         config.Config
	userService proto.UserClient
	lg          *logrus.Logger
}

func New(cfg config.Config, logger *logrus.Logger, userSrv proto.UserClient) (*ScholarLabs, error) {
	gin.SetMode(gin.ReleaseMode)

	ginEngine := gin.Default()

	p := ScholarLabs{
		service:     scholarlabs.New(),
		gEng:        ginEngine,
		jwt:         util.NewJWT(),
		cfg:         cfg,
		userService: userSrv,
		lg:          logger,
	}

	v1 := ginEngine.Group("/v1")
	v1.Use(p.corsMiddleware)
	{
		user := v1.Group("/user")
		{
			user.POST("/register", p.Register)
			user.POST("/login", p.Login)
		}
	}

	return &p, nil
}

func (p *ScholarLabs) Run() {
	if err := p.gEng.Run(p.cfg.Host); err != nil {
		p.lg.Fatalf("failed to run http server: %v", err)
	}
}
