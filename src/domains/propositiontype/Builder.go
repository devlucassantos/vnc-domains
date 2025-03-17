package propositiontype

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
		instance.invalidFields = append(instance.invalidFields, "The proposition type ID is invalid")
		return instance
	}
	instance.propositionType.id = id
	return instance
}

func (instance *builder) Description(description string) *builder {
	description = strings.TrimSpace(description)
	if len(description) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The proposition type description is invalid")
		return instance
	}
	instance.propositionType.description = description
	return instance
}

func (instance *builder) Codes(codes string) *builder {
	codes = strings.TrimSpace(codes)
	if len(codes) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The proposition type codes are invalid")
		return instance
	}
	instance.propositionType.codes = codes
	return instance
}

func (instance *builder) Color(color string) *builder {
	color = strings.TrimSpace(color)
	if len(color) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The proposition type color is invalid")
		return instance
	}
	instance.propositionType.color = color
	return instance
}

func (instance *builder) SortOrder(sortOrder int) *builder {
	if sortOrder <= 0 {
		instance.invalidFields = append(instance.invalidFields, "The proposition type sort order is invalid")
		return instance
	}
	instance.propositionType.sortOrder = sortOrder
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.propositionType.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date and time of the proposition type record is invalid")
		return instance
	}
	instance.propositionType.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date and time of the proposition type record is invalid")
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
