package propositiontype

import (
	"github.com/google/uuid"
	"reflect"
	"time"
)

type PropositionType struct {
	id          uuid.UUID
	description string
	codes       string
	color       string
	sortOrder   int
	active      bool
	createdAt   time.Time
	updatedAt   time.Time
}

func (instance *PropositionType) NewUpdater() *builder {
	return &builder{propositionType: instance}
}

func (instance *PropositionType) Id() uuid.UUID {
	return instance.id
}

func (instance *PropositionType) Description() string {
	return instance.description
}

func (instance *PropositionType) Codes() string {
	return instance.codes
}

func (instance *PropositionType) Color() string {
	return instance.color
}

func (instance *PropositionType) SortOrder() int {
	return instance.sortOrder
}

func (instance *PropositionType) Active() bool {
	return instance.active
}

func (instance *PropositionType) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *PropositionType) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *PropositionType) IsZero() bool {
	return reflect.DeepEqual(instance, &PropositionType{})
}
