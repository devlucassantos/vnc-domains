package eventagendaitem

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/domains/agendaitemregime"
	"github.com/devlucassantos/vnc-domains/src/domains/deputy"
	"github.com/devlucassantos/vnc-domains/src/domains/proposition"
	"github.com/devlucassantos/vnc-domains/src/domains/voting"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	eventAgendaItem *EventAgendaItem
	invalidFields   []string
}

func NewBuilder() *builder {
	return &builder{eventAgendaItem: &EventAgendaItem{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "The event agenda item ID is invalid")
		return instance
	}
	instance.eventAgendaItem.id = id
	return instance
}

func (instance *builder) Title(title string) *builder {
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The event agenda item title is invalid")
		return instance
	}
	instance.eventAgendaItem.title = title
	return instance
}

func (instance *builder) Topic(topic string) *builder {
	topic = strings.TrimSpace(topic)
	if len(topic) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The event agenda item topic is invalid")
		return instance
	}
	instance.eventAgendaItem.topic = topic
	return instance
}

func (instance *builder) Situation(situation string) *builder {
	situation = strings.TrimSpace(situation)
	if len(situation) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The event agenda item situation is invalid")
		return instance
	}
	instance.eventAgendaItem.situation = situation
	return instance
}

func (instance *builder) Regime(regime agendaitemregime.AgendaItemRegime) *builder {
	instance.eventAgendaItem.regime = regime
	return instance
}

func (instance *builder) Rapporteur(rapporteur deputy.Deputy) *builder {
	instance.eventAgendaItem.rapporteur = rapporteur
	return instance
}

func (instance *builder) Proposition(proposition proposition.Proposition) *builder {
	instance.eventAgendaItem.proposition = proposition
	return instance
}

func (instance *builder) RelatedProposition(relatedProposition proposition.Proposition) *builder {
	instance.eventAgendaItem.relatedProposition = relatedProposition
	return instance
}

func (instance *builder) Voting(voting voting.Voting) *builder {
	instance.eventAgendaItem.voting = voting
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.eventAgendaItem.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date and time of the event agenda item record is invalid")
		return instance
	}
	instance.eventAgendaItem.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date and time of the event agenda item record is invalid")
		return instance
	}
	instance.eventAgendaItem.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*EventAgendaItem, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.eventAgendaItem, nil
}
