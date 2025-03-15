package newsletter

import (
	"github.com/devlucassantos/vnc-domains/src/domains/article"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Newsletter struct {
	id            uuid.UUID
	referenceDate time.Time
	title         string
	description   string
	articles      []article.Article
	article       article.Article
	active        bool
	createdAt     time.Time
	updatedAt     time.Time
}

func (instance *Newsletter) NewUpdater() *builder {
	return &builder{newsletter: instance}
}

func (instance *Newsletter) Id() uuid.UUID {
	return instance.id
}

func (instance *Newsletter) ReferenceDate() time.Time {
	return instance.referenceDate
}

func (instance *Newsletter) Title() string {
	return instance.title
}

func (instance *Newsletter) Description() string {
	return instance.description
}

func (instance *Newsletter) Articles() []article.Article {
	return instance.articles
}

func (instance *Newsletter) Article() article.Article {
	return instance.article
}

func (instance *Newsletter) Active() bool {
	return instance.active
}

func (instance *Newsletter) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *Newsletter) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *Newsletter) IsZero() bool {
	return reflect.DeepEqual(instance, &Newsletter{})
}
