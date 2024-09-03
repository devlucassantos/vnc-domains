package articletype

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	articleType   *ArticleType
	invalidFields []string
}

func NewBuilder() *builder {
	return &builder{articleType: &ArticleType{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "O ID do tipo da matéria é inválido")
		return instance
	}
	instance.articleType.id = id
	return instance
}

func (instance *builder) Description(description string) *builder {
	description = strings.TrimSpace(description)
	if len(description) == 0 {
		instance.invalidFields = append(instance.invalidFields, "A descrição do tipo da matéria é inválida")
		return instance
	}
	instance.articleType.description = description
	return instance
}

func (instance *builder) Codes(codes string) *builder {
	instance.articleType.codes = codes
	return instance
}

func (instance *builder) Color(color string) *builder {
	instance.articleType.color = color
	return instance
}

func (instance *builder) SortOrder(sortOrder int) *builder {
	instance.articleType.sortOrder = sortOrder
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.articleType.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "A data de criação do registro do tipo da matéria é inválida")
		return instance
	}
	instance.articleType.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "A data de atualização do registro do tipo da matéria é inválida")
		return instance
	}
	instance.articleType.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*ArticleType, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.articleType, nil
}
