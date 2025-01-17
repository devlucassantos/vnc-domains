package legislativebody

import (
	"github.com/devlucassantos/vnc-domains/src/domains/legislativebodytype"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type LegislativeBody struct {
	id        uuid.UUID
	code      int
	name      string
	acronym   string
	_type     legislativebodytype.LegislativeBodyType
	active    bool
	createdAt time.Time
	updatedAt time.Time
}

func (instance *LegislativeBody) NewUpdater() *builder {
	return &builder{legislativeBody: instance}
}

func (instance *LegislativeBody) Id() uuid.UUID {
	return instance.id
}

func (instance *LegislativeBody) Code() int {
	return instance.code
}

func (instance *LegislativeBody) Name() string {
	return instance.name
}

func (instance *LegislativeBody) Acronym() string {
	return instance.acronym
}

func (instance *LegislativeBody) Type() legislativebodytype.LegislativeBodyType {
	return instance._type
}

func (instance *LegislativeBody) Active() bool {
	return instance.active
}

func (instance *LegislativeBody) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *LegislativeBody) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *LegislativeBody) IsZero() bool {
	return reflect.DeepEqual(instance, &LegislativeBody{})
}
