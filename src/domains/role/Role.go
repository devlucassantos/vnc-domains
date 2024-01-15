package role

import (
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Role struct {
	id          uuid.UUID
	description string
	active      bool
	createdAt   time.Time
	updatedAt   time.Time
}

func (instance *Role) NewUpdater() *builder {
	return &builder{role: instance}
}

func (instance *Role) Id() uuid.UUID {
	return instance.id
}

func (instance *Role) Description() string {
	return instance.description
}

func (instance *Role) Active() bool {
	return instance.active
}

func (instance *Role) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *Role) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *Role) IsZero() bool {
	return reflect.DeepEqual(instance, &Role{})
}
