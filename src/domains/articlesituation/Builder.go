package articlesituation

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	articleSituation *ArticleSituation
	invalidFields    []string
}

func NewBuilder() *builder {
	return &builder{articleSituation: &ArticleSituation{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "The article situation ID is invalid")
		return instance
	}
	instance.articleSituation.id = id
	return instance
}

func (instance *builder) Description(description string) *builder {
	description = strings.TrimSpace(description)
	if len(description) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The article situation description is invalid")
		return instance
	}
	instance.articleSituation.description = description
	return instance
}

func (instance *builder) Color(color string) *builder {
	color = strings.TrimSpace(color)
	if len(color) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The article situation color is invalid")
		return instance
	}
	instance.articleSituation.color = color
	return instance
}

func (instance *builder) StartsAt(startsAt time.Time) *builder {
	if startsAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The start date and time of the article is invalid")
		return instance
	}
	instance.articleSituation.startsAt = startsAt
	return instance
}

func (instance *builder) EndsAt(endsAt time.Time) *builder {
	if endsAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The end date and time of the article is invalid")
		return instance
	}
	instance.articleSituation.endsAt = endsAt
	return instance
}

func (instance *builder) Result(result string) *builder {
	result = strings.TrimSpace(result)
	if len(result) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The article result is invalid")
		return instance
	}
	instance.articleSituation.result = result
	return instance
}

func (instance *builder) ResultAnnouncedAt(resultAnnouncedAt time.Time) *builder {
	if resultAnnouncedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The date and time of the article result announcement is invalid")
		return instance
	}
	instance.articleSituation.resultAnnouncedAt = resultAnnouncedAt
	return instance
}

func (instance *builder) IsApproved(isApproved *bool) *builder {
	instance.articleSituation.isApproved = isApproved
	return instance
}

func (instance *builder) Build() (*ArticleSituation, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.articleSituation, nil
}
