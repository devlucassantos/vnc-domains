package article

import (
	"github.com/devlucassantos/vnc-domains/src/domains/articlesituation"
	"github.com/devlucassantos/vnc-domains/src/domains/articletype"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Article struct {
	id                    uuid.UUID
	title                 string
	content               string
	multimediaUrl         string
	multimediaDescription string
	situation             articlesituation.ArticleSituation
	averageRating         float64
	numberOfRatings       int
	userRating            int
	viewLater             bool
	viewLaterSetAt        time.Time
	_type                 articletype.ArticleType
	specificType          articletype.ArticleType
	referenceDateTime     time.Time
	active                bool
	createdAt             time.Time
	updatedAt             time.Time
}

func (instance *Article) NewUpdater() *builder {
	return &builder{article: instance}
}

func (instance *Article) Id() uuid.UUID {
	return instance.id
}

func (instance *Article) Title() string {
	return instance.title
}

func (instance *Article) Content() string {
	return instance.content
}

func (instance *Article) MultimediaUrl() string {
	return instance.multimediaUrl
}

func (instance *Article) MultimediaDescription() string {
	return instance.multimediaDescription
}

func (instance *Article) Situation() articlesituation.ArticleSituation {
	return instance.situation
}

func (instance *Article) AverageRating() float64 {
	return instance.averageRating
}

func (instance *Article) NumberOfRatings() int {
	return instance.numberOfRatings
}

func (instance *Article) UserRating() int {
	return instance.userRating
}

func (instance *Article) ViewLater() bool {
	return instance.viewLater
}

func (instance *Article) ViewLaterSetAt() time.Time {
	return instance.viewLaterSetAt
}

func (instance *Article) Type() articletype.ArticleType {
	return instance._type
}

func (instance *Article) SpecificType() articletype.ArticleType {
	return instance.specificType
}

func (instance *Article) ReferenceDateTime() time.Time {
	return instance.referenceDateTime
}

func (instance *Article) Active() bool {
	return instance.active
}

func (instance *Article) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *Article) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *Article) IsZero() bool {
	return reflect.DeepEqual(instance, &Article{})
}
