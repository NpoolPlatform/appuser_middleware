// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/app"
	"github.com/google/uuid"
)

// AppCreate is the builder for creating a App entity.
type AppCreate struct {
	config
	mutation *AppMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ac *AppCreate) SetCreatedAt(u uint32) *AppCreate {
	ac.mutation.SetCreatedAt(u)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AppCreate) SetNillableCreatedAt(u *uint32) *AppCreate {
	if u != nil {
		ac.SetCreatedAt(*u)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AppCreate) SetUpdatedAt(u uint32) *AppCreate {
	ac.mutation.SetUpdatedAt(u)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AppCreate) SetNillableUpdatedAt(u *uint32) *AppCreate {
	if u != nil {
		ac.SetUpdatedAt(*u)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *AppCreate) SetDeletedAt(u uint32) *AppCreate {
	ac.mutation.SetDeletedAt(u)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *AppCreate) SetNillableDeletedAt(u *uint32) *AppCreate {
	if u != nil {
		ac.SetDeletedAt(*u)
	}
	return ac
}

// SetEntID sets the "ent_id" field.
func (ac *AppCreate) SetEntID(u uuid.UUID) *AppCreate {
	ac.mutation.SetEntID(u)
	return ac
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ac *AppCreate) SetNillableEntID(u *uuid.UUID) *AppCreate {
	if u != nil {
		ac.SetEntID(*u)
	}
	return ac
}

// SetCreatedBy sets the "created_by" field.
func (ac *AppCreate) SetCreatedBy(u uuid.UUID) *AppCreate {
	ac.mutation.SetCreatedBy(u)
	return ac
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ac *AppCreate) SetNillableCreatedBy(u *uuid.UUID) *AppCreate {
	if u != nil {
		ac.SetCreatedBy(*u)
	}
	return ac
}

// SetName sets the "name" field.
func (ac *AppCreate) SetName(s string) *AppCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ac *AppCreate) SetNillableName(s *string) *AppCreate {
	if s != nil {
		ac.SetName(*s)
	}
	return ac
}

// SetLogo sets the "logo" field.
func (ac *AppCreate) SetLogo(s string) *AppCreate {
	ac.mutation.SetLogo(s)
	return ac
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (ac *AppCreate) SetNillableLogo(s *string) *AppCreate {
	if s != nil {
		ac.SetLogo(*s)
	}
	return ac
}

// SetDescription sets the "description" field.
func (ac *AppCreate) SetDescription(s string) *AppCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ac *AppCreate) SetNillableDescription(s *string) *AppCreate {
	if s != nil {
		ac.SetDescription(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AppCreate) SetID(u uint32) *AppCreate {
	ac.mutation.SetID(u)
	return ac
}

// Mutation returns the AppMutation object of the builder.
func (ac *AppCreate) Mutation() *AppMutation {
	return ac.mutation
}

// Save creates the App in the database.
func (ac *AppCreate) Save(ctx context.Context) (*App, error) {
	var (
		err  error
		node *App
	)
	if err := ac.defaults(); err != nil {
		return nil, err
	}
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*App)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AppCreate) SaveX(ctx context.Context) *App {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AppCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AppCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AppCreate) defaults() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		if app.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized app.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := app.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		if app.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized app.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := app.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.DeletedAt(); !ok {
		if app.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized app.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := app.DefaultDeletedAt()
		ac.mutation.SetDeletedAt(v)
	}
	if _, ok := ac.mutation.EntID(); !ok {
		if app.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized app.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := app.DefaultEntID()
		ac.mutation.SetEntID(v)
	}
	if _, ok := ac.mutation.CreatedBy(); !ok {
		if app.DefaultCreatedBy == nil {
			return fmt.Errorf("ent: uninitialized app.DefaultCreatedBy (forgotten import ent/runtime?)")
		}
		v := app.DefaultCreatedBy()
		ac.mutation.SetCreatedBy(v)
	}
	if _, ok := ac.mutation.Name(); !ok {
		v := app.DefaultName
		ac.mutation.SetName(v)
	}
	if _, ok := ac.mutation.Logo(); !ok {
		v := app.DefaultLogo
		ac.mutation.SetLogo(v)
	}
	if _, ok := ac.mutation.Description(); !ok {
		v := app.DefaultDescription
		ac.mutation.SetDescription(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ac *AppCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "App.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "App.updated_at"`)}
	}
	if _, ok := ac.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "App.deleted_at"`)}
	}
	if _, ok := ac.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "App.ent_id"`)}
	}
	return nil
}

func (ac *AppCreate) sqlSave(ctx context.Context) (*App, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
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

func (ac *AppCreate) createSpec() (*App, *sqlgraph.CreateSpec) {
	var (
		_node = &App{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: app.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: app.FieldID,
			},
		}
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: app.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: app.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: app.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := ac.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: app.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := ac.mutation.CreatedBy(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: app.FieldCreatedBy,
		})
		_node.CreatedBy = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ac.mutation.Logo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldLogo,
		})
		_node.Logo = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: app.FieldDescription,
		})
		_node.Description = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.App.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ac *AppCreate) OnConflict(opts ...sql.ConflictOption) *AppUpsertOne {
	ac.conflict = opts
	return &AppUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ac *AppCreate) OnConflictColumns(columns ...string) *AppUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AppUpsertOne{
		create: ac,
	}
}

type (
	// AppUpsertOne is the builder for "upsert"-ing
	//  one App node.
	AppUpsertOne struct {
		create *AppCreate
	}

	// AppUpsert is the "OnConflict" setter.
	AppUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AppUpsert) SetCreatedAt(v uint32) *AppUpsert {
	u.Set(app.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppUpsert) UpdateCreatedAt() *AppUpsert {
	u.SetExcluded(app.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppUpsert) AddCreatedAt(v uint32) *AppUpsert {
	u.Add(app.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppUpsert) SetUpdatedAt(v uint32) *AppUpsert {
	u.Set(app.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppUpsert) UpdateUpdatedAt() *AppUpsert {
	u.SetExcluded(app.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppUpsert) AddUpdatedAt(v uint32) *AppUpsert {
	u.Add(app.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppUpsert) SetDeletedAt(v uint32) *AppUpsert {
	u.Set(app.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppUpsert) UpdateDeletedAt() *AppUpsert {
	u.SetExcluded(app.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppUpsert) AddDeletedAt(v uint32) *AppUpsert {
	u.Add(app.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *AppUpsert) SetEntID(v uuid.UUID) *AppUpsert {
	u.Set(app.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppUpsert) UpdateEntID() *AppUpsert {
	u.SetExcluded(app.FieldEntID)
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *AppUpsert) SetCreatedBy(v uuid.UUID) *AppUpsert {
	u.Set(app.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *AppUpsert) UpdateCreatedBy() *AppUpsert {
	u.SetExcluded(app.FieldCreatedBy)
	return u
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *AppUpsert) ClearCreatedBy() *AppUpsert {
	u.SetNull(app.FieldCreatedBy)
	return u
}

// SetName sets the "name" field.
func (u *AppUpsert) SetName(v string) *AppUpsert {
	u.Set(app.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppUpsert) UpdateName() *AppUpsert {
	u.SetExcluded(app.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *AppUpsert) ClearName() *AppUpsert {
	u.SetNull(app.FieldName)
	return u
}

// SetLogo sets the "logo" field.
func (u *AppUpsert) SetLogo(v string) *AppUpsert {
	u.Set(app.FieldLogo, v)
	return u
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *AppUpsert) UpdateLogo() *AppUpsert {
	u.SetExcluded(app.FieldLogo)
	return u
}

// ClearLogo clears the value of the "logo" field.
func (u *AppUpsert) ClearLogo() *AppUpsert {
	u.SetNull(app.FieldLogo)
	return u
}

// SetDescription sets the "description" field.
func (u *AppUpsert) SetDescription(v string) *AppUpsert {
	u.Set(app.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppUpsert) UpdateDescription() *AppUpsert {
	u.SetExcluded(app.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *AppUpsert) ClearDescription() *AppUpsert {
	u.SetNull(app.FieldDescription)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(app.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppUpsertOne) UpdateNewValues() *AppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(app.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.App.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppUpsertOne) Ignore() *AppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUpsertOne) DoNothing() *AppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppCreate.OnConflict
// documentation for more info.
func (u *AppUpsertOne) Update(set func(*AppUpsert)) *AppUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppUpsertOne) SetCreatedAt(v uint32) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppUpsertOne) AddCreatedAt(v uint32) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateCreatedAt() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppUpsertOne) SetUpdatedAt(v uint32) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppUpsertOne) AddUpdatedAt(v uint32) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateUpdatedAt() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppUpsertOne) SetDeletedAt(v uint32) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppUpsertOne) AddDeletedAt(v uint32) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateDeletedAt() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AppUpsertOne) SetEntID(v uuid.UUID) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateEntID() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateEntID()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *AppUpsertOne) SetCreatedBy(v uuid.UUID) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateCreatedBy() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *AppUpsertOne) ClearCreatedBy() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.ClearCreatedBy()
	})
}

// SetName sets the "name" field.
func (u *AppUpsertOne) SetName(v string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateName() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *AppUpsertOne) ClearName() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.ClearName()
	})
}

// SetLogo sets the "logo" field.
func (u *AppUpsertOne) SetLogo(v string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetLogo(v)
	})
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateLogo() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateLogo()
	})
}

// ClearLogo clears the value of the "logo" field.
func (u *AppUpsertOne) ClearLogo() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.ClearLogo()
	})
}

// SetDescription sets the "description" field.
func (u *AppUpsertOne) SetDescription(v string) *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppUpsertOne) UpdateDescription() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *AppUpsertOne) ClearDescription() *AppUpsertOne {
	return u.Update(func(s *AppUpsert) {
		s.ClearDescription()
	})
}

// Exec executes the query.
func (u *AppUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppCreateBulk is the builder for creating many App entities in bulk.
type AppCreateBulk struct {
	config
	builders []*AppCreate
	conflict []sql.ConflictOption
}

// Save creates the App entities in the database.
func (acb *AppCreateBulk) Save(ctx context.Context) ([]*App, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*App, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AppCreateBulk) SaveX(ctx context.Context) []*App {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AppCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AppCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.App.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (acb *AppCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppUpsertBulk {
	acb.conflict = opts
	return &AppUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (acb *AppCreateBulk) OnConflictColumns(columns ...string) *AppUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AppUpsertBulk{
		create: acb,
	}
}

// AppUpsertBulk is the builder for "upsert"-ing
// a bulk of App nodes.
type AppUpsertBulk struct {
	create *AppCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(app.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppUpsertBulk) UpdateNewValues() *AppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(app.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.App.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppUpsertBulk) Ignore() *AppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUpsertBulk) DoNothing() *AppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppCreateBulk.OnConflict
// documentation for more info.
func (u *AppUpsertBulk) Update(set func(*AppUpsert)) *AppUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppUpsertBulk) SetCreatedAt(v uint32) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppUpsertBulk) AddCreatedAt(v uint32) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateCreatedAt() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppUpsertBulk) SetUpdatedAt(v uint32) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppUpsertBulk) AddUpdatedAt(v uint32) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateUpdatedAt() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppUpsertBulk) SetDeletedAt(v uint32) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppUpsertBulk) AddDeletedAt(v uint32) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateDeletedAt() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AppUpsertBulk) SetEntID(v uuid.UUID) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateEntID() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateEntID()
	})
}

// SetCreatedBy sets the "created_by" field.
func (u *AppUpsertBulk) SetCreatedBy(v uuid.UUID) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateCreatedBy() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateCreatedBy()
	})
}

// ClearCreatedBy clears the value of the "created_by" field.
func (u *AppUpsertBulk) ClearCreatedBy() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.ClearCreatedBy()
	})
}

// SetName sets the "name" field.
func (u *AppUpsertBulk) SetName(v string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateName() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *AppUpsertBulk) ClearName() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.ClearName()
	})
}

// SetLogo sets the "logo" field.
func (u *AppUpsertBulk) SetLogo(v string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetLogo(v)
	})
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateLogo() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateLogo()
	})
}

// ClearLogo clears the value of the "logo" field.
func (u *AppUpsertBulk) ClearLogo() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.ClearLogo()
	})
}

// SetDescription sets the "description" field.
func (u *AppUpsertBulk) SetDescription(v string) *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AppUpsertBulk) UpdateDescription() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *AppUpsertBulk) ClearDescription() *AppUpsertBulk {
	return u.Update(func(s *AppUpsert) {
		s.ClearDescription()
	})
}

// Exec executes the query.
func (u *AppUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
