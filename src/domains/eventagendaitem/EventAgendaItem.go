package eventagendaitem

import (
	"github.com/devlucassantos/vnc-domains/src/domains/agendaitemregime"
	"github.com/devlucassantos/vnc-domains/src/domains/deputy"
	"github.com/devlucassantos/vnc-domains/src/domains/proposition"
	"github.com/devlucassantos/vnc-domains/src/domains/voting"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type EventAgendaItem struct {
	id                 uuid.UUID
	title              string
	topic              string
	situation          string
	regime             agendaitemregime.AgendaItemRegime
	rapporteur         deputy.Deputy
	proposition        proposition.Proposition
	relatedProposition proposition.Proposition
	voting             voting.Voting
	active             bool
	createdAt          time.Time
	updatedAt          time.Time
}

func (instance *EventAgendaItem) NewUpdater() *builder {
	return &builder{eventAgendaItem: instance}
}

func (instance *EventAgendaItem) Id() uuid.UUID {
	return instance.id
}

func (instance *EventAgendaItem) Title() string {
	return instance.title
}

func (instance *EventAgendaItem) Topic() string {
	return instance.topic
}

func (instance *EventAgendaItem) Situation() string {
	return instance.situation
}

func (instance *EventAgendaItem) Regime() agendaitemregime.AgendaItemRegime {
	return instance.regime
}

func (instance *EventAgendaItem) Rapporteur() deputy.Deputy {
	return instance.rapporteur
}

func (instance *EventAgendaItem) Proposition() proposition.Proposition {
	return instance.proposition
}

func (instance *EventAgendaItem) RelatedProposition() proposition.Proposition {
	return instance.relatedProposition
}

func (instance *EventAgendaItem) Voting() voting.Voting {
	return instance.voting
}

func (instance *EventAgendaItem) Active() bool {
	return instance.active
}

func (instance *EventAgendaItem) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *EventAgendaItem) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *EventAgendaItem) IsZero() bool {
	return reflect.DeepEqual(instance, &EventAgendaItem{})
}
