// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/appuser"
	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppUserQuery is the builder for querying AppUser entities.
type AppUserQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppUser
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppUserQuery builder.
func (auq *AppUserQuery) Where(ps ...predicate.AppUser) *AppUserQuery {
	auq.predicates = append(auq.predicates, ps...)
	return auq
}

// Limit adds a limit step to the query.
func (auq *AppUserQuery) Limit(limit int) *AppUserQuery {
	auq.limit = &limit
	return auq
}

// Offset adds an offset step to the query.
func (auq *AppUserQuery) Offset(offset int) *AppUserQuery {
	auq.offset = &offset
	return auq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (auq *AppUserQuery) Unique(unique bool) *AppUserQuery {
	auq.unique = &unique
	return auq
}

// Order adds an order step to the query.
func (auq *AppUserQuery) Order(o ...OrderFunc) *AppUserQuery {
	auq.order = append(auq.order, o...)
	return auq
}

// First returns the first AppUser entity from the query.
// Returns a *NotFoundError when no AppUser was found.
func (auq *AppUserQuery) First(ctx context.Context) (*AppUser, error) {
	nodes, err := auq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (auq *AppUserQuery) FirstX(ctx context.Context) *AppUser {
	node, err := auq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppUser ID from the query.
// Returns a *NotFoundError when no AppUser ID was found.
func (auq *AppUserQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = auq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (auq *AppUserQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := auq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppUser entity is found.
// Returns a *NotFoundError when no AppUser entities are found.
func (auq *AppUserQuery) Only(ctx context.Context) (*AppUser, error) {
	nodes, err := auq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appuser.Label}
	default:
		return nil, &NotSingularError{appuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (auq *AppUserQuery) OnlyX(ctx context.Context) *AppUser {
	node, err := auq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppUser ID in the query.
// Returns a *NotSingularError when more than one AppUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (auq *AppUserQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = auq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appuser.Label}
	default:
		err = &NotSingularError{appuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (auq *AppUserQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := auq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppUsers.
func (auq *AppUserQuery) All(ctx context.Context) ([]*AppUser, error) {
	if err := auq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return auq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (auq *AppUserQuery) AllX(ctx context.Context) []*AppUser {
	nodes, err := auq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppUser IDs.
func (auq *AppUserQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := auq.Select(appuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (auq *AppUserQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := auq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (auq *AppUserQuery) Count(ctx context.Context) (int, error) {
	if err := auq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return auq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (auq *AppUserQuery) CountX(ctx context.Context) int {
	count, err := auq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (auq *AppUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := auq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return auq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (auq *AppUserQuery) ExistX(ctx context.Context) bool {
	exist, err := auq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (auq *AppUserQuery) Clone() *AppUserQuery {
	if auq == nil {
		return nil
	}
	return &AppUserQuery{
		config:     auq.config,
		limit:      auq.limit,
		offset:     auq.offset,
		order:      append([]OrderFunc{}, auq.order...),
		predicates: append([]predicate.AppUser{}, auq.predicates...),
		// clone intermediate query.
		sql:    auq.sql.Clone(),
		path:   auq.path,
		unique: auq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AppUser.Query().
//		GroupBy(appuser.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (auq *AppUserQuery) GroupBy(field string, fields ...string) *AppUserGroupBy {
	grbuild := &AppUserGroupBy{config: auq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := auq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return auq.sqlQuery(ctx), nil
	}
	grbuild.label = appuser.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.AppUser.Query().
//		Select(appuser.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (auq *AppUserQuery) Select(fields ...string) *AppUserSelect {
	auq.fields = append(auq.fields, fields...)
	selbuild := &AppUserSelect{AppUserQuery: auq}
	selbuild.label = appuser.Label
	selbuild.flds, selbuild.scan = &auq.fields, selbuild.Scan
	return selbuild
}

func (auq *AppUserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range auq.fields {
		if !appuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if auq.path != nil {
		prev, err := auq.path(ctx)
		if err != nil {
			return err
		}
		auq.sql = prev
	}
	if appuser.Policy == nil {
		return errors.New("ent: uninitialized appuser.Policy (forgotten import ent/runtime?)")
	}
	if err := appuser.Policy.EvalQuery(ctx, auq); err != nil {
		return err
	}
	return nil
}

func (auq *AppUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppUser, error) {
	var (
		nodes = []*AppUser{}
		_spec = auq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AppUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AppUser{config: auq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(auq.modifiers) > 0 {
		_spec.Modifiers = auq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, auq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (auq *AppUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := auq.querySpec()
	if len(auq.modifiers) > 0 {
		_spec.Modifiers = auq.modifiers
	}
	_spec.Node.Columns = auq.fields
	if len(auq.fields) > 0 {
		_spec.Unique = auq.unique != nil && *auq.unique
	}
	return sqlgraph.CountNodes(ctx, auq.driver, _spec)
}

func (auq *AppUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := auq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (auq *AppUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appuser.Table,
			Columns: appuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuser.FieldID,
			},
		},
		From:   auq.sql,
		Unique: true,
	}
	if unique := auq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := auq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appuser.FieldID)
		for i := range fields {
			if fields[i] != appuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := auq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := auq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := auq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := auq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (auq *AppUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(auq.driver.Dialect())
	t1 := builder.Table(appuser.Table)
	columns := auq.fields
	if len(columns) == 0 {
		columns = appuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if auq.sql != nil {
		selector = auq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if auq.unique != nil && *auq.unique {
		selector.Distinct()
	}
	for _, m := range auq.modifiers {
		m(selector)
	}
	for _, p := range auq.predicates {
		p(selector)
	}
	for _, p := range auq.order {
		p(selector)
	}
	if offset := auq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := auq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (auq *AppUserQuery) ForUpdate(opts ...sql.LockOption) *AppUserQuery {
	if auq.driver.Dialect() == dialect.Postgres {
		auq.Unique(false)
	}
	auq.modifiers = append(auq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return auq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (auq *AppUserQuery) ForShare(opts ...sql.LockOption) *AppUserQuery {
	if auq.driver.Dialect() == dialect.Postgres {
		auq.Unique(false)
	}
	auq.modifiers = append(auq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return auq
}

// AppUserGroupBy is the group-by builder for AppUser entities.
type AppUserGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (augb *AppUserGroupBy) Aggregate(fns ...AggregateFunc) *AppUserGroupBy {
	augb.fns = append(augb.fns, fns...)
	return augb
}

// Scan applies the group-by query and scans the result into the given value.
func (augb *AppUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := augb.path(ctx)
	if err != nil {
		return err
	}
	augb.sql = query
	return augb.sqlScan(ctx, v)
}

func (augb *AppUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range augb.fields {
		if !appuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := augb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := augb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (augb *AppUserGroupBy) sqlQuery() *sql.Selector {
	selector := augb.sql.Select()
	aggregation := make([]string, 0, len(augb.fns))
	for _, fn := range augb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(augb.fields)+len(augb.fns))
		for _, f := range augb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(augb.fields...)...)
}

// AppUserSelect is the builder for selecting fields of AppUser entities.
type AppUserSelect struct {
	*AppUserQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (aus *AppUserSelect) Scan(ctx context.Context, v interface{}) error {
	if err := aus.prepareQuery(ctx); err != nil {
		return err
	}
	aus.sql = aus.AppUserQuery.sqlQuery(ctx)
	return aus.sqlScan(ctx, v)
}

func (aus *AppUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := aus.sql.Query()
	if err := aus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
