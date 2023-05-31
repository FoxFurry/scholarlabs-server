package server

import (
	"github.com/FoxFurry/scholarlabs/services/environment/proto"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/httperr"
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ScholarLabs) GetPublicPrototypes(ctx *gin.Context) {
	prototypesResponse, err := s.environmentService.GetPublicPrototypes(ctx, &emptypb.Empty{})
	if err != nil {
		s.lg.WithError(err).Error("failed to get public prototypes")
		httperr.Handle(ctx, httperr.New("bad request", 400))
		return
	}

	ctx.JSON(200, protoToModelPrototypes(prototypesResponse.Prototypes))
}

func protoToModelPrototypes(prototypes []*proto.PrototypeShort) []models.Prototype {
	var modelsPrototypes []models.Prototype
	for _, prototype := range prototypes {
		modelsPrototypes = append(modelsPrototypes, models.Prototype{
			UUID:             prototype.UUID,
			Name:             prototype.Name,
			ShortDescription: prototype.ShortDescription,
		})
	}
	return modelsPrototypes
}
