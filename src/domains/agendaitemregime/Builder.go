package agendaitemregime

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	agendaItemRegime *AgendaItemRegime
	invalidFields    []string
}

func NewBuilder() *builder {
	return &builder{agendaItemRegime: &AgendaItemRegime{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "The agenda item regime ID is invalid")
		return instance
	}
	instance.agendaItemRegime.id = id
	return instance
}

func (instance *builder) Code(code int) *builder {
	if code <= 0 {
		instance.invalidFields = append(instance.invalidFields, "The agenda item regime code is invalid")
		return instance
	}
	instance.agendaItemRegime.code = code
	return instance
}

func (instance *builder) Description(description string) *builder {
	description = strings.TrimSpace(description)
	if len(description) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The agenda item regime description is invalid")
		return instance
	}
	instance.agendaItemRegime.description = description
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.agendaItemRegime.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date and time of the agenda item regime record is invalid")
		return instance
	}
	instance.agendaItemRegime.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date and time of the agenda item regime record is invalid")
		return instance
	}
	instance.agendaItemRegime.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*AgendaItemRegime, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.agendaItemRegime, nil
}
