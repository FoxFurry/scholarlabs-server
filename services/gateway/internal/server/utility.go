package server

import (
	"net/http"

	"github.com/FoxFurry/scholarlabs/services/gateway/internal/httperr"
	"github.com/gin-gonic/gin"
)

// ------------------------------------VARIABLES------------------------------------

const (
	authSchema  = "Bearer " // Space is required by auth header standard
	uuidKey     = "uuid"
	tokenIssuer = "scholarlabs"
)

// ------------------------------------MIDDLEWARES------------------------------------

func (s *ScholarLabs) jwtMiddleware(tokenSecret string) func(*gin.Context) {
	return func(c *gin.Context) {
		token, err := c.Cookie("X-Authorization")
		if err != nil {
			s.lg.WithError(err).Error("failed to get cookie")
			httperr.Handle(c, httperr.WrapHttp(err, "missing or invalid cookie", http.StatusUnauthorized))
			c.Abort()
			return
		}
		s.lg.Info("token: ", token)

		if token == "" {
			s.lg.Info("token is empty")
			httperr.Handle(c, httperr.New("missing or invalid cookie", http.StatusUnauthorized))
			c.Abort()
			return
		}

		uuid, err := s.jwt.ValidateToken(token, tokenIssuer, []byte(tokenSecret))
		if err != nil {
			httperr.Handle(c, httperr.WrapHttp(err, "could not validate JWT token", http.StatusUnauthorized))
			c.Abort()
			return
		}

		if uuid == "" {
			httperr.Handle(c, httperr.New("missing user uuid in JWT token", http.StatusUnauthorized))
			c.Abort()
			return
		}

		c.Set(uuidKey, uuid)
		c.Next()
	}
}

// ------------------------------------UTILITY FUNCS------------------------------------

func (s *ScholarLabs) getUUIDFromContext(c *gin.Context) (string, error) {
	userUUID := c.GetString(uuidKey)
	if userUUID == "" {
		return "", httperr.New("user uuid missing from context", http.StatusBadRequest)
	}

	return userUUID, nil
}
