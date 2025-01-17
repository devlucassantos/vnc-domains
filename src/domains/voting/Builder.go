package voting

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/domains/legislativebody"
	"github.com/devlucassantos/vnc-domains/src/domains/proposition"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	voting        *Voting
	invalidFields []string
}

func NewBuilder() *builder {
	return &builder{voting: &Voting{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "The voting ID is invalid")
		return instance
	}
	instance.voting.id = id
	return instance
}

func (instance *builder) Code(code string) *builder {
	code = strings.TrimSpace(code)
	if len(code) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The voting code is invalid")
		return instance
	}
	instance.voting.code = code
	return instance
}

func (instance *builder) Result(result string) *builder {
	result = strings.TrimSpace(result)
	if len(result) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The voting result is invalid")
		return instance
	}
	instance.voting.result = result
	return instance
}

func (instance *builder) ResultAnnouncedAt(resultAnnouncedAt time.Time) *builder {
	if resultAnnouncedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The date and time of the voting result announcement is invalid")
		return instance
	}
	instance.voting.resultAnnouncedAt = resultAnnouncedAt
	return instance
}

func (instance *builder) IsApproved(isApproved bool) *builder {
	instance.voting.isApproved = isApproved
	return instance
}

func (instance *builder) LegislativeBody(legislativeBody legislativebody.LegislativeBody) *builder {
	instance.voting.legislativeBody = legislativeBody
	return instance
}

func (instance *builder) MainProposition(mainProposition proposition.Proposition) *builder {
	instance.voting.mainProposition = mainProposition
	return instance
}

func (instance *builder) RelatedPropositions(relatedPropositions []proposition.Proposition) *builder {
	instance.voting.relatedPropositions = relatedPropositions
	return instance
}

func (instance *builder) AffectedPropositions(affectedPropositions []proposition.Proposition) *builder {
	instance.voting.affectedPropositions = affectedPropositions
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.voting.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date and time of the event type record is invalid")
		return instance
	}
	instance.voting.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date and time of the event type record is invalid")
		return instance
	}
	instance.voting.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*Voting, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.voting, nil
}
