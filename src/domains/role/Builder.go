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
	if !utils.IsUUIDValid(id) {
		instance.invalidFields = append(instance.invalidFields, "O ID do função é inválido")
		return instance
	}
	instance.role.id = id
	return instance
}

func (instance *builder) Description(description string) *builder {
	if len(description) < 2 {
		instance.invalidFields = append(instance.invalidFields, "A descrição da função é inválida")
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
		instance.invalidFields = append(instance.invalidFields, "A data de criação do registro da função é inválida")
		return instance
	}
	instance.role.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "A data de atualização do registro da função é inválida")
		return instance
	}
	instance.role.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*Role, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, ";"))
	}
	return instance.role, nil
}
