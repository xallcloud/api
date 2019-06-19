package datastore

import "time"

const (
	// KindCallpoints represents the "table" name (Kind) to store Callpoints
	KindCallpoints = "callpoints"
)

// Callpoint contains the structure to match the callpoints Datastore records
type Callpoint struct {
	// key
	ID int64 `datastore:"-"` // The integer ID used in the datastore.
	// Stored data
	CpID        string    `datastore:"cpID"`
	Label       string    `datastore:"label"`
	Type        int32     `datastore:"type"`
	Icon        string    `datastore:"icon,noindex"`
	Description string    `datastore:"description"`
	Priority    int32     `datastore:"priority"`
	AbsAddress  string    `datastore:"absAddress"`
	Created     time.Time `datastore:"created"`
	Changed     time.Time `datastore:"changed"`
	RawRequest  string    `datastore:"rawRequest,noindex"`
}
