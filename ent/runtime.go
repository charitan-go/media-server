// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/charitan-go/media-server/ent/media"
	"github.com/charitan-go/media-server/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	mediaFields := schema.Media{}.Fields()
	_ = mediaFields
	// mediaDescReadableID is the schema descriptor for readable_id field.
	mediaDescReadableID := mediaFields[1].Descriptor()
	// media.DefaultReadableID holds the default value on creation for the readable_id field.
	media.DefaultReadableID = mediaDescReadableID.Default.(func() uuid.UUID)
	// mediaDescID is the schema descriptor for id field.
	mediaDescID := mediaFields[0].Descriptor()
	// media.DefaultID holds the default value on creation for the id field.
	media.DefaultID = mediaDescID.Default.(func() uuid.UUID)
}
