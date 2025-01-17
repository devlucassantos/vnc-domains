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
		instance.invalidFields = append(instance.invalidFields, "The article type ID is invalid")
		return instance
	}
	instance.articleType.id = id
	return instance
}

func (instance *builder) Description(description string) *builder {
	description = strings.TrimSpace(description)
	if len(description) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The article type description is invalid")
		return instance
	}
	instance.articleType.description = description
	return instance
}

func (instance *builder) Codes(codes string) *builder {
	codes = strings.TrimSpace(codes)
	if len(codes) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The article type codes is invalid")
		return instance
	}
	instance.articleType.codes = codes
	return instance
}

func (instance *builder) Color(color string) *builder {
	color = strings.TrimSpace(color)
	if len(color) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The article type color is invalid")
		return instance
	}
	instance.articleType.color = color
	return instance
}

func (instance *builder) SortOrder(sortOrder int) *builder {
	if sortOrder <= 0 {
		instance.invalidFields = append(instance.invalidFields, "The article type sort order is invalid")
		return instance
	}
	instance.articleType.sortOrder = sortOrder
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.articleType.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date and time of the article type record is invalid")
		return instance
	}
	instance.articleType.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date and time of the article type record is invalid")
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
