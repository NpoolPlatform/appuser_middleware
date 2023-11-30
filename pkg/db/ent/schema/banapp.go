package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
)

// BanApp holds the schema definition for the BanApp entity.
type BanApp struct {
	ent.Schema
}

func (BanApp) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the BanApp.
func (BanApp) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("app_id", uuid.UUID{}),
		field.String("message").
			Default(""),
	}
}

// Edges of the BanApp.
func (BanApp) Edges() []ent.Edge {
	return nil
}
