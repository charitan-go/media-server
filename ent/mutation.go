// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/charitan-go/media-server/ent/predicate"
	"github.com/charitan-go/media-server/ent/media"
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
	op                        Op
	typ                       string
	id                        *uuid.UUID
	readable_id               *uuid.UUID
	name                      *string
	description               *string
	goal                      *float64
	addgoal                   *float64
	start_date                *time.Time
	end_date                  *time.Time
	category                  *dto.CategoryEnum
	countryCode               *string
	status                    *dto.StatusEnum
	owner_charity_readable_id *string
	clearedFields             map[string]struct{}
	done                      bool
	oldValue                  func(context.Context) (*Media, error)
	predicates                []predicate.Media
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

// SetName sets the "name" field.
func (m *MediaMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *MediaMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *MediaMutation) ResetName() {
	m.name = nil
}

// SetDescription sets the "description" field.
func (m *MediaMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *MediaMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *MediaMutation) ResetDescription() {
	m.description = nil
}

// SetGoal sets the "goal" field.
func (m *MediaMutation) SetGoal(f float64) {
	m.goal = &f
	m.addgoal = nil
}

// Goal returns the value of the "goal" field in the mutation.
func (m *MediaMutation) Goal() (r float64, exists bool) {
	v := m.goal
	if v == nil {
		return
	}
	return *v, true
}

// OldGoal returns the old "goal" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldGoal(ctx context.Context) (v float64, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGoal is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGoal requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGoal: %w", err)
	}
	return oldValue.Goal, nil
}

// AddGoal adds f to the "goal" field.
func (m *MediaMutation) AddGoal(f float64) {
	if m.addgoal != nil {
		*m.addgoal += f
	} else {
		m.addgoal = &f
	}
}

// AddedGoal returns the value that was added to the "goal" field in this mutation.
func (m *MediaMutation) AddedGoal() (r float64, exists bool) {
	v := m.addgoal
	if v == nil {
		return
	}
	return *v, true
}

// ResetGoal resets all changes to the "goal" field.
func (m *MediaMutation) ResetGoal() {
	m.goal = nil
	m.addgoal = nil
}

// SetStartDate sets the "start_date" field.
func (m *MediaMutation) SetStartDate(t time.Time) {
	m.start_date = &t
}

// StartDate returns the value of the "start_date" field in the mutation.
func (m *MediaMutation) StartDate() (r time.Time, exists bool) {
	v := m.start_date
	if v == nil {
		return
	}
	return *v, true
}

// OldStartDate returns the old "start_date" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldStartDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStartDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStartDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStartDate: %w", err)
	}
	return oldValue.StartDate, nil
}

// ResetStartDate resets all changes to the "start_date" field.
func (m *MediaMutation) ResetStartDate() {
	m.start_date = nil
}

// SetEndDate sets the "end_date" field.
func (m *MediaMutation) SetEndDate(t time.Time) {
	m.end_date = &t
}

// EndDate returns the value of the "end_date" field in the mutation.
func (m *MediaMutation) EndDate() (r time.Time, exists bool) {
	v := m.end_date
	if v == nil {
		return
	}
	return *v, true
}

// OldEndDate returns the old "end_date" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldEndDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEndDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEndDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEndDate: %w", err)
	}
	return oldValue.EndDate, nil
}

// ResetEndDate resets all changes to the "end_date" field.
func (m *MediaMutation) ResetEndDate() {
	m.end_date = nil
}

// SetCategory sets the "category" field.
func (m *MediaMutation) SetCategory(de dto.CategoryEnum) {
	m.category = &de
}

// Category returns the value of the "category" field in the mutation.
func (m *MediaMutation) Category() (r dto.CategoryEnum, exists bool) {
	v := m.category
	if v == nil {
		return
	}
	return *v, true
}

// OldCategory returns the old "category" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldCategory(ctx context.Context) (v dto.CategoryEnum, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCategory is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCategory requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCategory: %w", err)
	}
	return oldValue.Category, nil
}

// ResetCategory resets all changes to the "category" field.
func (m *MediaMutation) ResetCategory() {
	m.category = nil
}

// SetCountryCode sets the "countryCode" field.
func (m *MediaMutation) SetCountryCode(s string) {
	m.countryCode = &s
}

// CountryCode returns the value of the "countryCode" field in the mutation.
func (m *MediaMutation) CountryCode() (r string, exists bool) {
	v := m.countryCode
	if v == nil {
		return
	}
	return *v, true
}

// OldCountryCode returns the old "countryCode" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldCountryCode(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCountryCode is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCountryCode requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCountryCode: %w", err)
	}
	return oldValue.CountryCode, nil
}

// ResetCountryCode resets all changes to the "countryCode" field.
func (m *MediaMutation) ResetCountryCode() {
	m.countryCode = nil
}

// SetStatus sets the "status" field.
func (m *MediaMutation) SetStatus(de dto.StatusEnum) {
	m.status = &de
}

// Status returns the value of the "status" field in the mutation.
func (m *MediaMutation) Status() (r dto.StatusEnum, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldStatus(ctx context.Context) (v dto.StatusEnum, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *MediaMutation) ResetStatus() {
	m.status = nil
}

// SetOwnerCharityReadableID sets the "owner_charity_readable_id" field.
func (m *MediaMutation) SetOwnerCharityReadableID(s string) {
	m.owner_charity_readable_id = &s
}

// OwnerCharityReadableID returns the value of the "owner_charity_readable_id" field in the mutation.
func (m *MediaMutation) OwnerCharityReadableID() (r string, exists bool) {
	v := m.owner_charity_readable_id
	if v == nil {
		return
	}
	return *v, true
}

// OldOwnerCharityReadableID returns the old "owner_charity_readable_id" field's value of the Media entity.
// If the Media object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MediaMutation) OldOwnerCharityReadableID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldOwnerCharityReadableID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldOwnerCharityReadableID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldOwnerCharityReadableID: %w", err)
	}
	return oldValue.OwnerCharityReadableID, nil
}

// ResetOwnerCharityReadableID resets all changes to the "owner_charity_readable_id" field.
func (m *MediaMutation) ResetOwnerCharityReadableID() {
	m.owner_charity_readable_id = nil
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
	fields := make([]string, 0, 10)
	if m.readable_id != nil {
		fields = append(fields, media.FieldReadableID)
	}
	if m.name != nil {
		fields = append(fields, media.FieldName)
	}
	if m.description != nil {
		fields = append(fields, media.FieldDescription)
	}
	if m.goal != nil {
		fields = append(fields, media.FieldGoal)
	}
	if m.start_date != nil {
		fields = append(fields, media.FieldStartDate)
	}
	if m.end_date != nil {
		fields = append(fields, media.FieldEndDate)
	}
	if m.category != nil {
		fields = append(fields, media.FieldCategory)
	}
	if m.countryCode != nil {
		fields = append(fields, media.FieldCountryCode)
	}
	if m.status != nil {
		fields = append(fields, media.FieldStatus)
	}
	if m.owner_charity_readable_id != nil {
		fields = append(fields, media.FieldOwnerCharityReadableID)
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
	case media.FieldName:
		return m.Name()
	case media.FieldDescription:
		return m.Description()
	case media.FieldGoal:
		return m.Goal()
	case media.FieldStartDate:
		return m.StartDate()
	case media.FieldEndDate:
		return m.EndDate()
	case media.FieldCategory:
		return m.Category()
	case media.FieldCountryCode:
		return m.CountryCode()
	case media.FieldStatus:
		return m.Status()
	case media.FieldOwnerCharityReadableID:
		return m.OwnerCharityReadableID()
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
	case media.FieldName:
		return m.OldName(ctx)
	case media.FieldDescription:
		return m.OldDescription(ctx)
	case media.FieldGoal:
		return m.OldGoal(ctx)
	case media.FieldStartDate:
		return m.OldStartDate(ctx)
	case media.FieldEndDate:
		return m.OldEndDate(ctx)
	case media.FieldCategory:
		return m.OldCategory(ctx)
	case media.FieldCountryCode:
		return m.OldCountryCode(ctx)
	case media.FieldStatus:
		return m.OldStatus(ctx)
	case media.FieldOwnerCharityReadableID:
		return m.OldOwnerCharityReadableID(ctx)
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
	case media.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case media.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case media.FieldGoal:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGoal(v)
		return nil
	case media.FieldStartDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStartDate(v)
		return nil
	case media.FieldEndDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEndDate(v)
		return nil
	case media.FieldCategory:
		v, ok := value.(dto.CategoryEnum)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCategory(v)
		return nil
	case media.FieldCountryCode:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCountryCode(v)
		return nil
	case media.FieldStatus:
		v, ok := value.(dto.StatusEnum)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case media.FieldOwnerCharityReadableID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetOwnerCharityReadableID(v)
		return nil
	}
	return fmt.Errorf("unknown Media field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MediaMutation) AddedFields() []string {
	var fields []string
	if m.addgoal != nil {
		fields = append(fields, media.FieldGoal)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MediaMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case media.FieldGoal:
		return m.AddedGoal()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MediaMutation) AddField(name string, value ent.Value) error {
	switch name {
	case media.FieldGoal:
		v, ok := value.(float64)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddGoal(v)
		return nil
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
	case media.FieldName:
		m.ResetName()
		return nil
	case media.FieldDescription:
		m.ResetDescription()
		return nil
	case media.FieldGoal:
		m.ResetGoal()
		return nil
	case media.FieldStartDate:
		m.ResetStartDate()
		return nil
	case media.FieldEndDate:
		m.ResetEndDate()
		return nil
	case media.FieldCategory:
		m.ResetCategory()
		return nil
	case media.FieldCountryCode:
		m.ResetCountryCode()
		return nil
	case media.FieldStatus:
		m.ResetStatus()
		return nil
	case media.FieldOwnerCharityReadableID:
		m.ResetOwnerCharityReadableID()
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
