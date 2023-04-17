// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/subscriber"
	"github.com/google/uuid"
)

// SubscriberCreate is the builder for creating a Subscriber entity.
type SubscriberCreate struct {
	config
	mutation *SubscriberMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (sc *SubscriberCreate) SetCreatedAt(u uint32) *SubscriberCreate {
	sc.mutation.SetCreatedAt(u)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableCreatedAt(u *uint32) *SubscriberCreate {
	if u != nil {
		sc.SetCreatedAt(*u)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SubscriberCreate) SetUpdatedAt(u uint32) *SubscriberCreate {
	sc.mutation.SetUpdatedAt(u)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableUpdatedAt(u *uint32) *SubscriberCreate {
	if u != nil {
		sc.SetUpdatedAt(*u)
	}
	return sc
}

// SetDeletedAt sets the "deleted_at" field.
func (sc *SubscriberCreate) SetDeletedAt(u uint32) *SubscriberCreate {
	sc.mutation.SetDeletedAt(u)
	return sc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableDeletedAt(u *uint32) *SubscriberCreate {
	if u != nil {
		sc.SetDeletedAt(*u)
	}
	return sc
}

// SetAppID sets the "app_id" field.
func (sc *SubscriberCreate) SetAppID(u uuid.UUID) *SubscriberCreate {
	sc.mutation.SetAppID(u)
	return sc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableAppID(u *uuid.UUID) *SubscriberCreate {
	if u != nil {
		sc.SetAppID(*u)
	}
	return sc
}

// SetEmailAddress sets the "email_address" field.
func (sc *SubscriberCreate) SetEmailAddress(s string) *SubscriberCreate {
	sc.mutation.SetEmailAddress(s)
	return sc
}

// SetNillableEmailAddress sets the "email_address" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableEmailAddress(s *string) *SubscriberCreate {
	if s != nil {
		sc.SetEmailAddress(*s)
	}
	return sc
}

// SetRegistered sets the "registered" field.
func (sc *SubscriberCreate) SetRegistered(b bool) *SubscriberCreate {
	sc.mutation.SetRegistered(b)
	return sc
}

// SetNillableRegistered sets the "registered" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableRegistered(b *bool) *SubscriberCreate {
	if b != nil {
		sc.SetRegistered(*b)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SubscriberCreate) SetID(u uuid.UUID) *SubscriberCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SubscriberCreate) SetNillableID(u *uuid.UUID) *SubscriberCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// Mutation returns the SubscriberMutation object of the builder.
func (sc *SubscriberCreate) Mutation() *SubscriberMutation {
	return sc.mutation
}

// Save creates the Subscriber in the database.
func (sc *SubscriberCreate) Save(ctx context.Context) (*Subscriber, error) {
	var (
		err  error
		node *Subscriber
	)
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SubscriberMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (sc *SubscriberCreate) SaveX(ctx context.Context) *Subscriber {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubscriberCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubscriberCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SubscriberCreate) defaults() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		if subscriber.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized subscriber.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := subscriber.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		if subscriber.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized subscriber.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := subscriber.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.DeletedAt(); !ok {
		if subscriber.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized subscriber.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := subscriber.DefaultDeletedAt()
		sc.mutation.SetDeletedAt(v)
	}
	if _, ok := sc.mutation.AppID(); !ok {
		if subscriber.DefaultAppID == nil {
			return fmt.Errorf("ent: uninitialized subscriber.DefaultAppID (forgotten import ent/runtime?)")
		}
		v := subscriber.DefaultAppID()
		sc.mutation.SetAppID(v)
	}
	if _, ok := sc.mutation.EmailAddress(); !ok {
		v := subscriber.DefaultEmailAddress
		sc.mutation.SetEmailAddress(v)
	}
	if _, ok := sc.mutation.Registered(); !ok {
		v := subscriber.DefaultRegistered
		sc.mutation.SetRegistered(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		if subscriber.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized subscriber.DefaultID (forgotten import ent/runtime?)")
		}
		v := subscriber.DefaultID()
		sc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubscriberCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Subscriber.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Subscriber.updated_at"`)}
	}
	if _, ok := sc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Subscriber.deleted_at"`)}
	}
	return nil
}

func (sc *SubscriberCreate) sqlSave(ctx context.Context) (*Subscriber, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (sc *SubscriberCreate) createSpec() (*Subscriber, *sqlgraph.CreateSpec) {
	var (
		_node = &Subscriber{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: subscriber.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: subscriber.FieldID,
			},
		}
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscriber.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscriber.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: subscriber.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := sc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: subscriber.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := sc.mutation.EmailAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: subscriber.FieldEmailAddress,
		})
		_node.EmailAddress = value
	}
	if value, ok := sc.mutation.Registered(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: subscriber.FieldRegistered,
		})
		_node.Registered = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Subscriber.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SubscriberUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (sc *SubscriberCreate) OnConflict(opts ...sql.ConflictOption) *SubscriberUpsertOne {
	sc.conflict = opts
	return &SubscriberUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Subscriber.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (sc *SubscriberCreate) OnConflictColumns(columns ...string) *SubscriberUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SubscriberUpsertOne{
		create: sc,
	}
}

type (
	// SubscriberUpsertOne is the builder for "upsert"-ing
	//  one Subscriber node.
	SubscriberUpsertOne struct {
		create *SubscriberCreate
	}

	// SubscriberUpsert is the "OnConflict" setter.
	SubscriberUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *SubscriberUpsert) SetCreatedAt(v uint32) *SubscriberUpsert {
	u.Set(subscriber.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SubscriberUpsert) UpdateCreatedAt() *SubscriberUpsert {
	u.SetExcluded(subscriber.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *SubscriberUpsert) AddCreatedAt(v uint32) *SubscriberUpsert {
	u.Add(subscriber.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SubscriberUpsert) SetUpdatedAt(v uint32) *SubscriberUpsert {
	u.Set(subscriber.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SubscriberUpsert) UpdateUpdatedAt() *SubscriberUpsert {
	u.SetExcluded(subscriber.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *SubscriberUpsert) AddUpdatedAt(v uint32) *SubscriberUpsert {
	u.Add(subscriber.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SubscriberUpsert) SetDeletedAt(v uint32) *SubscriberUpsert {
	u.Set(subscriber.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SubscriberUpsert) UpdateDeletedAt() *SubscriberUpsert {
	u.SetExcluded(subscriber.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *SubscriberUpsert) AddDeletedAt(v uint32) *SubscriberUpsert {
	u.Add(subscriber.FieldDeletedAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *SubscriberUpsert) SetAppID(v uuid.UUID) *SubscriberUpsert {
	u.Set(subscriber.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *SubscriberUpsert) UpdateAppID() *SubscriberUpsert {
	u.SetExcluded(subscriber.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *SubscriberUpsert) ClearAppID() *SubscriberUpsert {
	u.SetNull(subscriber.FieldAppID)
	return u
}

// SetEmailAddress sets the "email_address" field.
func (u *SubscriberUpsert) SetEmailAddress(v string) *SubscriberUpsert {
	u.Set(subscriber.FieldEmailAddress, v)
	return u
}

// UpdateEmailAddress sets the "email_address" field to the value that was provided on create.
func (u *SubscriberUpsert) UpdateEmailAddress() *SubscriberUpsert {
	u.SetExcluded(subscriber.FieldEmailAddress)
	return u
}

// ClearEmailAddress clears the value of the "email_address" field.
func (u *SubscriberUpsert) ClearEmailAddress() *SubscriberUpsert {
	u.SetNull(subscriber.FieldEmailAddress)
	return u
}

// SetRegistered sets the "registered" field.
func (u *SubscriberUpsert) SetRegistered(v bool) *SubscriberUpsert {
	u.Set(subscriber.FieldRegistered, v)
	return u
}

// UpdateRegistered sets the "registered" field to the value that was provided on create.
func (u *SubscriberUpsert) UpdateRegistered() *SubscriberUpsert {
	u.SetExcluded(subscriber.FieldRegistered)
	return u
}

// ClearRegistered clears the value of the "registered" field.
func (u *SubscriberUpsert) ClearRegistered() *SubscriberUpsert {
	u.SetNull(subscriber.FieldRegistered)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Subscriber.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(subscriber.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *SubscriberUpsertOne) UpdateNewValues() *SubscriberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(subscriber.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Subscriber.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *SubscriberUpsertOne) Ignore() *SubscriberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SubscriberUpsertOne) DoNothing() *SubscriberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SubscriberCreate.OnConflict
// documentation for more info.
func (u *SubscriberUpsertOne) Update(set func(*SubscriberUpsert)) *SubscriberUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SubscriberUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *SubscriberUpsertOne) SetCreatedAt(v uint32) *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *SubscriberUpsertOne) AddCreatedAt(v uint32) *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SubscriberUpsertOne) UpdateCreatedAt() *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SubscriberUpsertOne) SetUpdatedAt(v uint32) *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *SubscriberUpsertOne) AddUpdatedAt(v uint32) *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SubscriberUpsertOne) UpdateUpdatedAt() *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SubscriberUpsertOne) SetDeletedAt(v uint32) *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *SubscriberUpsertOne) AddDeletedAt(v uint32) *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SubscriberUpsertOne) UpdateDeletedAt() *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *SubscriberUpsertOne) SetAppID(v uuid.UUID) *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *SubscriberUpsertOne) UpdateAppID() *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *SubscriberUpsertOne) ClearAppID() *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.ClearAppID()
	})
}

// SetEmailAddress sets the "email_address" field.
func (u *SubscriberUpsertOne) SetEmailAddress(v string) *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.SetEmailAddress(v)
	})
}

// UpdateEmailAddress sets the "email_address" field to the value that was provided on create.
func (u *SubscriberUpsertOne) UpdateEmailAddress() *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.UpdateEmailAddress()
	})
}

// ClearEmailAddress clears the value of the "email_address" field.
func (u *SubscriberUpsertOne) ClearEmailAddress() *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.ClearEmailAddress()
	})
}

// SetRegistered sets the "registered" field.
func (u *SubscriberUpsertOne) SetRegistered(v bool) *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.SetRegistered(v)
	})
}

// UpdateRegistered sets the "registered" field to the value that was provided on create.
func (u *SubscriberUpsertOne) UpdateRegistered() *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.UpdateRegistered()
	})
}

// ClearRegistered clears the value of the "registered" field.
func (u *SubscriberUpsertOne) ClearRegistered() *SubscriberUpsertOne {
	return u.Update(func(s *SubscriberUpsert) {
		s.ClearRegistered()
	})
}

// Exec executes the query.
func (u *SubscriberUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SubscriberCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SubscriberUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SubscriberUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: SubscriberUpsertOne.ID is not supported by MySQL driver. Use SubscriberUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SubscriberUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SubscriberCreateBulk is the builder for creating many Subscriber entities in bulk.
type SubscriberCreateBulk struct {
	config
	builders []*SubscriberCreate
	conflict []sql.ConflictOption
}

// Save creates the Subscriber entities in the database.
func (scb *SubscriberCreateBulk) Save(ctx context.Context) ([]*Subscriber, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Subscriber, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubscriberMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubscriberCreateBulk) SaveX(ctx context.Context) []*Subscriber {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubscriberCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubscriberCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Subscriber.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SubscriberUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (scb *SubscriberCreateBulk) OnConflict(opts ...sql.ConflictOption) *SubscriberUpsertBulk {
	scb.conflict = opts
	return &SubscriberUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Subscriber.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (scb *SubscriberCreateBulk) OnConflictColumns(columns ...string) *SubscriberUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SubscriberUpsertBulk{
		create: scb,
	}
}

// SubscriberUpsertBulk is the builder for "upsert"-ing
// a bulk of Subscriber nodes.
type SubscriberUpsertBulk struct {
	create *SubscriberCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Subscriber.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(subscriber.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *SubscriberUpsertBulk) UpdateNewValues() *SubscriberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(subscriber.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Subscriber.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *SubscriberUpsertBulk) Ignore() *SubscriberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SubscriberUpsertBulk) DoNothing() *SubscriberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SubscriberCreateBulk.OnConflict
// documentation for more info.
func (u *SubscriberUpsertBulk) Update(set func(*SubscriberUpsert)) *SubscriberUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SubscriberUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *SubscriberUpsertBulk) SetCreatedAt(v uint32) *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *SubscriberUpsertBulk) AddCreatedAt(v uint32) *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *SubscriberUpsertBulk) UpdateCreatedAt() *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *SubscriberUpsertBulk) SetUpdatedAt(v uint32) *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *SubscriberUpsertBulk) AddUpdatedAt(v uint32) *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *SubscriberUpsertBulk) UpdateUpdatedAt() *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *SubscriberUpsertBulk) SetDeletedAt(v uint32) *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *SubscriberUpsertBulk) AddDeletedAt(v uint32) *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *SubscriberUpsertBulk) UpdateDeletedAt() *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *SubscriberUpsertBulk) SetAppID(v uuid.UUID) *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *SubscriberUpsertBulk) UpdateAppID() *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *SubscriberUpsertBulk) ClearAppID() *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.ClearAppID()
	})
}

// SetEmailAddress sets the "email_address" field.
func (u *SubscriberUpsertBulk) SetEmailAddress(v string) *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.SetEmailAddress(v)
	})
}

// UpdateEmailAddress sets the "email_address" field to the value that was provided on create.
func (u *SubscriberUpsertBulk) UpdateEmailAddress() *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.UpdateEmailAddress()
	})
}

// ClearEmailAddress clears the value of the "email_address" field.
func (u *SubscriberUpsertBulk) ClearEmailAddress() *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.ClearEmailAddress()
	})
}

// SetRegistered sets the "registered" field.
func (u *SubscriberUpsertBulk) SetRegistered(v bool) *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.SetRegistered(v)
	})
}

// UpdateRegistered sets the "registered" field to the value that was provided on create.
func (u *SubscriberUpsertBulk) UpdateRegistered() *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.UpdateRegistered()
	})
}

// ClearRegistered clears the value of the "registered" field.
func (u *SubscriberUpsertBulk) ClearRegistered() *SubscriberUpsertBulk {
	return u.Update(func(s *SubscriberUpsert) {
		s.ClearRegistered()
	})
}

// Exec executes the query.
func (u *SubscriberUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the SubscriberCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for SubscriberCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SubscriberUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}