package external

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	externalAuthor *ExternalAuthor
	invalidFields  []string
}

func NewBuilder() *builder {
	return &builder{externalAuthor: &ExternalAuthor{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "O ID do autor externo é inválido")
		return instance
	}
	instance.externalAuthor.id = id
	return instance
}

func (instance *builder) Name(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		instance.invalidFields = append(instance.invalidFields, "O nome da autor externo é inválido")
		return instance
	}
	instance.externalAuthor.name = name
	return instance
}

func (instance *builder) Type(_type string) *builder {
	_type = strings.TrimSpace(_type)
	if len(_type) == 0 {
		instance.invalidFields = append(instance.invalidFields, "O tipo do autor externo é inválido")
		return instance
	}
	instance.externalAuthor._type = _type
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.externalAuthor.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "A data de criação do registro do autor externo é inválida")
		return instance
	}
	instance.externalAuthor.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "A data de atualização do registro do autor externo é inválida")
		return instance
	}
	instance.externalAuthor.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*ExternalAuthor, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.externalAuthor, nil
}
