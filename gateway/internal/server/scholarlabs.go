package server

import (
	"log"

	"github.com/FoxFurry/scholarlabs/gateway/internal/config"
	"github.com/FoxFurry/scholarlabs/gateway/internal/scholarlabs"
	"github.com/FoxFurry/scholarlabs/gateway/internal/util"
	"github.com/gin-gonic/gin"
)

type ScholarLabs struct {
	service scholarlabs.Service
	gEng    *gin.Engine
	jwt     util.JWTProvider
	cfg     config.Config
}

func New(cfg config.Config) (*ScholarLabs, error) {
	gin.SetMode(gin.ReleaseMode)

	ginEngine := gin.Default()

	p := ScholarLabs{
		service: scholarlabs.New(),
		gEng:    ginEngine,
		jwt:     util.NewJWT(),
		cfg:     cfg,
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
	log.Printf("Serving gateway on [%s]", p.cfg.GatewayHost)
	p.gEng.Run(p.cfg.GatewayHost)
}
