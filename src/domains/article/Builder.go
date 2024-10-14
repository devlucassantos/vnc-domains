package article

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/domains/articletype"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	article       *Article
	invalidFields []string
}

func NewBuilder() *builder {
	return &builder{article: &Article{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "The article ID is invalid")
		return instance
	}
	instance.article.id = id
	return instance
}

func (instance *builder) Title(title string) *builder {
	if len(title) < 10 {
		instance.invalidFields = append(instance.invalidFields, "The article title is invalid")
		return instance
	}
	instance.article.title = title
	return instance
}

func (instance *builder) Content(content string) *builder {
	content = strings.TrimSpace(content)
	if len(content) < 10 {
		instance.invalidFields = append(instance.invalidFields, "The content of the article is invalid")
		return instance
	}
	instance.article.content = content
	return instance
}

func (instance *builder) ImageUrl(imageUrl string) *builder {
	if !utils.IsUrlValid(imageUrl) {
		instance.invalidFields = append(instance.invalidFields, "The article image URL is invalid")
		return instance
	}
	instance.article.imageUrl = imageUrl
	return instance
}

func (instance *builder) AverageRating(averageRating float64) *builder {
	if averageRating < 0 || averageRating > 5 {
		instance.invalidFields = append(instance.invalidFields, "The value of the average public rating of the article is invalid")
		return instance
	}
	instance.article.averageRating = averageRating
	return instance
}

func (instance *builder) NumberOfRatings(numberOfRatings int) *builder {
	if numberOfRatings < 0 {
		instance.invalidFields = append(instance.invalidFields, "The number of public ratings of the article is invalid")
		return instance
	}
	instance.article.numberOfRatings = numberOfRatings
	return instance
}

func (instance *builder) UserRating(userRating int) *builder {
	if userRating < 0 || userRating > 5 {
		instance.invalidFields = append(instance.invalidFields, "The article rating value is invalid")
		return instance
	}
	instance.article.userRating = userRating
	return instance
}

func (instance *builder) ViewLater(viewLater bool) *builder {
	instance.article.viewLater = viewLater
	return instance
}

func (instance *builder) ViewLaterSetAt(viewLaterSetAt time.Time) *builder {
	if viewLaterSetAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The date and time the article was marked for later viewing are invalid")
		return instance
	}
	instance.article.viewLaterSetAt = viewLaterSetAt
	return instance
}

func (instance *builder) Type(_type articletype.ArticleType) *builder {
	instance.article._type = _type
	return instance
}

func (instance *builder) ReferenceDateTime(referenceDateTime time.Time) *builder {
	if referenceDateTime.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The reference date and time of the article are invalid")
		return instance
	}
	instance.article.referenceDateTime = referenceDateTime
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.article.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date of the article record is invalid")
		return instance
	}
	instance.article.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date of the article record is invalid")
		return instance
	}
	instance.article.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*Article, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.article, nil
}
