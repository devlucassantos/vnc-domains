package user

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/domains/role"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"strings"
	"time"
)

const (
	errorMessageWhenGeneratingUserTokens = "Error generating user tokens"
)

type builder struct {
	user          *User
	invalidFields []string
}

func NewBuilder() *builder {
	return &builder{user: &User{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "The user ID is invalid")
		return instance
	}
	instance.user.id = id
	return instance
}

func (instance *builder) FirstName(firstName string) *builder {
	if len(firstName) < 3 {
		instance.invalidFields = append(instance.invalidFields, "The user name is invalid")
		return instance
	}
	instance.user.firstName = firstName
	return instance
}

func (instance *builder) LastName(lastName string) *builder {
	if len(lastName) < 3 {
		instance.invalidFields = append(instance.invalidFields, "The user's surname is invalid")
		return instance
	}
	instance.user.lastName = lastName
	return instance
}

func (instance *builder) Email(email string) *builder {
	if !utils.IsEmailValid(email) {
		instance.invalidFields = append(instance.invalidFields, "The user's email address is invalid")
		return instance
	}
	instance.user.email = strings.ToLower(email)
	return instance
}

func (instance *builder) Password(password string) *builder {
	if !utils.IsPasswordValid(password) {
		instance.invalidFields = append(instance.invalidFields, "The user's password is invalid. The password "+
			"must be between 8 and 50 characters long and contain at least one letter and one number")
		return instance
	}
	instance.user.password = password
	return instance
}

func (instance *builder) HashedPassword(hashedPassword string) *builder {
	instance.user.password = hashedPassword
	return instance
}

func (instance *builder) Roles(roles []role.Role) *builder {
	instance.user.roles = roles
	return instance
}

func (instance *builder) Tokens(sessionId uuid.UUID) *builder {
	rsaPrivateKey, err := utils.GetRsaPrivateKeyFromEnvironmentVariable("SERVER_ACCESS_TOKEN_PRIVATE_KEY")
	if err != nil {
		log.Error("Error when extracting the server access token private key: ", err.Error())
		instance.invalidFields = append(instance.invalidFields, errorMessageWhenGeneratingUserTokens)
	}

	instance.user.accessToken, err = jwt.NewWithClaims(jwt.SigningMethodRS256, newAccessTokenClaims(
		instance.user.id.String(), sessionId.String(), instance.user.roles)).SignedString(rsaPrivateKey)
	if err != nil {
		log.Error("Error building the access token: ", err.Error())
		instance.invalidFields = append(instance.invalidFields, errorMessageWhenGeneratingUserTokens)
	}

	rsaPrivateKey, err = utils.GetRsaPrivateKeyFromEnvironmentVariable("SERVER_REFRESH_TOKEN_PRIVATE_KEY")
	if err != nil {
		log.Error("Error when extracting the server refresh token private key: ", err.Error())
		instance.invalidFields = append(instance.invalidFields, errorMessageWhenGeneratingUserTokens)
	}

	instance.user.refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodRS256, newRefreshTokenClaims(
		instance.user.id.String(), sessionId.String())).SignedString(rsaPrivateKey)
	if err != nil {
		log.Error("Error building the refresh token: ", err.Error())
		instance.invalidFields = append(instance.invalidFields, errorMessageWhenGeneratingUserTokens)
	}

	return instance
}

func (instance *builder) ActivationCode(activationCode string) *builder {
	if len(activationCode) != 6 {
		instance.invalidFields = append(instance.invalidFields, "The activation code is invalid. The activation "+
			"code must be exactly 6 characters long")
		return instance
	}
	instance.user.activationCode = activationCode
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.user.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date of the user record is invalid")
		return instance
	}
	instance.user.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date of the user record is invalid")
		return instance
	}
	instance.user.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*User, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.user, nil
}
