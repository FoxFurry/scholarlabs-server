package server

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *ScholarLabs) GetPublicPrototypes(ctx *gin.Context) {
	prototypes, err := s.environmentService.GetPublicPrototypes(ctx, &emptypb.Empty{})
	if err != nil {
		return
	}
}
