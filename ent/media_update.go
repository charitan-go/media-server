// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/charitan-go/media-server/ent/media"
	"github.com/charitan-go/media-server/ent/predicate"
	"github.com/charitan-go/media-server/internal/media/dto"
	"github.com/google/uuid"
)

// MediaUpdate is the builder for updating Media entities.
type MediaUpdate struct {
	config
	hooks    []Hook
	mutation *MediaMutation
}

// Where appends a list predicates to the MediaUpdate builder.
func (mu *MediaUpdate) Where(ps ...predicate.Media) *MediaUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetReadableID sets the "readable_id" field.
func (mu *MediaUpdate) SetReadableID(u uuid.UUID) *MediaUpdate {
	mu.mutation.SetReadableID(u)
	return mu
}

// SetNillableReadableID sets the "readable_id" field if the given value is not nil.
func (mu *MediaUpdate) SetNillableReadableID(u *uuid.UUID) *MediaUpdate {
	if u != nil {
		mu.SetReadableID(*u)
	}
	return mu
}

// SetMediaType sets the "media_type" field.
func (mu *MediaUpdate) SetMediaType(dte dto.MediaTypeEnum) *MediaUpdate {
	mu.mutation.SetMediaType(dte)
	return mu
}

// SetNillableMediaType sets the "media_type" field if the given value is not nil.
func (mu *MediaUpdate) SetNillableMediaType(dte *dto.MediaTypeEnum) *MediaUpdate {
	if dte != nil {
		mu.SetMediaType(*dte)
	}
	return mu
}

// SetProjectID sets the "project_id" field.
func (mu *MediaUpdate) SetProjectID(s string) *MediaUpdate {
	mu.mutation.SetProjectID(s)
	return mu
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (mu *MediaUpdate) SetNillableProjectID(s *string) *MediaUpdate {
	if s != nil {
		mu.SetProjectID(*s)
	}
	return mu
}

// Mutation returns the MediaMutation object of the builder.
func (mu *MediaUpdate) Mutation() *MediaMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MediaUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MediaUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MediaUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MediaUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MediaUpdate) check() error {
	if v, ok := mu.mutation.MediaType(); ok {
		if err := media.MediaTypeValidator(v); err != nil {
			return &ValidationError{Name: "media_type", err: fmt.Errorf(`ent: validator failed for field "Media.media_type": %w`, err)}
		}
	}
	return nil
}

func (mu *MediaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(media.Table, media.Columns, sqlgraph.NewFieldSpec(media.FieldID, field.TypeUUID))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.ReadableID(); ok {
		_spec.SetField(media.FieldReadableID, field.TypeUUID, value)
	}
	if value, ok := mu.mutation.MediaType(); ok {
		_spec.SetField(media.FieldMediaType, field.TypeEnum, value)
	}
	if value, ok := mu.mutation.ProjectID(); ok {
		_spec.SetField(media.FieldProjectID, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{media.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MediaUpdateOne is the builder for updating a single Media entity.
type MediaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MediaMutation
}

// SetReadableID sets the "readable_id" field.
func (muo *MediaUpdateOne) SetReadableID(u uuid.UUID) *MediaUpdateOne {
	muo.mutation.SetReadableID(u)
	return muo
}

// SetNillableReadableID sets the "readable_id" field if the given value is not nil.
func (muo *MediaUpdateOne) SetNillableReadableID(u *uuid.UUID) *MediaUpdateOne {
	if u != nil {
		muo.SetReadableID(*u)
	}
	return muo
}

// SetMediaType sets the "media_type" field.
func (muo *MediaUpdateOne) SetMediaType(dte dto.MediaTypeEnum) *MediaUpdateOne {
	muo.mutation.SetMediaType(dte)
	return muo
}

// SetNillableMediaType sets the "media_type" field if the given value is not nil.
func (muo *MediaUpdateOne) SetNillableMediaType(dte *dto.MediaTypeEnum) *MediaUpdateOne {
	if dte != nil {
		muo.SetMediaType(*dte)
	}
	return muo
}

// SetProjectID sets the "project_id" field.
func (muo *MediaUpdateOne) SetProjectID(s string) *MediaUpdateOne {
	muo.mutation.SetProjectID(s)
	return muo
}

// SetNillableProjectID sets the "project_id" field if the given value is not nil.
func (muo *MediaUpdateOne) SetNillableProjectID(s *string) *MediaUpdateOne {
	if s != nil {
		muo.SetProjectID(*s)
	}
	return muo
}

// Mutation returns the MediaMutation object of the builder.
func (muo *MediaUpdateOne) Mutation() *MediaMutation {
	return muo.mutation
}

// Where appends a list predicates to the MediaUpdate builder.
func (muo *MediaUpdateOne) Where(ps ...predicate.Media) *MediaUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MediaUpdateOne) Select(field string, fields ...string) *MediaUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Media entity.
func (muo *MediaUpdateOne) Save(ctx context.Context) (*Media, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MediaUpdateOne) SaveX(ctx context.Context) *Media {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MediaUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MediaUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MediaUpdateOne) check() error {
	if v, ok := muo.mutation.MediaType(); ok {
		if err := media.MediaTypeValidator(v); err != nil {
			return &ValidationError{Name: "media_type", err: fmt.Errorf(`ent: validator failed for field "Media.media_type": %w`, err)}
		}
	}
	return nil
}

func (muo *MediaUpdateOne) sqlSave(ctx context.Context) (_node *Media, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(media.Table, media.Columns, sqlgraph.NewFieldSpec(media.FieldID, field.TypeUUID))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Media.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, media.FieldID)
		for _, f := range fields {
			if !media.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != media.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.ReadableID(); ok {
		_spec.SetField(media.FieldReadableID, field.TypeUUID, value)
	}
	if value, ok := muo.mutation.MediaType(); ok {
		_spec.SetField(media.FieldMediaType, field.TypeEnum, value)
	}
	if value, ok := muo.mutation.ProjectID(); ok {
		_spec.SetField(media.FieldProjectID, field.TypeString, value)
	}
	_node = &Media{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{media.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
