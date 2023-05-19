package server

import (
	"github.com/FoxFurry/scholarlabs/services/gateway/internal/httperr"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ------------------------------------VARIABLES------------------------------------

const (
	contextKey  = "context"
	uuidKey     = "uuid"
	tokenIssuer = "scholarlabs"
)

// ------------------------------------MIDDLEWARES------------------------------------

func (s *ScholarLabs) jwtMiddleware(tokenSecret string) func(*gin.Context) {
	return func(c *gin.Context) {
		token, err := c.Cookie("X-Authorization")
		if err != nil {
			s.lg.WithError(err).Error("failed to get cookie")
			httperr.Handle(c, httperr.UnauthorizedError("missing or invalid cookie"))
			c.Abort()
			return
		}

		if token == "" {
			s.lg.Info("token is empty")
			httperr.Handle(c, httperr.UnauthorizedError("missing cookie"))
			c.Abort()
			return
		}

		uuid, err := s.jwt.ValidateToken(token, tokenIssuer, []byte(tokenSecret))
		if err != nil {
			httperr.Handle(c, httperr.UnauthorizedError("could not validate JWT token"))
			c.Abort()
			return
		}

		if uuid == "" {
			httperr.Handle(c, httperr.UnauthorizedError("missing user uuid in JWT token"))
			c.Abort()
			return
		}

		c.Set(uuidKey, uuid)
		c.Next()
	}
}

func (s *ScholarLabs) setContext() func(*gin.Context) {
	return func(c *gin.Context) {

		c.Set(contextKey, uuid.New().String())
		c.Next()
	}
}

// ------------------------------------UTILITY FUNCS------------------------------------

func (s *ScholarLabs) getUUIDFromContext(c *gin.Context) (string, error) {
	userUUID := c.GetString(uuidKey)
	if userUUID == "" {
		return "", errUserUUIDMissing
	}

	return userUUID, nil
}
