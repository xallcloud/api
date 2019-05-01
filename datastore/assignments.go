package datastore

import "time"

const (
	// KindAssignments represents the "table" name (Kind) to store Assignments
	KindAssignments = "assignments"
)

// Assignment contains the structure to match the devices Datastore records
type Assignment struct {
	// Stored data
	AsID        string    `datastore:"asID"`
	Created     time.Time `datastore:"created"`
	Changed     time.Time `datastore:"changed"`
	Description string    `datastore:"description"`

	CpID string `datastore:"cpID"`
	DvID string `datastore:"dvID"`

	Level int32 `datastore:"level"`

	// no indexing
	Settings   string `datastore:"settings,noindex"`
	RawRequest string `datastore:"rawRequest,noindex"`
	// key
	ID int64 `datastore:"-"` // The integer ID used in the datastore.

	CallpointObj Callpoint `datastore:"-"`
	DeviceObj    Device    `datastore:"-"`
}
