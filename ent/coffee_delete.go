// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"main/ent/coffee"
	"main/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CoffeeDelete is the builder for deleting a Coffee entity.
type CoffeeDelete struct {
	config
	hooks    []Hook
	mutation *CoffeeMutation
}

// Where appends a list predicates to the CoffeeDelete builder.
func (cd *CoffeeDelete) Where(ps ...predicate.Coffee) *CoffeeDelete {
	cd.mutation.Where(ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *CoffeeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cd.sqlExec, cd.mutation, cd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CoffeeDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *CoffeeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(coffee.Table, sqlgraph.NewFieldSpec(coffee.FieldID, field.TypeInt))
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cd.mutation.done = true
	return affected, err
}

// CoffeeDeleteOne is the builder for deleting a single Coffee entity.
type CoffeeDeleteOne struct {
	cd *CoffeeDelete
}

// Where appends a list predicates to the CoffeeDelete builder.
func (cdo *CoffeeDeleteOne) Where(ps ...predicate.Coffee) *CoffeeDeleteOne {
	cdo.cd.mutation.Where(ps...)
	return cdo
}

// Exec executes the deletion query.
func (cdo *CoffeeDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{coffee.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CoffeeDeleteOne) ExecX(ctx context.Context) {
	if err := cdo.Exec(ctx); err != nil {
		panic(err)
	}
}