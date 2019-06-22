package datastore

import "time"

const (
	// KindAssignments represents the "table" name (Kind) to store Assignments
	KindAssignments = "assignments"
)

// Assignment contains the structure to match the assignments Datastore records
type Assignment struct {
	// key
	ID int64 `datastore:"-"` // The integer ID used in the datastore.
	// Stored data
	AsID        string    `datastore:"asID"`
	DvID        string    `datastore:"dvID"`
	CpID        string    `datastore:"cpID"`
	Description string    `datastore:"description"`
	Level       int32     `datastore:"level"`
	Created     time.Time `datastore:"created"`
	Changed     time.Time `datastore:"changed"`
	Settings    string    `datastore:"settings,noindex"`
	RawRequest  string    `datastore:"rawRequest,noindex"`
	// Query information
	CallpointObj Callpoint `datastore:"-"`
	DeviceObj    Device    `datastore:"-"`
}
