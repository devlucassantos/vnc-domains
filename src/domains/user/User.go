package user

import (
	"github.com/devlucassantos/vnc-domains/src/domains/role"
	"github.com/google/uuid"
	"reflect"
	"time"
)

const (
	RefreshTokenTimeout = time.Hour * 24 * 7
	AccessTokenTimeout  = time.Minute * 15
)

type User struct {
	id             uuid.UUID
	firstName      string
	lastName       string
	email          string
	password       string
	roles          []role.Role
	accessToken    string
	refreshToken   string
	activationCode string
	active         bool
	createdAt      time.Time
	updatedAt      time.Time
}

func (instance *User) NewUpdater() *builder {
	return &builder{user: instance}
}

func (instance *User) Id() uuid.UUID {
	return instance.id
}

func (instance *User) FirstName() string {
	return instance.firstName
}

func (instance *User) LastName() string {
	return instance.lastName
}

func (instance *User) Email() string {
	return instance.email
}

func (instance *User) Password() string {
	return instance.password
}

func (instance *User) Roles() []role.Role {
	return instance.roles
}

func (instance *User) AccessToken() string {
	return instance.accessToken
}

func (instance *User) RefreshToken() string {
	return instance.refreshToken
}

func (instance *User) ActivationCode() string {
	return instance.activationCode
}

func (instance *User) Active() bool {
	return instance.active
}

func (instance *User) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *User) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *User) IsZero() bool {
	return reflect.DeepEqual(instance, &User{})
}
