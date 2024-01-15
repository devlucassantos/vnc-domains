package user

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/domains/role"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	user          *User
	invalidFields []string
}

func NewBuilder() *builder {
	return &builder{user: &User{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUUIDValid(id) {
		instance.invalidFields = append(instance.invalidFields, "O ID do usuário é inválido")
		return instance
	}
	instance.user.id = id
	return instance
}

func (instance *builder) FirstName(firstName string) *builder {
	if len(firstName) < 3 {
		instance.invalidFields = append(instance.invalidFields, "O nome do usuário é inválido")
		return instance
	}
	instance.user.firstName = firstName
	return instance
}

func (instance *builder) LastName(lastName string) *builder {
	if len(lastName) < 3 {
		instance.invalidFields = append(instance.invalidFields, "O sobrenome do usuário é inválido")
		return instance
	}
	instance.user.lastName = lastName
	return instance
}

func (instance *builder) Email(email string) *builder {
	if !utils.IsEmailValid(email) {
		instance.invalidFields = append(instance.invalidFields, "O email do usuário é inválido")
		return instance
	}
	instance.user.email = email
	return instance
}

func (instance *builder) Password(password string) *builder {
	if !utils.IsPasswordValid(password) {
		instance.invalidFields = append(instance.invalidFields, "A senha do usuário é inválida")
		return instance
	}
	instance.user.password = password
	return instance
}

func (instance *builder) Hash(hash string) *builder {
	if len(hash) < 3 {
		instance.invalidFields = append(instance.invalidFields, "O hash do usuário é inválido")
		return instance
	}
	instance.user.hash = hash
	return instance
}

func (instance *builder) Roles(roles []role.Role) *builder {
	instance.user.roles = roles
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.user.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "A data de criação do registro do usuário é inválida")
		return instance
	}
	instance.user.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "A data de atualização do registro do usuário é inválida")
		return instance
	}
	instance.user.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*User, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, ";"))
	}
	return instance.user, nil
}
