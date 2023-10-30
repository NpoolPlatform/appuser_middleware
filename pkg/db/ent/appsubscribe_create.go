// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appsubscribe"
	"github.com/google/uuid"
)

// AppSubscribeCreate is the builder for creating a AppSubscribe entity.
type AppSubscribeCreate struct {
	config
	mutation *AppSubscribeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (asc *AppSubscribeCreate) SetCreatedAt(u uint32) *AppSubscribeCreate {
	asc.mutation.SetCreatedAt(u)
	return asc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (asc *AppSubscribeCreate) SetNillableCreatedAt(u *uint32) *AppSubscribeCreate {
	if u != nil {
		asc.SetCreatedAt(*u)
	}
	return asc
}

// SetUpdatedAt sets the "updated_at" field.
func (asc *AppSubscribeCreate) SetUpdatedAt(u uint32) *AppSubscribeCreate {
	asc.mutation.SetUpdatedAt(u)
	return asc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (asc *AppSubscribeCreate) SetNillableUpdatedAt(u *uint32) *AppSubscribeCreate {
	if u != nil {
		asc.SetUpdatedAt(*u)
	}
	return asc
}

// SetDeletedAt sets the "deleted_at" field.
func (asc *AppSubscribeCreate) SetDeletedAt(u uint32) *AppSubscribeCreate {
	asc.mutation.SetDeletedAt(u)
	return asc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (asc *AppSubscribeCreate) SetNillableDeletedAt(u *uint32) *AppSubscribeCreate {
	if u != nil {
		asc.SetDeletedAt(*u)
	}
	return asc
}

// SetEntID sets the "ent_id" field.
func (asc *AppSubscribeCreate) SetEntID(u uuid.UUID) *AppSubscribeCreate {
	asc.mutation.SetEntID(u)
	return asc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (asc *AppSubscribeCreate) SetNillableEntID(u *uuid.UUID) *AppSubscribeCreate {
	if u != nil {
		asc.SetEntID(*u)
	}
	return asc
}

// SetAppID sets the "app_id" field.
func (asc *AppSubscribeCreate) SetAppID(u uuid.UUID) *AppSubscribeCreate {
	asc.mutation.SetAppID(u)
	return asc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (asc *AppSubscribeCreate) SetNillableAppID(u *uuid.UUID) *AppSubscribeCreate {
	if u != nil {
		asc.SetAppID(*u)
	}
	return asc
}

// SetSubscribeAppID sets the "subscribe_app_id" field.
func (asc *AppSubscribeCreate) SetSubscribeAppID(u uuid.UUID) *AppSubscribeCreate {
	asc.mutation.SetSubscribeAppID(u)
	return asc
}

// SetNillableSubscribeAppID sets the "subscribe_app_id" field if the given value is not nil.
func (asc *AppSubscribeCreate) SetNillableSubscribeAppID(u *uuid.UUID) *AppSubscribeCreate {
	if u != nil {
		asc.SetSubscribeAppID(*u)
	}
	return asc
}

// SetID sets the "id" field.
func (asc *AppSubscribeCreate) SetID(u uint32) *AppSubscribeCreate {
	asc.mutation.SetID(u)
	return asc
}

// Mutation returns the AppSubscribeMutation object of the builder.
func (asc *AppSubscribeCreate) Mutation() *AppSubscribeMutation {
	return asc.mutation
}

// Save creates the AppSubscribe in the database.
func (asc *AppSubscribeCreate) Save(ctx context.Context) (*AppSubscribe, error) {
	var (
		err  error
		node *AppSubscribe
	)
	if err := asc.defaults(); err != nil {
		return nil, err
	}
	if len(asc.hooks) == 0 {
		if err = asc.check(); err != nil {
			return nil, err
		}
		node, err = asc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppSubscribeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = asc.check(); err != nil {
				return nil, err
			}
			asc.mutation = mutation
			if node, err = asc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(asc.hooks) - 1; i >= 0; i-- {
			if asc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = asc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, asc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppSubscribe)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppSubscribeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (asc *AppSubscribeCreate) SaveX(ctx context.Context) *AppSubscribe {
	v, err := asc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (asc *AppSubscribeCreate) Exec(ctx context.Context) error {
	_, err := asc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asc *AppSubscribeCreate) ExecX(ctx context.Context) {
	if err := asc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asc *AppSubscribeCreate) defaults() error {
	if _, ok := asc.mutation.CreatedAt(); !ok {
		if appsubscribe.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized appsubscribe.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := appsubscribe.DefaultCreatedAt()
		asc.mutation.SetCreatedAt(v)
	}
	if _, ok := asc.mutation.UpdatedAt(); !ok {
		if appsubscribe.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appsubscribe.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appsubscribe.DefaultUpdatedAt()
		asc.mutation.SetUpdatedAt(v)
	}
	if _, ok := asc.mutation.DeletedAt(); !ok {
		if appsubscribe.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized appsubscribe.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := appsubscribe.DefaultDeletedAt()
		asc.mutation.SetDeletedAt(v)
	}
	if _, ok := asc.mutation.EntID(); !ok {
		if appsubscribe.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized appsubscribe.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := appsubscribe.DefaultEntID()
		asc.mutation.SetEntID(v)
	}
	if _, ok := asc.mutation.AppID(); !ok {
		if appsubscribe.DefaultAppID == nil {
			return fmt.Errorf("ent: uninitialized appsubscribe.DefaultAppID (forgotten import ent/runtime?)")
		}
		v := appsubscribe.DefaultAppID()
		asc.mutation.SetAppID(v)
	}
	if _, ok := asc.mutation.SubscribeAppID(); !ok {
		if appsubscribe.DefaultSubscribeAppID == nil {
			return fmt.Errorf("ent: uninitialized appsubscribe.DefaultSubscribeAppID (forgotten import ent/runtime?)")
		}
		v := appsubscribe.DefaultSubscribeAppID()
		asc.mutation.SetSubscribeAppID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (asc *AppSubscribeCreate) check() error {
	if _, ok := asc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppSubscribe.created_at"`)}
	}
	if _, ok := asc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppSubscribe.updated_at"`)}
	}
	if _, ok := asc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "AppSubscribe.deleted_at"`)}
	}
	if _, ok := asc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "AppSubscribe.ent_id"`)}
	}
	return nil
}

func (asc *AppSubscribeCreate) sqlSave(ctx context.Context) (*AppSubscribe, error) {
	_node, _spec := asc.createSpec()
	if err := sqlgraph.CreateNode(ctx, asc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (asc *AppSubscribeCreate) createSpec() (*AppSubscribe, *sqlgraph.CreateSpec) {
	var (
		_node = &AppSubscribe{config: asc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appsubscribe.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appsubscribe.FieldID,
			},
		}
	)
	_spec.OnConflict = asc.conflict
	if id, ok := asc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := asc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsubscribe.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := asc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsubscribe.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := asc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsubscribe.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := asc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsubscribe.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := asc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsubscribe.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := asc.mutation.SubscribeAppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsubscribe.FieldSubscribeAppID,
		})
		_node.SubscribeAppID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppSubscribe.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppSubscribeUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (asc *AppSubscribeCreate) OnConflict(opts ...sql.ConflictOption) *AppSubscribeUpsertOne {
	asc.conflict = opts
	return &AppSubscribeUpsertOne{
		create: asc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppSubscribe.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (asc *AppSubscribeCreate) OnConflictColumns(columns ...string) *AppSubscribeUpsertOne {
	asc.conflict = append(asc.conflict, sql.ConflictColumns(columns...))
	return &AppSubscribeUpsertOne{
		create: asc,
	}
}

type (
	// AppSubscribeUpsertOne is the builder for "upsert"-ing
	//  one AppSubscribe node.
	AppSubscribeUpsertOne struct {
		create *AppSubscribeCreate
	}

	// AppSubscribeUpsert is the "OnConflict" setter.
	AppSubscribeUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AppSubscribeUpsert) SetCreatedAt(v uint32) *AppSubscribeUpsert {
	u.Set(appsubscribe.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppSubscribeUpsert) UpdateCreatedAt() *AppSubscribeUpsert {
	u.SetExcluded(appsubscribe.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppSubscribeUpsert) AddCreatedAt(v uint32) *AppSubscribeUpsert {
	u.Add(appsubscribe.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppSubscribeUpsert) SetUpdatedAt(v uint32) *AppSubscribeUpsert {
	u.Set(appsubscribe.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppSubscribeUpsert) UpdateUpdatedAt() *AppSubscribeUpsert {
	u.SetExcluded(appsubscribe.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppSubscribeUpsert) AddUpdatedAt(v uint32) *AppSubscribeUpsert {
	u.Add(appsubscribe.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppSubscribeUpsert) SetDeletedAt(v uint32) *AppSubscribeUpsert {
	u.Set(appsubscribe.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppSubscribeUpsert) UpdateDeletedAt() *AppSubscribeUpsert {
	u.SetExcluded(appsubscribe.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppSubscribeUpsert) AddDeletedAt(v uint32) *AppSubscribeUpsert {
	u.Add(appsubscribe.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *AppSubscribeUpsert) SetEntID(v uuid.UUID) *AppSubscribeUpsert {
	u.Set(appsubscribe.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppSubscribeUpsert) UpdateEntID() *AppSubscribeUpsert {
	u.SetExcluded(appsubscribe.FieldEntID)
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppSubscribeUpsert) SetAppID(v uuid.UUID) *AppSubscribeUpsert {
	u.Set(appsubscribe.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppSubscribeUpsert) UpdateAppID() *AppSubscribeUpsert {
	u.SetExcluded(appsubscribe.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppSubscribeUpsert) ClearAppID() *AppSubscribeUpsert {
	u.SetNull(appsubscribe.FieldAppID)
	return u
}

// SetSubscribeAppID sets the "subscribe_app_id" field.
func (u *AppSubscribeUpsert) SetSubscribeAppID(v uuid.UUID) *AppSubscribeUpsert {
	u.Set(appsubscribe.FieldSubscribeAppID, v)
	return u
}

// UpdateSubscribeAppID sets the "subscribe_app_id" field to the value that was provided on create.
func (u *AppSubscribeUpsert) UpdateSubscribeAppID() *AppSubscribeUpsert {
	u.SetExcluded(appsubscribe.FieldSubscribeAppID)
	return u
}

// ClearSubscribeAppID clears the value of the "subscribe_app_id" field.
func (u *AppSubscribeUpsert) ClearSubscribeAppID() *AppSubscribeUpsert {
	u.SetNull(appsubscribe.FieldSubscribeAppID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppSubscribe.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appsubscribe.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppSubscribeUpsertOne) UpdateNewValues() *AppSubscribeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appsubscribe.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppSubscribe.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppSubscribeUpsertOne) Ignore() *AppSubscribeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppSubscribeUpsertOne) DoNothing() *AppSubscribeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppSubscribeCreate.OnConflict
// documentation for more info.
func (u *AppSubscribeUpsertOne) Update(set func(*AppSubscribeUpsert)) *AppSubscribeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppSubscribeUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppSubscribeUpsertOne) SetCreatedAt(v uint32) *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppSubscribeUpsertOne) AddCreatedAt(v uint32) *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppSubscribeUpsertOne) UpdateCreatedAt() *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppSubscribeUpsertOne) SetUpdatedAt(v uint32) *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppSubscribeUpsertOne) AddUpdatedAt(v uint32) *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppSubscribeUpsertOne) UpdateUpdatedAt() *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppSubscribeUpsertOne) SetDeletedAt(v uint32) *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppSubscribeUpsertOne) AddDeletedAt(v uint32) *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppSubscribeUpsertOne) UpdateDeletedAt() *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AppSubscribeUpsertOne) SetEntID(v uuid.UUID) *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppSubscribeUpsertOne) UpdateEntID() *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppSubscribeUpsertOne) SetAppID(v uuid.UUID) *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppSubscribeUpsertOne) UpdateAppID() *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppSubscribeUpsertOne) ClearAppID() *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.ClearAppID()
	})
}

// SetSubscribeAppID sets the "subscribe_app_id" field.
func (u *AppSubscribeUpsertOne) SetSubscribeAppID(v uuid.UUID) *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.SetSubscribeAppID(v)
	})
}

// UpdateSubscribeAppID sets the "subscribe_app_id" field to the value that was provided on create.
func (u *AppSubscribeUpsertOne) UpdateSubscribeAppID() *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.UpdateSubscribeAppID()
	})
}

// ClearSubscribeAppID clears the value of the "subscribe_app_id" field.
func (u *AppSubscribeUpsertOne) ClearSubscribeAppID() *AppSubscribeUpsertOne {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.ClearSubscribeAppID()
	})
}

// Exec executes the query.
func (u *AppSubscribeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppSubscribeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppSubscribeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppSubscribeUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppSubscribeUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppSubscribeCreateBulk is the builder for creating many AppSubscribe entities in bulk.
type AppSubscribeCreateBulk struct {
	config
	builders []*AppSubscribeCreate
	conflict []sql.ConflictOption
}

// Save creates the AppSubscribe entities in the database.
func (ascb *AppSubscribeCreateBulk) Save(ctx context.Context) ([]*AppSubscribe, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ascb.builders))
	nodes := make([]*AppSubscribe, len(ascb.builders))
	mutators := make([]Mutator, len(ascb.builders))
	for i := range ascb.builders {
		func(i int, root context.Context) {
			builder := ascb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppSubscribeMutation)
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
					_, err = mutators[i+1].Mutate(root, ascb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ascb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ascb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, ascb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ascb *AppSubscribeCreateBulk) SaveX(ctx context.Context) []*AppSubscribe {
	v, err := ascb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ascb *AppSubscribeCreateBulk) Exec(ctx context.Context) error {
	_, err := ascb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ascb *AppSubscribeCreateBulk) ExecX(ctx context.Context) {
	if err := ascb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppSubscribe.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppSubscribeUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ascb *AppSubscribeCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppSubscribeUpsertBulk {
	ascb.conflict = opts
	return &AppSubscribeUpsertBulk{
		create: ascb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppSubscribe.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ascb *AppSubscribeCreateBulk) OnConflictColumns(columns ...string) *AppSubscribeUpsertBulk {
	ascb.conflict = append(ascb.conflict, sql.ConflictColumns(columns...))
	return &AppSubscribeUpsertBulk{
		create: ascb,
	}
}

// AppSubscribeUpsertBulk is the builder for "upsert"-ing
// a bulk of AppSubscribe nodes.
type AppSubscribeUpsertBulk struct {
	create *AppSubscribeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppSubscribe.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appsubscribe.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppSubscribeUpsertBulk) UpdateNewValues() *AppSubscribeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appsubscribe.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppSubscribe.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppSubscribeUpsertBulk) Ignore() *AppSubscribeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppSubscribeUpsertBulk) DoNothing() *AppSubscribeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppSubscribeCreateBulk.OnConflict
// documentation for more info.
func (u *AppSubscribeUpsertBulk) Update(set func(*AppSubscribeUpsert)) *AppSubscribeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppSubscribeUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppSubscribeUpsertBulk) SetCreatedAt(v uint32) *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppSubscribeUpsertBulk) AddCreatedAt(v uint32) *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppSubscribeUpsertBulk) UpdateCreatedAt() *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppSubscribeUpsertBulk) SetUpdatedAt(v uint32) *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppSubscribeUpsertBulk) AddUpdatedAt(v uint32) *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppSubscribeUpsertBulk) UpdateUpdatedAt() *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppSubscribeUpsertBulk) SetDeletedAt(v uint32) *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppSubscribeUpsertBulk) AddDeletedAt(v uint32) *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppSubscribeUpsertBulk) UpdateDeletedAt() *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AppSubscribeUpsertBulk) SetEntID(v uuid.UUID) *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppSubscribeUpsertBulk) UpdateEntID() *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.UpdateEntID()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppSubscribeUpsertBulk) SetAppID(v uuid.UUID) *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppSubscribeUpsertBulk) UpdateAppID() *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppSubscribeUpsertBulk) ClearAppID() *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.ClearAppID()
	})
}

// SetSubscribeAppID sets the "subscribe_app_id" field.
func (u *AppSubscribeUpsertBulk) SetSubscribeAppID(v uuid.UUID) *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.SetSubscribeAppID(v)
	})
}

// UpdateSubscribeAppID sets the "subscribe_app_id" field to the value that was provided on create.
func (u *AppSubscribeUpsertBulk) UpdateSubscribeAppID() *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.UpdateSubscribeAppID()
	})
}

// ClearSubscribeAppID clears the value of the "subscribe_app_id" field.
func (u *AppSubscribeUpsertBulk) ClearSubscribeAppID() *AppSubscribeUpsertBulk {
	return u.Update(func(s *AppSubscribeUpsert) {
		s.ClearSubscribeAppID()
	})
}

// Exec executes the query.
func (u *AppSubscribeUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppSubscribeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppSubscribeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppSubscribeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
