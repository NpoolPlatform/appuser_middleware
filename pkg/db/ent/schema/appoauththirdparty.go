package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// AppOAuthThirdParty holds the schema definition for the AppOAuthThirdParty entity.
type AppOAuthThirdParty struct {
	ent.Schema
}

func (AppOAuthThirdParty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the AppOAuthThirdParty.
func (AppOAuthThirdParty) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("third_party_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
	}
}

// Edges of the AppOAuthThirdParty.
func (AppOAuthThirdParty) Edges() []ent.Edge {
	return nil
}
