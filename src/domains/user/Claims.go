package user

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/domains/role"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"strings"
	"time"
)

type Claims struct {
	Issuer         string           `json:"iss,omitempty"`
	Subject        string           `json:"sub,omitempty"`
	Audience       jwt.ClaimStrings `json:"aud,omitempty"`
	IssuedAt       jwt.NumericDate  `json:"iat,omitempty"`
	NotBefore      jwt.NumericDate  `json:"nbf,omitempty"`
	ExpirationTime jwt.NumericDate  `json:"exp,omitempty"`
	SessionId      string           `json:"session_id,omitempty"`
	Roles          jwt.ClaimStrings `json:"roles,omitempty"`
}

func (instance Claims) Valid() error {
	if instance.NotBefore.Time.Before(time.Now()) || instance.ExpirationTime.Time.After(time.Now()) {
		return errors.New("este token é inválido")
	}

	return nil
}

func newAccessTokenClaims(userId string, sessionId string, roles []role.Role) *Claims {
	var userRoles []string
	for _, userRole := range roles {
		userRoles = append(userRoles, userRole.Code())
	}

	return &Claims{
		Issuer:         "Você na Câmara API",
		Subject:        userId,
		Audience:       strings.Split(os.Getenv("SERVER_ALLOWED_HOSTS"), ","),
		IssuedAt:       *jwt.NewNumericDate(time.Now()),
		NotBefore:      *jwt.NewNumericDate(time.Now()),
		ExpirationTime: *jwt.NewNumericDate(time.Now().Add(AccessTokenTimeout)),
		SessionId:      sessionId,
		Roles:          userRoles,
	}
}

func newRefreshTokenClaims(userId string, sessionId string) *Claims {
	return &Claims{
		Issuer:         "Você na Câmara API",
		Subject:        userId,
		Audience:       strings.Split(os.Getenv("SERVER_ALLOWED_HOSTS"), ","),
		IssuedAt:       *jwt.NewNumericDate(time.Now()),
		NotBefore:      *jwt.NewNumericDate(time.Now()),
		ExpirationTime: *jwt.NewNumericDate(time.Now().Add(RefreshTokenTimeout)),
		SessionId:      sessionId,
	}
}
