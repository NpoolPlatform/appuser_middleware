// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuserthirdparty"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppUserThirdPartyUpdate is the builder for updating AppUserThirdParty entities.
type AppUserThirdPartyUpdate struct {
	config
	hooks     []Hook
	mutation  *AppUserThirdPartyMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AppUserThirdPartyUpdate builder.
func (autpu *AppUserThirdPartyUpdate) Where(ps ...predicate.AppUserThirdParty) *AppUserThirdPartyUpdate {
	autpu.mutation.Where(ps...)
	return autpu
}

// SetCreatedAt sets the "created_at" field.
func (autpu *AppUserThirdPartyUpdate) SetCreatedAt(u uint32) *AppUserThirdPartyUpdate {
	autpu.mutation.ResetCreatedAt()
	autpu.mutation.SetCreatedAt(u)
	return autpu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (autpu *AppUserThirdPartyUpdate) SetNillableCreatedAt(u *uint32) *AppUserThirdPartyUpdate {
	if u != nil {
		autpu.SetCreatedAt(*u)
	}
	return autpu
}

// AddCreatedAt adds u to the "created_at" field.
func (autpu *AppUserThirdPartyUpdate) AddCreatedAt(u int32) *AppUserThirdPartyUpdate {
	autpu.mutation.AddCreatedAt(u)
	return autpu
}

// SetUpdatedAt sets the "updated_at" field.
func (autpu *AppUserThirdPartyUpdate) SetUpdatedAt(u uint32) *AppUserThirdPartyUpdate {
	autpu.mutation.ResetUpdatedAt()
	autpu.mutation.SetUpdatedAt(u)
	return autpu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (autpu *AppUserThirdPartyUpdate) AddUpdatedAt(u int32) *AppUserThirdPartyUpdate {
	autpu.mutation.AddUpdatedAt(u)
	return autpu
}

// SetDeletedAt sets the "deleted_at" field.
func (autpu *AppUserThirdPartyUpdate) SetDeletedAt(u uint32) *AppUserThirdPartyUpdate {
	autpu.mutation.ResetDeletedAt()
	autpu.mutation.SetDeletedAt(u)
	return autpu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (autpu *AppUserThirdPartyUpdate) SetNillableDeletedAt(u *uint32) *AppUserThirdPartyUpdate {
	if u != nil {
		autpu.SetDeletedAt(*u)
	}
	return autpu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (autpu *AppUserThirdPartyUpdate) AddDeletedAt(u int32) *AppUserThirdPartyUpdate {
	autpu.mutation.AddDeletedAt(u)
	return autpu
}

// SetAppID sets the "app_id" field.
func (autpu *AppUserThirdPartyUpdate) SetAppID(u uuid.UUID) *AppUserThirdPartyUpdate {
	autpu.mutation.SetAppID(u)
	return autpu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (autpu *AppUserThirdPartyUpdate) SetNillableAppID(u *uuid.UUID) *AppUserThirdPartyUpdate {
	if u != nil {
		autpu.SetAppID(*u)
	}
	return autpu
}

// ClearAppID clears the value of the "app_id" field.
func (autpu *AppUserThirdPartyUpdate) ClearAppID() *AppUserThirdPartyUpdate {
	autpu.mutation.ClearAppID()
	return autpu
}

// SetUserID sets the "user_id" field.
func (autpu *AppUserThirdPartyUpdate) SetUserID(u uuid.UUID) *AppUserThirdPartyUpdate {
	autpu.mutation.SetUserID(u)
	return autpu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (autpu *AppUserThirdPartyUpdate) SetNillableUserID(u *uuid.UUID) *AppUserThirdPartyUpdate {
	if u != nil {
		autpu.SetUserID(*u)
	}
	return autpu
}

// ClearUserID clears the value of the "user_id" field.
func (autpu *AppUserThirdPartyUpdate) ClearUserID() *AppUserThirdPartyUpdate {
	autpu.mutation.ClearUserID()
	return autpu
}

// SetThirdPartyUserID sets the "third_party_user_id" field.
func (autpu *AppUserThirdPartyUpdate) SetThirdPartyUserID(s string) *AppUserThirdPartyUpdate {
	autpu.mutation.SetThirdPartyUserID(s)
	return autpu
}

// SetNillableThirdPartyUserID sets the "third_party_user_id" field if the given value is not nil.
func (autpu *AppUserThirdPartyUpdate) SetNillableThirdPartyUserID(s *string) *AppUserThirdPartyUpdate {
	if s != nil {
		autpu.SetThirdPartyUserID(*s)
	}
	return autpu
}

// ClearThirdPartyUserID clears the value of the "third_party_user_id" field.
func (autpu *AppUserThirdPartyUpdate) ClearThirdPartyUserID() *AppUserThirdPartyUpdate {
	autpu.mutation.ClearThirdPartyUserID()
	return autpu
}

// SetThirdPartyID sets the "third_party_id" field.
func (autpu *AppUserThirdPartyUpdate) SetThirdPartyID(u uuid.UUID) *AppUserThirdPartyUpdate {
	autpu.mutation.SetThirdPartyID(u)
	return autpu
}

// SetNillableThirdPartyID sets the "third_party_id" field if the given value is not nil.
func (autpu *AppUserThirdPartyUpdate) SetNillableThirdPartyID(u *uuid.UUID) *AppUserThirdPartyUpdate {
	if u != nil {
		autpu.SetThirdPartyID(*u)
	}
	return autpu
}

// ClearThirdPartyID clears the value of the "third_party_id" field.
func (autpu *AppUserThirdPartyUpdate) ClearThirdPartyID() *AppUserThirdPartyUpdate {
	autpu.mutation.ClearThirdPartyID()
	return autpu
}

// SetThirdPartyUsername sets the "third_party_username" field.
func (autpu *AppUserThirdPartyUpdate) SetThirdPartyUsername(s string) *AppUserThirdPartyUpdate {
	autpu.mutation.SetThirdPartyUsername(s)
	return autpu
}

// SetNillableThirdPartyUsername sets the "third_party_username" field if the given value is not nil.
func (autpu *AppUserThirdPartyUpdate) SetNillableThirdPartyUsername(s *string) *AppUserThirdPartyUpdate {
	if s != nil {
		autpu.SetThirdPartyUsername(*s)
	}
	return autpu
}

// ClearThirdPartyUsername clears the value of the "third_party_username" field.
func (autpu *AppUserThirdPartyUpdate) ClearThirdPartyUsername() *AppUserThirdPartyUpdate {
	autpu.mutation.ClearThirdPartyUsername()
	return autpu
}

// SetThirdPartyAvatar sets the "third_party_avatar" field.
func (autpu *AppUserThirdPartyUpdate) SetThirdPartyAvatar(s string) *AppUserThirdPartyUpdate {
	autpu.mutation.SetThirdPartyAvatar(s)
	return autpu
}

// SetNillableThirdPartyAvatar sets the "third_party_avatar" field if the given value is not nil.
func (autpu *AppUserThirdPartyUpdate) SetNillableThirdPartyAvatar(s *string) *AppUserThirdPartyUpdate {
	if s != nil {
		autpu.SetThirdPartyAvatar(*s)
	}
	return autpu
}

// ClearThirdPartyAvatar clears the value of the "third_party_avatar" field.
func (autpu *AppUserThirdPartyUpdate) ClearThirdPartyAvatar() *AppUserThirdPartyUpdate {
	autpu.mutation.ClearThirdPartyAvatar()
	return autpu
}

// Mutation returns the AppUserThirdPartyMutation object of the builder.
func (autpu *AppUserThirdPartyUpdate) Mutation() *AppUserThirdPartyMutation {
	return autpu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (autpu *AppUserThirdPartyUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := autpu.defaults(); err != nil {
		return 0, err
	}
	if len(autpu.hooks) == 0 {
		affected, err = autpu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserThirdPartyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			autpu.mutation = mutation
			affected, err = autpu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(autpu.hooks) - 1; i >= 0; i-- {
			if autpu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = autpu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, autpu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (autpu *AppUserThirdPartyUpdate) SaveX(ctx context.Context) int {
	affected, err := autpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (autpu *AppUserThirdPartyUpdate) Exec(ctx context.Context) error {
	_, err := autpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (autpu *AppUserThirdPartyUpdate) ExecX(ctx context.Context) {
	if err := autpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (autpu *AppUserThirdPartyUpdate) defaults() error {
	if _, ok := autpu.mutation.UpdatedAt(); !ok {
		if appuserthirdparty.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appuserthirdparty.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appuserthirdparty.UpdateDefaultUpdatedAt()
		autpu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (autpu *AppUserThirdPartyUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppUserThirdPartyUpdate {
	autpu.modifiers = append(autpu.modifiers, modifiers...)
	return autpu
}

func (autpu *AppUserThirdPartyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appuserthirdparty.Table,
			Columns: appuserthirdparty.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuserthirdparty.FieldID,
			},
		},
	}
	if ps := autpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := autpu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldCreatedAt,
		})
	}
	if value, ok := autpu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldCreatedAt,
		})
	}
	if value, ok := autpu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldUpdatedAt,
		})
	}
	if value, ok := autpu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldUpdatedAt,
		})
	}
	if value, ok := autpu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldDeletedAt,
		})
	}
	if value, ok := autpu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldDeletedAt,
		})
	}
	if value, ok := autpu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserthirdparty.FieldAppID,
		})
	}
	if autpu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appuserthirdparty.FieldAppID,
		})
	}
	if value, ok := autpu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserthirdparty.FieldUserID,
		})
	}
	if autpu.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appuserthirdparty.FieldUserID,
		})
	}
	if value, ok := autpu.mutation.ThirdPartyUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthirdparty.FieldThirdPartyUserID,
		})
	}
	if autpu.mutation.ThirdPartyUserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appuserthirdparty.FieldThirdPartyUserID,
		})
	}
	if value, ok := autpu.mutation.ThirdPartyID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserthirdparty.FieldThirdPartyID,
		})
	}
	if autpu.mutation.ThirdPartyIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appuserthirdparty.FieldThirdPartyID,
		})
	}
	if value, ok := autpu.mutation.ThirdPartyUsername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthirdparty.FieldThirdPartyUsername,
		})
	}
	if autpu.mutation.ThirdPartyUsernameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appuserthirdparty.FieldThirdPartyUsername,
		})
	}
	if value, ok := autpu.mutation.ThirdPartyAvatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthirdparty.FieldThirdPartyAvatar,
		})
	}
	if autpu.mutation.ThirdPartyAvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appuserthirdparty.FieldThirdPartyAvatar,
		})
	}
	_spec.Modifiers = autpu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, autpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appuserthirdparty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AppUserThirdPartyUpdateOne is the builder for updating a single AppUserThirdParty entity.
type AppUserThirdPartyUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AppUserThirdPartyMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (autpuo *AppUserThirdPartyUpdateOne) SetCreatedAt(u uint32) *AppUserThirdPartyUpdateOne {
	autpuo.mutation.ResetCreatedAt()
	autpuo.mutation.SetCreatedAt(u)
	return autpuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (autpuo *AppUserThirdPartyUpdateOne) SetNillableCreatedAt(u *uint32) *AppUserThirdPartyUpdateOne {
	if u != nil {
		autpuo.SetCreatedAt(*u)
	}
	return autpuo
}

// AddCreatedAt adds u to the "created_at" field.
func (autpuo *AppUserThirdPartyUpdateOne) AddCreatedAt(u int32) *AppUserThirdPartyUpdateOne {
	autpuo.mutation.AddCreatedAt(u)
	return autpuo
}

// SetUpdatedAt sets the "updated_at" field.
func (autpuo *AppUserThirdPartyUpdateOne) SetUpdatedAt(u uint32) *AppUserThirdPartyUpdateOne {
	autpuo.mutation.ResetUpdatedAt()
	autpuo.mutation.SetUpdatedAt(u)
	return autpuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (autpuo *AppUserThirdPartyUpdateOne) AddUpdatedAt(u int32) *AppUserThirdPartyUpdateOne {
	autpuo.mutation.AddUpdatedAt(u)
	return autpuo
}

// SetDeletedAt sets the "deleted_at" field.
func (autpuo *AppUserThirdPartyUpdateOne) SetDeletedAt(u uint32) *AppUserThirdPartyUpdateOne {
	autpuo.mutation.ResetDeletedAt()
	autpuo.mutation.SetDeletedAt(u)
	return autpuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (autpuo *AppUserThirdPartyUpdateOne) SetNillableDeletedAt(u *uint32) *AppUserThirdPartyUpdateOne {
	if u != nil {
		autpuo.SetDeletedAt(*u)
	}
	return autpuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (autpuo *AppUserThirdPartyUpdateOne) AddDeletedAt(u int32) *AppUserThirdPartyUpdateOne {
	autpuo.mutation.AddDeletedAt(u)
	return autpuo
}

// SetAppID sets the "app_id" field.
func (autpuo *AppUserThirdPartyUpdateOne) SetAppID(u uuid.UUID) *AppUserThirdPartyUpdateOne {
	autpuo.mutation.SetAppID(u)
	return autpuo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (autpuo *AppUserThirdPartyUpdateOne) SetNillableAppID(u *uuid.UUID) *AppUserThirdPartyUpdateOne {
	if u != nil {
		autpuo.SetAppID(*u)
	}
	return autpuo
}

// ClearAppID clears the value of the "app_id" field.
func (autpuo *AppUserThirdPartyUpdateOne) ClearAppID() *AppUserThirdPartyUpdateOne {
	autpuo.mutation.ClearAppID()
	return autpuo
}

// SetUserID sets the "user_id" field.
func (autpuo *AppUserThirdPartyUpdateOne) SetUserID(u uuid.UUID) *AppUserThirdPartyUpdateOne {
	autpuo.mutation.SetUserID(u)
	return autpuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (autpuo *AppUserThirdPartyUpdateOne) SetNillableUserID(u *uuid.UUID) *AppUserThirdPartyUpdateOne {
	if u != nil {
		autpuo.SetUserID(*u)
	}
	return autpuo
}

// ClearUserID clears the value of the "user_id" field.
func (autpuo *AppUserThirdPartyUpdateOne) ClearUserID() *AppUserThirdPartyUpdateOne {
	autpuo.mutation.ClearUserID()
	return autpuo
}

// SetThirdPartyUserID sets the "third_party_user_id" field.
func (autpuo *AppUserThirdPartyUpdateOne) SetThirdPartyUserID(s string) *AppUserThirdPartyUpdateOne {
	autpuo.mutation.SetThirdPartyUserID(s)
	return autpuo
}

// SetNillableThirdPartyUserID sets the "third_party_user_id" field if the given value is not nil.
func (autpuo *AppUserThirdPartyUpdateOne) SetNillableThirdPartyUserID(s *string) *AppUserThirdPartyUpdateOne {
	if s != nil {
		autpuo.SetThirdPartyUserID(*s)
	}
	return autpuo
}

// ClearThirdPartyUserID clears the value of the "third_party_user_id" field.
func (autpuo *AppUserThirdPartyUpdateOne) ClearThirdPartyUserID() *AppUserThirdPartyUpdateOne {
	autpuo.mutation.ClearThirdPartyUserID()
	return autpuo
}

// SetThirdPartyID sets the "third_party_id" field.
func (autpuo *AppUserThirdPartyUpdateOne) SetThirdPartyID(u uuid.UUID) *AppUserThirdPartyUpdateOne {
	autpuo.mutation.SetThirdPartyID(u)
	return autpuo
}

// SetNillableThirdPartyID sets the "third_party_id" field if the given value is not nil.
func (autpuo *AppUserThirdPartyUpdateOne) SetNillableThirdPartyID(u *uuid.UUID) *AppUserThirdPartyUpdateOne {
	if u != nil {
		autpuo.SetThirdPartyID(*u)
	}
	return autpuo
}

// ClearThirdPartyID clears the value of the "third_party_id" field.
func (autpuo *AppUserThirdPartyUpdateOne) ClearThirdPartyID() *AppUserThirdPartyUpdateOne {
	autpuo.mutation.ClearThirdPartyID()
	return autpuo
}

// SetThirdPartyUsername sets the "third_party_username" field.
func (autpuo *AppUserThirdPartyUpdateOne) SetThirdPartyUsername(s string) *AppUserThirdPartyUpdateOne {
	autpuo.mutation.SetThirdPartyUsername(s)
	return autpuo
}

// SetNillableThirdPartyUsername sets the "third_party_username" field if the given value is not nil.
func (autpuo *AppUserThirdPartyUpdateOne) SetNillableThirdPartyUsername(s *string) *AppUserThirdPartyUpdateOne {
	if s != nil {
		autpuo.SetThirdPartyUsername(*s)
	}
	return autpuo
}

// ClearThirdPartyUsername clears the value of the "third_party_username" field.
func (autpuo *AppUserThirdPartyUpdateOne) ClearThirdPartyUsername() *AppUserThirdPartyUpdateOne {
	autpuo.mutation.ClearThirdPartyUsername()
	return autpuo
}

// SetThirdPartyAvatar sets the "third_party_avatar" field.
func (autpuo *AppUserThirdPartyUpdateOne) SetThirdPartyAvatar(s string) *AppUserThirdPartyUpdateOne {
	autpuo.mutation.SetThirdPartyAvatar(s)
	return autpuo
}

// SetNillableThirdPartyAvatar sets the "third_party_avatar" field if the given value is not nil.
func (autpuo *AppUserThirdPartyUpdateOne) SetNillableThirdPartyAvatar(s *string) *AppUserThirdPartyUpdateOne {
	if s != nil {
		autpuo.SetThirdPartyAvatar(*s)
	}
	return autpuo
}

// ClearThirdPartyAvatar clears the value of the "third_party_avatar" field.
func (autpuo *AppUserThirdPartyUpdateOne) ClearThirdPartyAvatar() *AppUserThirdPartyUpdateOne {
	autpuo.mutation.ClearThirdPartyAvatar()
	return autpuo
}

// Mutation returns the AppUserThirdPartyMutation object of the builder.
func (autpuo *AppUserThirdPartyUpdateOne) Mutation() *AppUserThirdPartyMutation {
	return autpuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (autpuo *AppUserThirdPartyUpdateOne) Select(field string, fields ...string) *AppUserThirdPartyUpdateOne {
	autpuo.fields = append([]string{field}, fields...)
	return autpuo
}

// Save executes the query and returns the updated AppUserThirdParty entity.
func (autpuo *AppUserThirdPartyUpdateOne) Save(ctx context.Context) (*AppUserThirdParty, error) {
	var (
		err  error
		node *AppUserThirdParty
	)
	if err := autpuo.defaults(); err != nil {
		return nil, err
	}
	if len(autpuo.hooks) == 0 {
		node, err = autpuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserThirdPartyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			autpuo.mutation = mutation
			node, err = autpuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(autpuo.hooks) - 1; i >= 0; i-- {
			if autpuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = autpuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, autpuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppUserThirdParty)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppUserThirdPartyMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (autpuo *AppUserThirdPartyUpdateOne) SaveX(ctx context.Context) *AppUserThirdParty {
	node, err := autpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (autpuo *AppUserThirdPartyUpdateOne) Exec(ctx context.Context) error {
	_, err := autpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (autpuo *AppUserThirdPartyUpdateOne) ExecX(ctx context.Context) {
	if err := autpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (autpuo *AppUserThirdPartyUpdateOne) defaults() error {
	if _, ok := autpuo.mutation.UpdatedAt(); !ok {
		if appuserthirdparty.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appuserthirdparty.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appuserthirdparty.UpdateDefaultUpdatedAt()
		autpuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (autpuo *AppUserThirdPartyUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppUserThirdPartyUpdateOne {
	autpuo.modifiers = append(autpuo.modifiers, modifiers...)
	return autpuo
}

func (autpuo *AppUserThirdPartyUpdateOne) sqlSave(ctx context.Context) (_node *AppUserThirdParty, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appuserthirdparty.Table,
			Columns: appuserthirdparty.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuserthirdparty.FieldID,
			},
		},
	}
	id, ok := autpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppUserThirdParty.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := autpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appuserthirdparty.FieldID)
		for _, f := range fields {
			if !appuserthirdparty.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appuserthirdparty.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := autpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := autpuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldCreatedAt,
		})
	}
	if value, ok := autpuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldCreatedAt,
		})
	}
	if value, ok := autpuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldUpdatedAt,
		})
	}
	if value, ok := autpuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldUpdatedAt,
		})
	}
	if value, ok := autpuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldDeletedAt,
		})
	}
	if value, ok := autpuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserthirdparty.FieldDeletedAt,
		})
	}
	if value, ok := autpuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserthirdparty.FieldAppID,
		})
	}
	if autpuo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appuserthirdparty.FieldAppID,
		})
	}
	if value, ok := autpuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserthirdparty.FieldUserID,
		})
	}
	if autpuo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appuserthirdparty.FieldUserID,
		})
	}
	if value, ok := autpuo.mutation.ThirdPartyUserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthirdparty.FieldThirdPartyUserID,
		})
	}
	if autpuo.mutation.ThirdPartyUserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appuserthirdparty.FieldThirdPartyUserID,
		})
	}
	if value, ok := autpuo.mutation.ThirdPartyID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserthirdparty.FieldThirdPartyID,
		})
	}
	if autpuo.mutation.ThirdPartyIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appuserthirdparty.FieldThirdPartyID,
		})
	}
	if value, ok := autpuo.mutation.ThirdPartyUsername(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthirdparty.FieldThirdPartyUsername,
		})
	}
	if autpuo.mutation.ThirdPartyUsernameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appuserthirdparty.FieldThirdPartyUsername,
		})
	}
	if value, ok := autpuo.mutation.ThirdPartyAvatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserthirdparty.FieldThirdPartyAvatar,
		})
	}
	if autpuo.mutation.ThirdPartyAvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appuserthirdparty.FieldThirdPartyAvatar,
		})
	}
	_spec.Modifiers = autpuo.modifiers
	_node = &AppUserThirdParty{config: autpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, autpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appuserthirdparty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
