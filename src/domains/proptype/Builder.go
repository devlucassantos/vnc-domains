package proptype

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	propositionType *PropositionType
	invalidFields   []string
}

func NewBuilder() *builder {
	return &builder{propositionType: &PropositionType{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "O ID do tipo da proposição é inválido")
		return instance
	}
	instance.propositionType.id = id
	return instance
}

func (instance *builder) Description(description string) *builder {
	description = strings.TrimSpace(description)
	if len(description) == 0 {
		instance.invalidFields = append(instance.invalidFields, "A descrição do tipo da proposição é inválida")
		return instance
	}
	instance.propositionType.description = description
	return instance
}

func (instance *builder) Codes(codes string) *builder {
	instance.propositionType.codes = codes
	return instance
}

func (instance *builder) Color(color string) *builder {
	instance.propositionType.color = color
	return instance
}

func (instance *builder) SortOrder(sortOrder int) *builder {
	instance.propositionType.sortOrder = sortOrder
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.propositionType.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "A data de criação do registro do tipo da proposição é inválida")
		return instance
	}
	instance.propositionType.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "A data de atualização do registro do tipo da proposição é inválida")
		return instance
	}
	instance.propositionType.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*PropositionType, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.propositionType, nil
}
