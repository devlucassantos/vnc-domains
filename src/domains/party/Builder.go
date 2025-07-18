package party

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	party         *Party
	invalidFields []string
}

func NewBuilder() *builder {
	return &builder{party: &Party{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "The party ID is invalid")
		return instance
	}
	instance.party.id = id
	return instance
}

func (instance *builder) Code(code int) *builder {
	if code <= 0 {
		instance.invalidFields = append(instance.invalidFields, "The party code is invalid")
		return instance
	}
	instance.party.code = code
	return instance
}

func (instance *builder) Name(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The party name is invalid")
		return instance
	}
	instance.party.name = name
	return instance
}

func (instance *builder) Acronym(acronym string) *builder {
	acronym = strings.TrimSpace(acronym)
	if len(acronym) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The party acronym is invalid")
		return instance
	}
	instance.party.acronym = acronym
	return instance
}

func (instance *builder) ImageUrl(imageUrl string) *builder {
	if !utils.IsUrlValid(imageUrl) {
		instance.invalidFields = append(instance.invalidFields, "The party image URL is invalid")
		return instance
	}
	instance.party.imageUrl = imageUrl
	return instance
}

func (instance *builder) ImageDescription(imageDescription string) *builder {
	imageDescription = strings.TrimSpace(imageDescription)
	if len(imageDescription) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The party image description is invalid")
		return instance
	}
	instance.party.imageDescription = imageDescription
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.party.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date and time of the party record is invalid")
		return instance
	}
	instance.party.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date and time of the party record is invalid")
		return instance
	}
	instance.party.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*Party, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.party, nil
}
