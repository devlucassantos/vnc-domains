package legislativebody

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/domains/legislativebodytype"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	legislativeBody *LegislativeBody
	invalidFields   []string
}

func NewBuilder() *builder {
	return &builder{legislativeBody: &LegislativeBody{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "The legislative body ID is invalid")
		return instance
	}
	instance.legislativeBody.id = id
	return instance
}

func (instance *builder) Code(code int) *builder {
	if code <= 0 {
		instance.invalidFields = append(instance.invalidFields, "The legislative body code is invalid")
		return instance
	}
	instance.legislativeBody.code = code
	return instance
}

func (instance *builder) Name(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The legislative body name is invalid")
		return instance
	}
	instance.legislativeBody.name = name
	return instance
}

func (instance *builder) Acronym(acronym string) *builder {
	acronym = strings.TrimSpace(acronym)
	if len(acronym) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The legislative body acronym is invalid")
		return instance
	}
	instance.legislativeBody.acronym = acronym
	return instance
}

func (instance *builder) Type(_type legislativebodytype.LegislativeBodyType) *builder {
	instance.legislativeBody._type = _type
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.legislativeBody.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date and time of the legislative body record is invalid")
		return instance
	}
	instance.legislativeBody.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date and time of the legislative body record is invalid")
		return instance
	}
	instance.legislativeBody.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*LegislativeBody, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.legislativeBody, nil
}
