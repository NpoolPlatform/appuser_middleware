// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appoauththirdparty"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/predicate"
)

// AppOAuthThirdPartyDelete is the builder for deleting a AppOAuthThirdParty entity.
type AppOAuthThirdPartyDelete struct {
	config
	hooks    []Hook
	mutation *AppOAuthThirdPartyMutation
}

// Where appends a list predicates to the AppOAuthThirdPartyDelete builder.
func (aotpd *AppOAuthThirdPartyDelete) Where(ps ...predicate.AppOAuthThirdParty) *AppOAuthThirdPartyDelete {
	aotpd.mutation.Where(ps...)
	return aotpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (aotpd *AppOAuthThirdPartyDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(aotpd.hooks) == 0 {
		affected, err = aotpd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppOAuthThirdPartyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aotpd.mutation = mutation
			affected, err = aotpd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aotpd.hooks) - 1; i >= 0; i-- {
			if aotpd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aotpd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aotpd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (aotpd *AppOAuthThirdPartyDelete) ExecX(ctx context.Context) int {
	n, err := aotpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (aotpd *AppOAuthThirdPartyDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: appoauththirdparty.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appoauththirdparty.FieldID,
			},
		},
	}
	if ps := aotpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, aotpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AppOAuthThirdPartyDeleteOne is the builder for deleting a single AppOAuthThirdParty entity.
type AppOAuthThirdPartyDeleteOne struct {
	aotpd *AppOAuthThirdPartyDelete
}

// Exec executes the deletion query.
func (aotpdo *AppOAuthThirdPartyDeleteOne) Exec(ctx context.Context) error {
	n, err := aotpdo.aotpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appoauththirdparty.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (aotpdo *AppOAuthThirdPartyDeleteOne) ExecX(ctx context.Context) {
	aotpdo.aotpd.ExecX(ctx)
}
