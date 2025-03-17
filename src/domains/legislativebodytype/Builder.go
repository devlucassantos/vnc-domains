package legislativebodytype

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	legislativeBodyType *LegislativeBodyType
	invalidFields       []string
}

func NewBuilder() *builder {
	return &builder{legislativeBodyType: &LegislativeBodyType{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "The legislative body type ID is invalid")
		return instance
	}
	instance.legislativeBodyType.id = id
	return instance
}

func (instance *builder) Code(code int) *builder {
	if code <= 0 {
		instance.invalidFields = append(instance.invalidFields, "The legislative body type code is invalid")
		return instance
	}
	instance.legislativeBodyType.code = code
	return instance
}

func (instance *builder) Description(description string) *builder {
	description = strings.TrimSpace(description)
	if len(description) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The legislative body type description is invalid")
		return instance
	}
	instance.legislativeBodyType.description = description
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.legislativeBodyType.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date and time of the legislative body type record is invalid")
		return instance
	}
	instance.legislativeBodyType.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date and time of the legislative body type record is invalid")
		return instance
	}
	instance.legislativeBodyType.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*LegislativeBodyType, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.legislativeBodyType, nil
}
