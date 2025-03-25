package eventtype

import (
	"github.com/google/uuid"
	"reflect"
	"time"
)

type EventType struct {
	id          uuid.UUID
	description string
	codes       string
	color       string
	sortOrder   int
	active      bool
	createdAt   time.Time
	updatedAt   time.Time
}

func (instance *EventType) NewUpdater() *builder {
	return &builder{eventType: instance}
}

func (instance *EventType) Id() uuid.UUID {
	return instance.id
}

func (instance *EventType) Description() string {
	return instance.description
}

func (instance *EventType) Codes() string {
	return instance.codes
}

func (instance *EventType) Color() string {
	return instance.color
}

func (instance *EventType) SortOrder() int {
	return instance.sortOrder
}

func (instance *EventType) Active() bool {
	return instance.active
}

func (instance *EventType) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *EventType) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *EventType) IsZero() bool {
	return reflect.DeepEqual(instance, &EventType{})
}
