package eventsituation

import (
	"github.com/google/uuid"
	"reflect"
	"time"
)

type EventSituation struct {
	id          uuid.UUID
	description string
	codes       string
	color       string
	active      bool
	createdAt   time.Time
	updatedAt   time.Time
}

func (instance *EventSituation) NewUpdater() *builder {
	return &builder{eventSituation: instance}
}

func (instance *EventSituation) Id() uuid.UUID {
	return instance.id
}

func (instance *EventSituation) Description() string {
	return instance.description
}

func (instance *EventSituation) Codes() string {
	return instance.codes
}

func (instance *EventSituation) Color() string {
	return instance.color
}

func (instance *EventSituation) Active() bool {
	return instance.active
}

func (instance *EventSituation) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *EventSituation) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *EventSituation) IsZero() bool {
	return reflect.DeepEqual(instance, &EventSituation{})
}
