package agendaitemregime

import (
	"github.com/google/uuid"
	"reflect"
	"time"
)

type AgendaItemRegime struct {
	id          uuid.UUID
	code        int
	description string
	active      bool
	createdAt   time.Time
	updatedAt   time.Time
}

func (instance *AgendaItemRegime) NewUpdater() *builder {
	return &builder{agendaItemRegime: instance}
}

func (instance *AgendaItemRegime) Id() uuid.UUID {
	return instance.id
}

func (instance *AgendaItemRegime) Code() int {
	return instance.code
}

func (instance *AgendaItemRegime) Description() string {
	return instance.description
}

func (instance *AgendaItemRegime) Active() bool {
	return instance.active
}

func (instance *AgendaItemRegime) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *AgendaItemRegime) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *AgendaItemRegime) IsZero() bool {
	return reflect.DeepEqual(instance, &AgendaItemRegime{})
}
