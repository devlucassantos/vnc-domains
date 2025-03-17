package externalauthortype

import (
	"github.com/google/uuid"
	"reflect"
	"time"
)

type ExternalAuthorType struct {
	id          uuid.UUID
	code        int
	description string
	active      bool
	createdAt   time.Time
	updatedAt   time.Time
}

func (instance *ExternalAuthorType) NewUpdater() *builder {
	return &builder{externalAuthorType: instance}
}

func (instance *ExternalAuthorType) Id() uuid.UUID {
	return instance.id
}

func (instance *ExternalAuthorType) Code() int {
	return instance.code
}

func (instance *ExternalAuthorType) Description() string {
	return instance.description
}

func (instance *ExternalAuthorType) Active() bool {
	return instance.active
}

func (instance *ExternalAuthorType) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *ExternalAuthorType) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *ExternalAuthorType) IsZero() bool {
	return reflect.DeepEqual(instance, &ExternalAuthorType{})
}
