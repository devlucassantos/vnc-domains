package role

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	role          *Role
	invalidFields []string
}

func NewBuilder() *builder {
	return &builder{role: &Role{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "The role ID is invalid")
		return instance
	}
	instance.role.id = id
	return instance
}

func (instance *builder) Code(code string) *builder {
	code = strings.TrimSpace(code)
	if len(code) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The role code is invalid")
		return instance
	}
	instance.role.code = code
	return instance
}

func (instance *builder) Description(description string) *builder {
	description = strings.TrimSpace(description)
	if len(description) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The role description is invalid")
		return instance
	}
	instance.role.description = description
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.role.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date and time of the role record is invalid")
		return instance
	}
	instance.role.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date and time of the role record is invalid")
		return instance
	}
	instance.role.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*Role, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.role, nil
}
