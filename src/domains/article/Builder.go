package article

import (
	"errors"
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
		instance.invalidFields = append(instance.invalidFields, "O ID da matéria é inválido")
		return instance
	}
	instance.article.id = id
	return instance
}

func (instance *builder) Title(title string) *builder {
	if len(title) < 10 {
		instance.invalidFields = append(instance.invalidFields, "O título da matéria é inválido")
		return instance
	}
	instance.article.title = title
	return instance
}

func (instance *builder) Content(content string) *builder {
	content = strings.TrimSpace(content)
	if len(content) < 10 {
		instance.invalidFields = append(instance.invalidFields, "O conteúdo da matéria é inválido")
		return instance
	}
	instance.article.content = content
	return instance
}

func (instance *builder) ImageUrl(imageUrl string) *builder {
	if !utils.IsUrlValid(imageUrl) {
		instance.invalidFields = append(instance.invalidFields, "A URL da imagem da matéria é inválida")
		return instance
	}
	instance.article.imageUrl = imageUrl
	return instance
}

func (instance *builder) AverageRating(averageRating float64) *builder {
	if averageRating < 1 && averageRating > 5 {
		instance.invalidFields = append(instance.invalidFields, "O valor da média da avaliação da matéria pelo público é inválido")
		return instance
	}
	instance.article.averageRating = averageRating
	return instance
}

func (instance *builder) NumberOfRatings(NumberOfRatings int) *builder {
	if NumberOfRatings < 0 {
		instance.invalidFields = append(instance.invalidFields, "O número de avaliações da matéria pelo público é inválido")
		return instance
	}
	instance.article.numberOfRatings = NumberOfRatings
	return instance
}

func (instance *builder) UserRating(userRating int) *builder {
	if userRating < 1 || userRating > 5 {
		instance.invalidFields = append(instance.invalidFields, "O valor da avaliação da matéria é inválido")
		return instance
	}
	instance.article.userRating = userRating
	return instance
}

func (instance *builder) ViewLater(viewLater bool) *builder {
	instance.article.viewLater = viewLater
	return instance
}

func (instance *builder) Type(_type string) *builder {
	_type = strings.TrimSpace(_type)
	if len(_type) == 0 {
		instance.invalidFields = append(instance.invalidFields, "O tipo da matéria é inválido")
		return instance
	}
	instance.article._type = _type
	return instance
}

func (instance *builder) ReferenceDateTime(referenceDateTime time.Time) *builder {
	if referenceDateTime.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "A data e hora de referência da matéria são inválidas")
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
		instance.invalidFields = append(instance.invalidFields, "A data de criação do registro da matéria é inválida")
		return instance
	}
	instance.article.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "A data de atualização do registro da matéria é inválida")
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
