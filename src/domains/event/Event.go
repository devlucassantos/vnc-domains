package event

import (
	"github.com/devlucassantos/vnc-domains/src/domains/article"
	"github.com/devlucassantos/vnc-domains/src/domains/eventagendaitem"
	"github.com/devlucassantos/vnc-domains/src/domains/eventsituation"
	"github.com/devlucassantos/vnc-domains/src/domains/eventtype"
	"github.com/devlucassantos/vnc-domains/src/domains/legislativebody"
	"github.com/devlucassantos/vnc-domains/src/domains/proposition"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Event struct {
	id                uuid.UUID
	code              int
	title             string
	description       string
	startsAt          time.Time
	endsAt            time.Time
	location          string
	isInternal        bool
	videoUrl          string
	specificType      string
	_type             eventtype.EventType
	specificSituation string
	situation         eventsituation.EventSituation
	legislativeBodies []legislativebody.LegislativeBody
	requirements      []proposition.Proposition
	agendaItems       []eventagendaitem.EventAgendaItem
	article           article.Article
	relatedArticles   []article.Article
	active            bool
	createdAt         time.Time
	updatedAt         time.Time
}

func (instance *Event) NewUpdater() *builder {
	return &builder{event: instance}
}

func (instance *Event) Id() uuid.UUID {
	return instance.id
}

func (instance *Event) Code() int {
	return instance.code
}

func (instance *Event) Title() string {
	return instance.title
}

func (instance *Event) Description() string {
	return instance.description
}

func (instance *Event) StartsAt() time.Time {
	return instance.startsAt
}

func (instance *Event) EndsAt() time.Time {
	return instance.endsAt
}

func (instance *Event) Location() string {
	return instance.location
}

func (instance *Event) IsInternal() bool {
	return instance.isInternal
}

func (instance *Event) VideoUrl() string {
	return instance.videoUrl
}

func (instance *Event) SpecificType() string {
	return instance.specificType
}

func (instance *Event) Type() eventtype.EventType {
	return instance._type
}

func (instance *Event) SpecificSituation() string {
	return instance.specificSituation
}

func (instance *Event) Situation() eventsituation.EventSituation {
	return instance.situation
}

func (instance *Event) LegislativeBodies() []legislativebody.LegislativeBody {
	return instance.legislativeBodies
}

func (instance *Event) Requirements() []proposition.Proposition {
	return instance.requirements
}

func (instance *Event) AgendaItems() []eventagendaitem.EventAgendaItem {
	return instance.agendaItems
}

func (instance *Event) Article() article.Article {
	return instance.article
}

func (instance *Event) RelatedArticles() []article.Article {
	return instance.relatedArticles
}

func (instance *Event) Active() bool {
	return instance.active
}

func (instance *Event) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *Event) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *Event) IsEqual(event Event) bool {
	return instance.code == event.code &&
		instance.description == event.description &&
		instance.startsAt == event.startsAt &&
		instance.endsAt == event.endsAt &&
		instance.location == event.location &&
		instance.isInternal == event.isInternal &&
		instance.videoUrl == event.videoUrl &&
		instance.specificSituation == event.specificSituation &&
		instance.situation.IsEqual(event.situation)
}

func (instance *Event) IsZero() bool {
	return reflect.DeepEqual(instance, &Event{})
}
