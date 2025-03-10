// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/charitan-go/media-server/ent/media"
	"github.com/charitan-go/media-server/ent/predicate"
	"github.com/charitan-go/media-server/internal/media/dto"
	"github.com/google/uuid"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeMedia = "Media"
)

// MediaMutation represents an operation that mutates the Media nodes in the graph.
type MediaMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	readable_id   *uuid.UUID
	media_type    *dto.MediaTypeEnum
	project_id    *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Media, error)
	predicates    []predicate.Media
}

var _ ent.Mutation = (*MediaMutation)(nil)

// mediaOption allows management of the mutation configuration using functional options.
type mediaOption func(*MediaMutation)

// newMediaMutation creates new mutation for the Media entity.
func newMediaMutation(c config, op Op, opts ...mediaOption) *MediaMutation {
	m := &MediaMutation{
		config:        c,
		op:            op,
		typ:           TypeMedia,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMediaID sets the ID field of the mutation.
func withMediaID(id uuid.UUID) mediaOption {
	return func(m *MediaMutation) {
		var (
			err   error
			once  sync.Once
			value *Media
		)
		m.oldValue = func(ctx context.Context) (*Media, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Media.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMedia sets the old Media of the mutation.
func withMedia(node *Media) mediaOption {
	return func(m *MediaMutation) {
		m.oldValue = func(context.Context) (*Media, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MediaMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MediaMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Media entities.
func (m *MediaMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MediaMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *MediaMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Media.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetReadableID sets the "readable_id" field.
func (m *MediaMutation) SetReadableID(u uuid.UUID) {
	m.readable_id = &u
}

// ReadableID returns the value of the "readable_id" field in the mutation.
func (m *MediaMutation) ReadableID() (r uuid.UUID, exists bool) {
	v := m.readable_id
	if v == nil {
		return
	}
	return *v, true
}

// OldReadableID returns the old "readable_id" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldReadableID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldReadableID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldReadableID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldReadableID: %w", err)
	}
	return oldValue.ReadableID, nil
}

// ResetReadableID resets all changes to the "readable_id" field.
func (m *MediaMutation) ResetReadableID() {
	m.readable_id = nil
}

// SetMediaType sets the "media_type" field.
func (m *MediaMutation) SetMediaType(dte dto.MediaTypeEnum) {
	m.media_type = &dte
}

// MediaType returns the value of the "media_type" field in the mutation.
func (m *MediaMutation) MediaType() (r dto.MediaTypeEnum, exists bool) {
	v := m.media_type
	if v == nil {
		return
	}
	return *v, true
}

// OldMediaType returns the old "media_type" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldMediaType(ctx context.Context) (v dto.MediaTypeEnum, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldMediaType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldMediaType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMediaType: %w", err)
	}
	return oldValue.MediaType, nil
}

// ResetMediaType resets all changes to the "media_type" field.
func (m *MediaMutation) ResetMediaType() {
	m.media_type = nil
}

// SetProjectID sets the "project_id" field.
func (m *MediaMutation) SetProjectID(s string) {
	m.project_id = &s
}

// ProjectID returns the value of the "project_id" field in the mutation.
func (m *MediaMutation) ProjectID() (r string, exists bool) {
	v := m.project_id
	if v == nil {
		return
	}
	return *v, true
}

// OldProjectID returns the old "project_id" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldProjectID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldProjectID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldProjectID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldProjectID: %w", err)
	}
	return oldValue.ProjectID, nil
}

// ResetProjectID resets all changes to the "project_id" field.
func (m *MediaMutation) ResetProjectID() {
	m.project_id = nil
}

// Where appends a list predicates to the MediaMutation builder.
func (m *MediaMutation) Where(ps ...predicate.Media) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the MediaMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *MediaMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Media, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *MediaMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *MediaMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Media).
func (m *MediaMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MediaMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.readable_id != nil {
		fields = append(fields, media.FieldReadableID)
	}
	if m.media_type != nil {
		fields = append(fields, media.FieldMediaType)
	}
	if m.project_id != nil {
		fields = append(fields, media.FieldProjectID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MediaMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case media.FieldReadableID:
		return m.ReadableID()
	case media.FieldMediaType:
		return m.MediaType()
	case media.FieldProjectID:
		return m.ProjectID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MediaMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case media.FieldReadableID:
		return m.OldReadableID(ctx)
	case media.FieldMediaType:
		return m.OldMediaType(ctx)
	case media.FieldProjectID:
		return m.OldProjectID(ctx)
	}
	return nil, fmt.Errorf("unknown Media field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MediaMutation) SetField(name string, value ent.Value) error {
	switch name {
	case media.FieldReadableID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetReadableID(v)
		return nil
	case media.FieldMediaType:
		v, ok := value.(dto.MediaTypeEnum)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMediaType(v)
		return nil
	case media.FieldProjectID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetProjectID(v)
		return nil
	}
	return fmt.Errorf("unknown Media field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MediaMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MediaMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MediaMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Media numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MediaMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MediaMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MediaMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Media nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MediaMutation) ResetField(name string) error {
	switch name {
	case media.FieldReadableID:
		m.ResetReadableID()
		return nil
	case media.FieldMediaType:
		m.ResetMediaType()
		return nil
	case media.FieldProjectID:
		m.ResetProjectID()
		return nil
	}
	return fmt.Errorf("unknown Media field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MediaMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MediaMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MediaMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MediaMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MediaMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MediaMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MediaMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Media unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MediaMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Media edge %s", name)
}
