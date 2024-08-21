package deputy

import (
	"github.com/devlucassantos/vnc-domains/src/domains/party"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Deputy struct {
	id                    uuid.UUID
	code                  int
	cpf                   string
	name                  string
	electoralName         string
	imageUrl              string
	party                 party.Party
	partyInTheProposition party.Party
	active                bool
	createdAt             time.Time
	updatedAt             time.Time
}

func (instance *Deputy) NewUpdater() *builder {
	return &builder{deputy: instance}
}

func (instance *Deputy) Id() uuid.UUID {
	return instance.id
}

func (instance *Deputy) Code() int {
	return instance.code
}

func (instance *Deputy) Cpf() string {
	return instance.cpf
}

func (instance *Deputy) Name() string {
	return instance.name
}

func (instance *Deputy) ElectoralName() string {
	return instance.electoralName
}

func (instance *Deputy) ImageUrl() string {
	return instance.imageUrl
}

func (instance *Deputy) Party() party.Party {
	return instance.party
}

func (instance *Deputy) PartyInTheProposition() party.Party {
	return instance.partyInTheProposition
}

func (instance *Deputy) Active() bool {
	return instance.active
}

func (instance *Deputy) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *Deputy) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *Deputy) IsEqual(deputy Deputy) bool {
	return instance.code == deputy.code &&
		instance.cpf == deputy.cpf &&
		instance.name == deputy.name &&
		instance.electoralName == deputy.electoralName &&
		instance.imageUrl == deputy.imageUrl &&
		instance.party.IsEqual(deputy.party)
}

func (instance *Deputy) IsZero() bool {
	return reflect.DeepEqual(instance, &Deputy{})
}
