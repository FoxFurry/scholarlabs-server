package server

import (
	"net/http"

	"github.com/FoxFurry/scholarlabs/services/gateway/internal/httperr"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ------------------------------------VARIABLES------------------------------------

const (
	authSchema  = "Bearer " // Space is required by auth header standard
	contextKey  = "context"
	uuidKey     = "uuid"
	tokenIssuer = "scholarlabs"
)

// ------------------------------------MIDDLEWARES------------------------------------

func (s *ScholarLabs) jwtMiddleware(tokenSecret string) func(*gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) <= len(authSchema) {
			httperr.Handle(c, httperr.New("missing or invalid JWT token", http.StatusUnauthorized))
			c.Abort()
			return
		}

		token := authHeader[len(authSchema):]

		if token == "" {
			s.lg.Info("token is empty")
			httperr.Handle(c, httperr.UnauthorizedError("missing cookie"))
			c.Abort()
			return
		}

		userUUID, err := s.jwt.ValidateToken(token, tokenIssuer, []byte(tokenSecret))
		if err != nil {
			httperr.Handle(c, httperr.UnauthorizedError("could not validate JWT token"))
			c.Abort()
			return
		}

		if userUUID == "" {
			httperr.Handle(c, httperr.UnauthorizedError("missing user uuid in JWT token"))
			c.Abort()
			return
		}

		c.Set(uuidKey, userUUID)
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
