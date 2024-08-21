package utils

import (
	"crypto/rsa"
	"encoding/base64"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/gommon/log"
	"os"
)

func GetRsaPrivateKeyFromEnvironmentVariable(environmentVariable string) (*rsa.PrivateKey, error) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(os.Getenv(environmentVariable))
	if err != nil {
		log.Error("Erro durante a decodificação da chave privada: ", err.Error())
		return nil, err
	}

	rsaPrivateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		log.Error("Erro durante a construção da chave privada: ", err.Error())
		return nil, err
	}

	return rsaPrivateKey, nil
}
