package event

import (
	"time"

	"github.com/bluesky-social/indigo/atproto/identity"
)

// information about a repo/account/identity, always pre-populated and relevant to many rules
type AccountMeta struct {
	Identity             *identity.Identity
	Profile              ProfileSummary
	Private              *AccountPrivate
	AccountLabels        []string
	AccountNegatedLabels []string
	AccountFlags         []string
	FollowersCount       int64
	FollowsCount         int64
	PostsCount           int64
	Takendown            bool
}

type ProfileSummary struct {
	HasAvatar   bool
	Description *string
	DisplayName *string
}

type AccountPrivate struct {
	Email          string
	EmailConfirmed bool
	IndexedAt      time.Time
}
