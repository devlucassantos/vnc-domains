package newsletter

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/domains/article"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	newsletter    *Newsletter
	invalidFields []string
}

func NewBuilder() *builder {
	return &builder{newsletter: &Newsletter{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "The newsletter ID is invalid")
		return instance
	}
	instance.newsletter.id = id
	return instance
}

func (instance *builder) ReferenceDate(referenceDate time.Time) *builder {
	if referenceDate.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The reference date of the newsletter is invalid")
		return instance
	}
	instance.newsletter.referenceDate = referenceDate
	return instance
}

func (instance *builder) Title(title string) *builder {
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The newsletter title is invalid")
		return instance
	}
	instance.newsletter.title = title
	return instance
}

func (instance *builder) Description(description string) *builder {
	description = strings.TrimSpace(description)
	if len(description) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The newsletter description is invalid")
		return instance
	}
	instance.newsletter.description = description
	return instance
}

func (instance *builder) Article(article article.Article) *builder {
	instance.newsletter.article = article
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.newsletter.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date and time of the newsletter record is invalid")
		return instance
	}
	instance.newsletter.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date and time of the newsletter record is invalid")
		return instance
	}
	instance.newsletter.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*Newsletter, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.newsletter, nil
}
