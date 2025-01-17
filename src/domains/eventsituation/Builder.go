package eventsituation

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	eventSituation *EventSituation
	invalidFields  []string
}

func NewBuilder() *builder {
	return &builder{eventSituation: &EventSituation{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "The event situation ID is invalid")
		return instance
	}
	instance.eventSituation.id = id
	return instance
}

func (instance *builder) Description(description string) *builder {
	description = strings.TrimSpace(description)
	if len(description) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The event situation description is invalid")
		return instance
	}
	instance.eventSituation.description = description
	return instance
}

func (instance *builder) Codes(codes string) *builder {
	codes = strings.TrimSpace(codes)
	if len(codes) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The event situation codes is invalid")
		return instance
	}
	instance.eventSituation.codes = codes
	return instance
}

func (instance *builder) Color(color string) *builder {
	color = strings.TrimSpace(color)
	if len(color) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The event situation color is invalid")
		return instance
	}
	instance.eventSituation.color = color
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.eventSituation.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date and time of the event situation record is invalid")
		return instance
	}
	instance.eventSituation.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date and time of the event situation record is invalid")
		return instance
	}
	instance.eventSituation.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*EventSituation, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.eventSituation, nil
}
