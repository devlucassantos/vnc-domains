package legislativebodytype

import (
	"github.com/google/uuid"
	"reflect"
	"time"
)

type LegislativeBodyType struct {
	id          uuid.UUID
	code        int
	description string
	active      bool
	createdAt   time.Time
	updatedAt   time.Time
}

func (instance *LegislativeBodyType) NewUpdater() *builder {
	return &builder{legislativeBodyType: instance}
}

func (instance *LegislativeBodyType) Id() uuid.UUID {
	return instance.id
}

func (instance *LegislativeBodyType) Code() int {
	return instance.code
}

func (instance *LegislativeBodyType) Description() string {
	return instance.description
}

func (instance *LegislativeBodyType) Active() bool {
	return instance.active
}

func (instance *LegislativeBodyType) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *LegislativeBodyType) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *LegislativeBodyType) IsZero() bool {
	return reflect.DeepEqual(instance, &LegislativeBodyType{})
}
