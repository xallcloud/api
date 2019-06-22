package datastore

import "time"

const (
	// KindActions represents the "table" name (Kind) to store Actions
	KindActions = "actions"
)

// Action contains the structure to match the devices Datastore records
type Action struct {
	// key
	ID int64 `datastore:"-"` // The integer ID used in the datastore.
	// Stored data
	AcID        string    `datastore:"acID"`
	CpID        string    `datastore:"cpID"`
	Action      string    `datastore:"action"`
	Description string    `datastore:"description"`
	Created     time.Time `datastore:"created"`
	RawRequest  string    `datastore:"rawRequest,noindex"`
}
