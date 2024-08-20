package article

import (
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Article struct {
	id                uuid.UUID
	title             string
	content           string
	imageUrl          string
	averageRating     float64
	numberOfRatings   int
	userRating        int
	viewLater         bool
	_type             string
	referenceDateTime time.Time
	active            bool
	createdAt         time.Time
	updatedAt         time.Time
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

func (instance *Article) ImageUrl() string {
	return instance.imageUrl
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

func (instance *Article) Type() string {
	return instance._type
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
