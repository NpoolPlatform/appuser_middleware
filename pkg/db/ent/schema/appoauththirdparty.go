package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
)

// AppOAuthThirdParty holds the schema definition for the AppOAuthThirdParty entity.
type AppOAuthThirdParty struct {
	ent.Schema
}

func (AppOAuthThirdParty) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppOAuthThirdParty.
func (AppOAuthThirdParty) Fields() []ent.Field {
	return []ent.Field{
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
		field.
			String("client_id").
			Optional().
			Default(""),
		field.
			String("client_secret").
			Optional().
			Default(""),
		field.
			String("callback_url").
			Optional().
			Default(""),
		field.
			String("salt").
			Optional().
			Default(""),
	}
}

// Edges of the AppOAuthThirdParty.
func (AppOAuthThirdParty) Edges() []ent.Edge {
	return nil
}
