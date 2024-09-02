package articletype

import (
	"github.com/google/uuid"
	"reflect"
	"time"
)

type ArticleType struct {
	id          uuid.UUID
	description string
	codes       string
	color       string
	sortOrder   int
	active      bool
	createdAt   time.Time
	updatedAt   time.Time
}

func (instance *ArticleType) NewUpdater() *builder {
	return &builder{articleType: instance}
}

func (instance *ArticleType) Id() uuid.UUID {
	return instance.id
}

func (instance *ArticleType) Description() string {
	return instance.description
}

func (instance *ArticleType) Codes() string {
	return instance.codes
}

func (instance *ArticleType) Color() string {
	return instance.color
}

func (instance *ArticleType) SortOrder() int {
	return instance.sortOrder
}

func (instance *ArticleType) Active() bool {
	return instance.active
}

func (instance *ArticleType) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *ArticleType) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *ArticleType) IsZero() bool {
	return reflect.DeepEqual(instance, &ArticleType{})
}
