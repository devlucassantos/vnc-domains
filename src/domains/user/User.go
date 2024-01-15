package user

import (
	"github.com/devlucassantos/vnc-domains/src/domains/role"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type User struct {
	id        uuid.UUID
	firstName string
	lastName  string
	email     string
	password  string
	hash      string
	roles     []role.Role
	active    bool
	createdAt time.Time
	updatedAt time.Time
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

func (instance *User) Hash() string {
	return instance.hash
}

func (instance *User) Roles() []role.Role {
	return instance.roles
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
