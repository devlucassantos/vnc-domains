package articlesituation

import (
	"github.com/google/uuid"
	"reflect"
	"time"
)

type ArticleSituation struct {
	id          uuid.UUID
	description string
	color       string
	startsAt    time.Time
	endsAt      time.Time
	isApproved  bool
}

func (instance *ArticleSituation) NewUpdater() *builder {
	return &builder{articleSituation: instance}
}

func (instance *ArticleSituation) Id() uuid.UUID {
	return instance.id
}

func (instance *ArticleSituation) Description() string {
	return instance.description
}

func (instance *ArticleSituation) Color() string {
	return instance.color
}

func (instance *ArticleSituation) StartsAt() time.Time {
	return instance.startsAt
}

func (instance *ArticleSituation) EndsAt() time.Time {
	return instance.endsAt
}

func (instance *ArticleSituation) IsApproved() bool {
	return instance.isApproved
}

func (instance *ArticleSituation) IsZero() bool {
	return reflect.DeepEqual(instance, &ArticleSituation{})
}
