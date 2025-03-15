package proposition

import (
	"github.com/devlucassantos/vnc-domains/src/domains/article"
	"github.com/devlucassantos/vnc-domains/src/domains/deputy"
	"github.com/devlucassantos/vnc-domains/src/domains/externalauthor"
	"github.com/devlucassantos/vnc-domains/src/domains/propositiontype"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Proposition struct {
	id                   uuid.UUID
	code                 int
	originalTextUrl      string
	originalTextMimeType string
	title                string
	content              string
	submittedAt          time.Time
	imageUrl             string
	imageDescription     string
	specificType         string
	_type                propositiontype.PropositionType
	deputies             []deputy.Deputy
	externalAuthors      []externalauthor.ExternalAuthor
	article              article.Article
	relatedArticles      []article.Article
	active               bool
	createdAt            time.Time
	updatedAt            time.Time
}

func (instance *Proposition) NewUpdater() *builder {
	return &builder{proposition: instance}
}

func (instance *Proposition) Id() uuid.UUID {
	return instance.id
}

func (instance *Proposition) Code() int {
	return instance.code
}

func (instance *Proposition) OriginalTextUrl() string {
	return instance.originalTextUrl
}

func (instance *Proposition) OriginalTextMimeType() string {
	return instance.originalTextMimeType
}

func (instance *Proposition) Title() string {
	return instance.title
}

func (instance *Proposition) Content() string {
	return instance.content
}

func (instance *Proposition) SubmittedAt() time.Time {
	return instance.submittedAt
}

func (instance *Proposition) ImageUrl() string {
	return instance.imageUrl
}

func (instance *Proposition) ImageDescription() string {
	return instance.imageDescription
}

func (instance *Proposition) SpecificType() string {
	return instance.specificType
}

func (instance *Proposition) Type() propositiontype.PropositionType {
	return instance._type
}

func (instance *Proposition) Deputies() []deputy.Deputy {
	return instance.deputies
}

func (instance *Proposition) ExternalAuthors() []externalauthor.ExternalAuthor {
	return instance.externalAuthors
}

func (instance *Proposition) Article() article.Article {
	return instance.article
}

func (instance *Proposition) RelatedArticles() []article.Article {
	return instance.relatedArticles
}

func (instance *Proposition) Active() bool {
	return instance.active
}

func (instance *Proposition) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *Proposition) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *Proposition) IsZero() bool {
	return reflect.DeepEqual(instance, &Proposition{})
}
