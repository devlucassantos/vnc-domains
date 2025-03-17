package externalauthor

import (
	"github.com/devlucassantos/vnc-domains/src/domains/externalauthortype"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type ExternalAuthor struct {
	id        uuid.UUID
	name      string
	_type     externalauthortype.ExternalAuthorType
	active    bool
	createdAt time.Time
	updatedAt time.Time
}

func (instance *ExternalAuthor) NewUpdater() *builder {
	return &builder{externalAuthor: instance}
}

func (instance *ExternalAuthor) Id() uuid.UUID {
	return instance.id
}

func (instance *ExternalAuthor) Name() string {
	return instance.name
}

func (instance *ExternalAuthor) Type() externalauthortype.ExternalAuthorType {
	return instance._type
}

func (instance *ExternalAuthor) Active() bool {
	return instance.active
}

func (instance *ExternalAuthor) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *ExternalAuthor) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *ExternalAuthor) IsZero() bool {
	return reflect.DeepEqual(instance, &ExternalAuthor{})
}
