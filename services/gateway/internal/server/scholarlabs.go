package server

import (
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/config"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/scholarlabs"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ScholarLabs struct {
	service scholarlabs.Service
	gEng    *gin.Engine
	jwt     util.JWTProvider
	cfg     config.Config
	lg      *logrus.Logger
}

func New(cfg config.Config, logger *logrus.Logger) (*ScholarLabs, error) {
	gin.SetMode(gin.ReleaseMode)

	ginEngine := gin.Default()

	p := ScholarLabs{
		service: scholarlabs.New(),
		gEng:    ginEngine,
		jwt:     util.NewJWT(),
		cfg:     cfg,
		lg:      logger,
	}

	v1 := ginEngine.Group("/v1")
	v1.Use(p.corsMiddleware)
	{

		//plan := v1.Group("/plan")
		//{
		//	plan.POST("/", p.jwtMiddleware, p.CreatePlan)
		//	plan.DELETE("/:planUUID", p.jwtMiddleware, p.DeletePlan)
		//}
	}

	return &p, nil
}

func (p *ScholarLabs) Run() {
	p.gEng.Run(p.cfg.Host)
}
