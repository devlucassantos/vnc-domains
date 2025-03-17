package deputy

import (
	"errors"
	"github.com/devlucassantos/vnc-domains/src/domains/party"
	"github.com/devlucassantos/vnc-domains/src/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type builder struct {
	deputy        *Deputy
	invalidFields []string
}

func NewBuilder() *builder {
	return &builder{deputy: &Deputy{}}
}

func (instance *builder) Id(id uuid.UUID) *builder {
	if !utils.IsUuidValid(id) {
		instance.invalidFields = append(instance.invalidFields, "The deputy ID is invalid")
		return instance
	}
	instance.deputy.id = id
	return instance
}

func (instance *builder) Code(code int) *builder {
	if code <= 0 {
		instance.invalidFields = append(instance.invalidFields, "The deputy code is invalid")
		return instance
	}
	instance.deputy.code = code
	return instance
}

func (instance *builder) Cpf(cpf string) *builder {
	cpf = strings.TrimSpace(cpf)
	if len(cpf) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The deputy CPF is invalid")
		return instance
	}
	instance.deputy.cpf = cpf
	return instance
}

func (instance *builder) Name(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The deputy name is invalid")
		return instance
	}
	instance.deputy.name = name
	return instance
}

func (instance *builder) ElectoralName(electoralName string) *builder {
	electoralName = strings.TrimSpace(electoralName)
	if len(electoralName) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The deputy electoral name is invalid")
		return instance
	}
	instance.deputy.electoralName = electoralName
	return instance
}

func (instance *builder) ImageUrl(imageUrl string) *builder {
	if !utils.IsUrlValid(imageUrl) {
		instance.invalidFields = append(instance.invalidFields, "The deputy image URL is invalid")
		return instance
	}
	instance.deputy.imageUrl = imageUrl
	return instance
}

func (instance *builder) ImageDescription(imageDescription string) *builder {
	imageDescription = strings.TrimSpace(imageDescription)
	if len(imageDescription) == 0 {
		instance.invalidFields = append(instance.invalidFields, "The deputy image description is invalid")
		return instance
	}
	instance.deputy.imageDescription = imageDescription
	return instance
}

func (instance *builder) Party(party party.Party) *builder {
	instance.deputy.party = party
	return instance
}

func (instance *builder) FederatedUnit(federatedUnit string) *builder {
	federatedUnit = strings.TrimSpace(federatedUnit)
	if len(federatedUnit) != 2 {
		instance.invalidFields = append(instance.invalidFields, "The deputy federated unit is invalid")
		return instance
	}
	instance.deputy.federatedUnit = federatedUnit
	return instance
}

func (instance *builder) PreviousParty(previousParty party.Party) *builder {
	instance.deputy.previousParty = previousParty
	return instance
}

func (instance *builder) PreviousFederatedUnit(previousFederatedUnit string) *builder {
	previousFederatedUnit = strings.TrimSpace(previousFederatedUnit)
	if len(previousFederatedUnit) != 2 {
		instance.invalidFields = append(instance.invalidFields, "The deputy previous federated unit is invalid")
		return instance
	}
	instance.deputy.previousFederatedUnit = previousFederatedUnit
	return instance
}

func (instance *builder) Active(active bool) *builder {
	instance.deputy.active = active
	return instance
}

func (instance *builder) CreatedAt(createdAt time.Time) *builder {
	if createdAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The creation date and time of the deputy record is invalid")
		return instance
	}
	instance.deputy.createdAt = createdAt
	return instance
}

func (instance *builder) UpdatedAt(updatedAt time.Time) *builder {
	if updatedAt.IsZero() {
		instance.invalidFields = append(instance.invalidFields, "The update date and time of the deputy record is invalid")
		return instance
	}
	instance.deputy.updatedAt = updatedAt
	return instance
}

func (instance *builder) Build() (*Deputy, error) {
	if len(instance.invalidFields) > 0 {
		return nil, errors.New(strings.Join(instance.invalidFields, "; "))
	}
	return instance.deputy, nil
}
