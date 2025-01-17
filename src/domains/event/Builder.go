package event

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/domains/eventagendaitem"
	"github.com/devlucassantos/vnc-domains/src/domains/eventsituation"
	"github.com/devlucassantos/vnc-domains/src/domains/eventtype"
	"github.com/devlucassantos/vnc-domains/src/domains/legislativebody"
	"github.com/devlucassantos/vnc-domains/src/domains/proposition"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	event         *Event
	invalidFields []string
}

func NewBuilder() *builder {
	return &builder{event: &Event{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "The event ID is invalid")
		return instance
	}
	instance.event.id = id
	return instance
}

func (instance *builder) Code(code int) *builder {
	if code <= 0 {
		instance.invalidFields = append(instance.invalidFields, "The event code is invalid")
		return instance
	}
	instance.event.code = code
	return instance
}

func (instance *builder) Title(title string) *builder {
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The event title is invalid")
		return instance
	}
	instance.event.title = title
	return instance
}

func (instance *builder) Description(description string) *builder {
	description = strings.TrimSpace(description)
	if len(description) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The event description is invalid")
		return instance
	}
	instance.event.description = description
	return instance
}

func (instance *builder) StartsAt(startsAt time.Time) *builder {
	if startsAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The start date and time of the event is invalid")
		return instance
	}
	instance.event.startsAt = startsAt
	return instance
}

func (instance *builder) EndsAt(endsAt time.Time) *builder {
	if endsAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The end date and time of the event is invalid")
		return instance
	}
	instance.event.endsAt = endsAt
	return instance
}

func (instance *builder) Location(location string) *builder {
	location = strings.TrimSpace(location)
	if len(location) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The event location is invalid")
		return instance
	}
	instance.event.location = location
	return instance
}

func (instance *builder) IsInternal(isInternal bool) *builder {
	instance.event.isInternal = isInternal
	return instance
}

func (instance *builder) VideoUrl(videoUrl string) *builder {
	if !utils.IsUrlValid(videoUrl) {
		instance.invalidFields = append(instance.invalidFields, "The event video URL is invalid")
		return instance
	}
	instance.event.videoUrl = videoUrl
	return instance
}

func (instance *builder) SpecificType(specificType string) *builder {
	specificType = strings.TrimSpace(specificType)
	if len(specificType) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The event specific type is invalid")
		return instance
	}
	instance.event.specificType = specificType
	return instance
}

func (instance *builder) Type(_type eventtype.EventType) *builder {
	instance.event._type = _type
	return instance
}

func (instance *builder) SpecificSituation(specificSituation string) *builder {
	specificSituation = strings.TrimSpace(specificSituation)
	if len(specificSituation) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The event specific situation is invalid")
		return instance
	}
	instance.event.specificSituation = specificSituation
	return instance
}

func (instance *builder) Situation(situation eventsituation.EventSituation) *builder {
	instance.event.situation = situation
	return instance
}

func (instance *builder) LegislativeBodies(legislativeBodies []legislativebody.LegislativeBody) *builder {
	instance.event.legislativeBodies = legislativeBodies
	return instance
}

func (instance *builder) Requirements(requirements []proposition.Proposition) *builder {
	instance.event.requirements = requirements
	return instance
}

func (instance *builder) AgendaItems(agendaItems []eventagendaitem.EventAgendaItem) *builder {
	instance.event.agendaItems = agendaItems
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.event.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date and time of the event record is invalid")
		return instance
	}
	instance.event.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date and time of the event record is invalid")
		return instance
	}
	instance.event.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*Event, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.event, nil
}
