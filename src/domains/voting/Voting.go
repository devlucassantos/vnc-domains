package voting

import (
	"github.com/devlucassantos/vnc-domains/src/domains/article"
	"github.com/devlucassantos/vnc-domains/src/domains/legislativebody"
	"github.com/devlucassantos/vnc-domains/src/domains/proposition"
	"github.com/google/uuid"
	"reflect"
	"time"
)

type Voting struct {
	id                   uuid.UUID
	code                 string
	title                string
	description          string
	result               string
	resultAnnouncedAt    time.Time
	isApproved           *bool
	legislativeBody      legislativebody.LegislativeBody
	mainProposition      proposition.Proposition
	relatedPropositions  []proposition.Proposition
	affectedPropositions []proposition.Proposition
	article              article.Article
	relatedArticles      []article.Article
	active               bool
	createdAt            time.Time
	updatedAt            time.Time
}

func (instance *Voting) NewUpdater() *builder {
	return &builder{voting: instance}
}

func (instance *Voting) Id() uuid.UUID {
	return instance.id
}

func (instance *Voting) Code() string {
	return instance.code
}

func (instance *Voting) Title() string {
	return instance.title
}

func (instance *Voting) Description() string {
	return instance.description
}

func (instance *Voting) Result() string {
	return instance.result
}

func (instance *Voting) ResultAnnouncedAt() time.Time {
	return instance.resultAnnouncedAt
}

func (instance *Voting) IsApproved() *bool {
	return instance.isApproved
}

func (instance *Voting) LegislativeBody() legislativebody.LegislativeBody {
	return instance.legislativeBody
}

func (instance *Voting) MainProposition() proposition.Proposition {
	return instance.mainProposition
}

func (instance *Voting) RelatedPropositions() []proposition.Proposition {
	return instance.relatedPropositions
}

func (instance *Voting) AffectedPropositions() []proposition.Proposition {
	return instance.affectedPropositions
}

func (instance *Voting) Article() article.Article {
	return instance.article
}

func (instance *Voting) RelatedArticles() []article.Article {
	return instance.relatedArticles
}

func (instance *Voting) Active() bool {
	return instance.active
}

func (instance *Voting) CreatedAt() time.Time {
	return instance.createdAt
}

func (instance *Voting) UpdatedAt() time.Time {
	return instance.updatedAt
}

func (instance *Voting) IsZero() bool {
	return reflect.DeepEqual(instance, &Voting{})
}
