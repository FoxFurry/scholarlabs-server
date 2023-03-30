package server

import (
	"net/http"

	"github.com/FoxFurry/scholarlabs/gateway/internal/httperr"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// ------------------------------------VARIABLES------------------------------------

const (
	authSchema  = "Bearer " // Space is required by auth header standard
	uuidKey     = "uuid"
	tokenIssuer = "scholarlabs_dev"
)

// ------------------------------------MIDDLEWARES------------------------------------

func (p *ScholarLabs) jwtMiddleware(tokenSecret string) func(*gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) <= len(authSchema) {
			httperr.Handle(c, httperr.New("missing or invalid JWT token", http.StatusUnauthorized))
			return
		}

		token := authHeader[len(authSchema):]
		viper.GetString("")
		uuid, err := p.jwt.ValidateToken(token, tokenIssuer, []byte(tokenSecret))
		if err != nil {
			httperr.Handle(c, httperr.WrapHttp(err, "could not validate JWT token", http.StatusUnauthorized))
			return
		}

		c.Set(uuidKey, uuid)
		c.Next()
	}
}

func (p *ScholarLabs) corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

// ------------------------------------UTILITY FUNCS------------------------------------

func (p *ScholarLabs) getUUIDFromContext(c *gin.Context) (string, error) {
	userUUID := c.GetString(uuidKey)
	if userUUID == "" {
		return "", httperr.New("user uuid missing from context", http.StatusBadRequest)
	}

	return userUUID, nil
}