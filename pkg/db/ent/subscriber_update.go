// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/subscriber"
	"github.com/google/uuid"
)

// SubscriberUpdate is the builder for updating Subscriber entities.
type SubscriberUpdate struct {
	config
	hooks    []Hook
	mutation *SubscriberMutation
}

// Where appends a list predicates to the SubscriberUpdate builder.
func (su *SubscriberUpdate) Where(ps ...predicate.Subscriber) *SubscriberUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *SubscriberUpdate) SetCreatedAt(u uint32) *SubscriberUpdate {
	su.mutation.ResetCreatedAt()
	su.mutation.SetCreatedAt(u)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *SubscriberUpdate) SetNillableCreatedAt(u *uint32) *SubscriberUpdate {
	if u != nil {
		su.SetCreatedAt(*u)
	}
	return su
}

// AddCreatedAt adds u to the "created_at" field.
func (su *SubscriberUpdate) AddCreatedAt(u int32) *SubscriberUpdate {
	su.mutation.AddCreatedAt(u)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SubscriberUpdate) SetUpdatedAt(u uint32) *SubscriberUpdate {
	su.mutation.ResetUpdatedAt()
	su.mutation.SetUpdatedAt(u)
	return su
}

// AddUpdatedAt adds u to the "updated_at" field.
func (su *SubscriberUpdate) AddUpdatedAt(u int32) *SubscriberUpdate {
	su.mutation.AddUpdatedAt(u)
	return su
}

// SetDeletedAt sets the "deleted_at" field.
func (su *SubscriberUpdate) SetDeletedAt(u uint32) *SubscriberUpdate {
	su.mutation.ResetDeletedAt()
	su.mutation.SetDeletedAt(u)
	return su
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (su *SubscriberUpdate) SetNillableDeletedAt(u *uint32) *SubscriberUpdate {
	if u != nil {
		su.SetDeletedAt(*u)
	}
	return su
}

// AddDeletedAt adds u to the "deleted_at" field.
func (su *SubscriberUpdate) AddDeletedAt(u int32) *SubscriberUpdate {
	su.mutation.AddDeletedAt(u)
	return su
}

// SetAppID sets the "app_id" field.
func (su *SubscriberUpdate) SetAppID(u uuid.UUID) *SubscriberUpdate {
	su.mutation.SetAppID(u)
	return su
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (su *SubscriberUpdate) SetNillableAppID(u *uuid.UUID) *SubscriberUpdate {
	if u != nil {
		su.SetAppID(*u)
	}
	return su
}

// ClearAppID clears the value of the "app_id" field.
func (su *SubscriberUpdate) ClearAppID() *SubscriberUpdate {
	su.mutation.ClearAppID()
	return su
}

// SetEmailAddress sets the "email_address" field.
func (su *SubscriberUpdate) SetEmailAddress(s string) *SubscriberUpdate {
	su.mutation.SetEmailAddress(s)
	return su
}

// SetNillableEmailAddress sets the "email_address" field if the given value is not nil.
func (su *SubscriberUpdate) SetNillableEmailAddress(s *string) *SubscriberUpdate {
	if s != nil {
		su.SetEmailAddress(*s)
	}
	return su
}

// ClearEmailAddress clears the value of the "email_address" field.
func (su *SubscriberUpdate) ClearEmailAddress() *SubscriberUpdate {
	su.mutation.ClearEmailAddress()
	return su
}

// SetRegistered sets the "registered" field.
func (su *SubscriberUpdate) SetRegistered(b bool) *SubscriberUpdate {
	su.mutation.SetRegistered(b)
	return su
}

// SetNillableRegistered sets the "registered" field if the given value is not nil.
func (su *SubscriberUpdate) SetNillableRegistered(b *bool) *SubscriberUpdate {
	if b != nil {
		su.SetRegistered(*b)
	}
	return su
}

// ClearRegistered clears the value of the "registered" field.
func (su *SubscriberUpdate) ClearRegistered() *SubscriberUpdate {
	su.mutation.ClearRegistered()
	return su
}

// Mutation returns the SubscriberMutation object of the builder.
func (su *SubscriberUpdate) Mutation() *SubscriberMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SubscriberUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := su.defaults(); err != nil {
		return 0, err
	}
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubscriberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SubscriberUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SubscriberUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SubscriberUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SubscriberUpdate) defaults() error {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		if subscriber.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized subscriber.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := subscriber.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (su *SubscriberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subscriber.Table,
			Columns: subscriber.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: subscriber.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscriber.FieldCreatedAt,
		})
	}
	if value, ok := su.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscriber.FieldCreatedAt,
		})
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscriber.FieldUpdatedAt,
		})
	}
	if value, ok := su.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscriber.FieldUpdatedAt,
		})
	}
	if value, ok := su.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscriber.FieldDeletedAt,
		})
	}
	if value, ok := su.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscriber.FieldDeletedAt,
		})
	}
	if value, ok := su.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: subscriber.FieldAppID,
		})
	}
	if su.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: subscriber.FieldAppID,
		})
	}
	if value, ok := su.mutation.EmailAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscriber.FieldEmailAddress,
		})
	}
	if su.mutation.EmailAddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: subscriber.FieldEmailAddress,
		})
	}
	if value, ok := su.mutation.Registered(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: subscriber.FieldRegistered,
		})
	}
	if su.mutation.RegisteredCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: subscriber.FieldRegistered,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscriber.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SubscriberUpdateOne is the builder for updating a single Subscriber entity.
type SubscriberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SubscriberMutation
}

// SetCreatedAt sets the "created_at" field.
func (suo *SubscriberUpdateOne) SetCreatedAt(u uint32) *SubscriberUpdateOne {
	suo.mutation.ResetCreatedAt()
	suo.mutation.SetCreatedAt(u)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *SubscriberUpdateOne) SetNillableCreatedAt(u *uint32) *SubscriberUpdateOne {
	if u != nil {
		suo.SetCreatedAt(*u)
	}
	return suo
}

// AddCreatedAt adds u to the "created_at" field.
func (suo *SubscriberUpdateOne) AddCreatedAt(u int32) *SubscriberUpdateOne {
	suo.mutation.AddCreatedAt(u)
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SubscriberUpdateOne) SetUpdatedAt(u uint32) *SubscriberUpdateOne {
	suo.mutation.ResetUpdatedAt()
	suo.mutation.SetUpdatedAt(u)
	return suo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (suo *SubscriberUpdateOne) AddUpdatedAt(u int32) *SubscriberUpdateOne {
	suo.mutation.AddUpdatedAt(u)
	return suo
}

// SetDeletedAt sets the "deleted_at" field.
func (suo *SubscriberUpdateOne) SetDeletedAt(u uint32) *SubscriberUpdateOne {
	suo.mutation.ResetDeletedAt()
	suo.mutation.SetDeletedAt(u)
	return suo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suo *SubscriberUpdateOne) SetNillableDeletedAt(u *uint32) *SubscriberUpdateOne {
	if u != nil {
		suo.SetDeletedAt(*u)
	}
	return suo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (suo *SubscriberUpdateOne) AddDeletedAt(u int32) *SubscriberUpdateOne {
	suo.mutation.AddDeletedAt(u)
	return suo
}

// SetAppID sets the "app_id" field.
func (suo *SubscriberUpdateOne) SetAppID(u uuid.UUID) *SubscriberUpdateOne {
	suo.mutation.SetAppID(u)
	return suo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (suo *SubscriberUpdateOne) SetNillableAppID(u *uuid.UUID) *SubscriberUpdateOne {
	if u != nil {
		suo.SetAppID(*u)
	}
	return suo
}

// ClearAppID clears the value of the "app_id" field.
func (suo *SubscriberUpdateOne) ClearAppID() *SubscriberUpdateOne {
	suo.mutation.ClearAppID()
	return suo
}

// SetEmailAddress sets the "email_address" field.
func (suo *SubscriberUpdateOne) SetEmailAddress(s string) *SubscriberUpdateOne {
	suo.mutation.SetEmailAddress(s)
	return suo
}

// SetNillableEmailAddress sets the "email_address" field if the given value is not nil.
func (suo *SubscriberUpdateOne) SetNillableEmailAddress(s *string) *SubscriberUpdateOne {
	if s != nil {
		suo.SetEmailAddress(*s)
	}
	return suo
}

// ClearEmailAddress clears the value of the "email_address" field.
func (suo *SubscriberUpdateOne) ClearEmailAddress() *SubscriberUpdateOne {
	suo.mutation.ClearEmailAddress()
	return suo
}

// SetRegistered sets the "registered" field.
func (suo *SubscriberUpdateOne) SetRegistered(b bool) *SubscriberUpdateOne {
	suo.mutation.SetRegistered(b)
	return suo
}

// SetNillableRegistered sets the "registered" field if the given value is not nil.
func (suo *SubscriberUpdateOne) SetNillableRegistered(b *bool) *SubscriberUpdateOne {
	if b != nil {
		suo.SetRegistered(*b)
	}
	return suo
}

// ClearRegistered clears the value of the "registered" field.
func (suo *SubscriberUpdateOne) ClearRegistered() *SubscriberUpdateOne {
	suo.mutation.ClearRegistered()
	return suo
}

// Mutation returns the SubscriberMutation object of the builder.
func (suo *SubscriberUpdateOne) Mutation() *SubscriberMutation {
	return suo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SubscriberUpdateOne) Select(field string, fields ...string) *SubscriberUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Subscriber entity.
func (suo *SubscriberUpdateOne) Save(ctx context.Context) (*Subscriber, error) {
	var (
		err  error
		node *Subscriber
	)
	if err := suo.defaults(); err != nil {
		return nil, err
	}
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubscriberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Subscriber)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SubscriberMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SubscriberUpdateOne) SaveX(ctx context.Context) *Subscriber {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SubscriberUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SubscriberUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SubscriberUpdateOne) defaults() error {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		if subscriber.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized subscriber.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := subscriber.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (suo *SubscriberUpdateOne) sqlSave(ctx context.Context) (_node *Subscriber, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   subscriber.Table,
			Columns: subscriber.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: subscriber.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Subscriber.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subscriber.FieldID)
		for _, f := range fields {
			if !subscriber.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != subscriber.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscriber.FieldCreatedAt,
		})
	}
	if value, ok := suo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscriber.FieldCreatedAt,
		})
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscriber.FieldUpdatedAt,
		})
	}
	if value, ok := suo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscriber.FieldUpdatedAt,
		})
	}
	if value, ok := suo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscriber.FieldDeletedAt,
		})
	}
	if value, ok := suo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscriber.FieldDeletedAt,
		})
	}
	if value, ok := suo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: subscriber.FieldAppID,
		})
	}
	if suo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: subscriber.FieldAppID,
		})
	}
	if value, ok := suo.mutation.EmailAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscriber.FieldEmailAddress,
		})
	}
	if suo.mutation.EmailAddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: subscriber.FieldEmailAddress,
		})
	}
	if value, ok := suo.mutation.Registered(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: subscriber.FieldRegistered,
		})
	}
	if suo.mutation.RegisteredCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: subscriber.FieldRegistered,
		})
	}
	_node = &Subscriber{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{subscriber.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
