package externalauthortype

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	externalAuthorType *ExternalAuthorType
	invalidFields      []string
}

func NewBuilder() *builder {
	return &builder{externalAuthorType: &ExternalAuthorType{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "The external author type ID is invalid")
		return instance
	}
	instance.externalAuthorType.id = id
	return instance
}

func (instance *builder) Code(code int) *builder {
	if code <= 0 {
		instance.invalidFields = append(instance.invalidFields, "The external author type code is invalid")
		return instance
	}
	instance.externalAuthorType.code = code
	return instance
}

func (instance *builder) Description(description string) *builder {
	description = strings.TrimSpace(description)
	if len(description) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The external author type description is invalid")
		return instance
	}
	instance.externalAuthorType.description = description
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.externalAuthorType.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date and time of the external author type record is invalid")
		return instance
	}
	instance.externalAuthorType.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date and time of the external author type record is invalid")
		return instance
	}
	instance.externalAuthorType.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*ExternalAuthorType, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.externalAuthorType, nil
}
