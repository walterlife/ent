// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package blob

import (
	"github.com/facebookincubator/ent/entc/integration/customid/ent/schema"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the blob type in the database.
	Label = "blob"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid vertex property in the database.
	FieldUUID = "uuid"

	// Table holds the table name of the blob in the database.
	Table = "blobs"
)

// Columns holds all SQL columns are blob fields.
var Columns = []string{
	FieldID,
	FieldUUID,
}

var (
	fields = schema.Blob{}.Fields()

	// descUUID is the schema descriptor for uuid field.
	descUUID = fields[1].Descriptor()
	// DefaultUUID holds the default value on creation for the uuid field.
	DefaultUUID = descUUID.Default.(func() uuid.UUID)
)
