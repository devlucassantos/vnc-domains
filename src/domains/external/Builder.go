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
		instance.invalidFields = append(instance.invalidFields, "The external author ID is invalid")
		return instance
	}
	instance.externalAuthor.id = id
	return instance
}

func (instance *builder) Name(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The external author name is invalid")
		return instance
	}
	instance.externalAuthor.name = name
	return instance
}

func (instance *builder) Type(_type string) *builder {
	_type = strings.TrimSpace(_type)
	if len(_type) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The external author type is invalid")
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
		instance.invalidFields = append(instance.invalidFields, "The creation date of the external author record is invalid")
		return instance
	}
	instance.externalAuthor.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date of the external author record is invalid")
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
