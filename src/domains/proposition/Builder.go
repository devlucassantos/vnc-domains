package proposition

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/domains/article"
	"github.com/devlucassantos/vnc-domains/src/domains/deputy"
	"github.com/devlucassantos/vnc-domains/src/domains/externalauthor"
	"github.com/devlucassantos/vnc-domains/src/domains/propositiontype"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	proposition   *Proposition
	invalidFields []string
}

func NewBuilder() *builder {
	return &builder{proposition: &Proposition{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "The proposition ID is invalid")
		return instance
	}
	instance.proposition.id = id
	return instance
}

func (instance *builder) Code(code int) *builder {
	if code <= 0 {
		instance.invalidFields = append(instance.invalidFields, "The proposition code is invalid")
		return instance
	}
	instance.proposition.code = code
	return instance
}

func (instance *builder) OriginalTextUrl(originalTextUrl string) *builder {
	if !utils.IsUrlValid(originalTextUrl) {
		instance.invalidFields = append(instance.invalidFields, "The proposition original text URL is invalid")
		return instance
	}
	instance.proposition.originalTextUrl = originalTextUrl
	return instance
}

func (instance *builder) OriginalTextMimeType(originalTextMimeType string) *builder {
	originalTextMimeType = strings.TrimSpace(originalTextMimeType)
	if len(originalTextMimeType) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The proposition original text mime type is invalid")
		return instance
	}
	instance.proposition.originalTextMimeType = originalTextMimeType
	return instance
}

func (instance *builder) Title(title string) *builder {
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The proposition title is invalid")
		return instance
	}
	instance.proposition.title = title
	return instance
}

func (instance *builder) Content(content string) *builder {
	content = strings.TrimSpace(content)
	if len(content) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The proposition content is invalid")
		return instance
	}
	instance.proposition.content = content
	return instance
}

func (instance *builder) SubmittedAt(submittedAt time.Time) *builder {
	if submittedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The submission date and time of the proposition is invalid")
		return instance
	}
	instance.proposition.submittedAt = submittedAt
	return instance
}

func (instance *builder) ImageUrl(imageUrl string) *builder {
	if !utils.IsUrlValid(imageUrl) {
		instance.invalidFields = append(instance.invalidFields, "The proposition image URL is invalid")
		return instance
	}
	instance.proposition.imageUrl = imageUrl
	return instance
}

func (instance *builder) ImageDescription(imageDescription string) *builder {
	imageDescription = strings.TrimSpace(imageDescription)
	if len(imageDescription) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The proposition image description is invalid")
		return instance
	}
	instance.proposition.imageDescription = imageDescription
	return instance
}

func (instance *builder) SpecificType(specificType string) *builder {
	specificType = strings.TrimSpace(specificType)
	if len(specificType) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The proposition specific type is invalid")
		return instance
	}
	instance.proposition.specificType = specificType
	return instance
}

func (instance *builder) Type(_type propositiontype.PropositionType) *builder {
	instance.proposition._type = _type
	return instance
}

func (instance *builder) Deputies(deputies []deputy.Deputy) *builder {
	instance.proposition.deputies = deputies
	return instance
}

func (instance *builder) ExternalAuthors(externalAuthors []externalauthor.ExternalAuthor) *builder {
	instance.proposition.externalAuthors = externalAuthors
	return instance
}

func (instance *builder) Article(article article.Article) *builder {
	instance.proposition.article = article
	return instance
}

func (instance *builder) RelatedArticles(relatedArticles []article.Article) *builder {
	instance.proposition.relatedArticles = relatedArticles
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.proposition.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date and time of the proposition record is invalid")
		return instance
	}
	instance.proposition.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date and time of the proposition record is invalid")
		return instance
	}
	instance.proposition.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*Proposition, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.proposition, nil
}
