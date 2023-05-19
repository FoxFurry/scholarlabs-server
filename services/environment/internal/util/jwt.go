package util

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

const (
	issuerKey = "iss"
	expKey    = "exp"
	uuidKey   = "uuid"
)

type JWTProvider interface {
	CreateSignedToken(UUID string, expDate int64, issuer string, key []byte) (string, error)
	ValidateToken(token, issuer string, key []byte) (string, error)
}

type jwtprovider struct{}

func NewJWT() JWTProvider {
	return &jwtprovider{}
}

func (j *jwtprovider) CreateSignedToken(UUID string, expDate int64, issuer string, key []byte) (string, error) {
	claims := jwt.MapClaims{
		uuidKey:   UUID,
		expKey:    expDate,
		issuerKey: issuer,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *jwtprovider) ValidateToken(token, issuer string, key []byte) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}

		return key, nil
	})
	if err != nil {
		return "", err
	}

	var uuidInToken string
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if claims[issuerKey] != issuer {
			return "", errors.New("issuers mismatch")
		}
		uuidInToken, ok = claims[uuidKey].(string)
		if !ok {
			return "", errors.New("missing uuid claim")
		}
	}

	return uuidInToken, nil
}
