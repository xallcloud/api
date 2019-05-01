package datastore

import "time"

const (
	// KindCallpoints represents the "table" name (Kind) to store Callpoints
	KindCallpoints = "callpoints"
)

// Callpoint contains the structure to match the callpoints Datastore records
type Callpoint struct {
	// key
	ID int64 // The integer ID used in the datastore.
	// Stored data
	CpID        string    `datastore:"cpID"`
	Created     time.Time `datastore:"created"`
	AbsAddress  string    `datastore:"absAddress"`
	Label       string    `datastore:"label"`
	Description string    `datastore:"description"`
	Type        int32     `datastore:"type"`
	Priority    int32     `datastore:"priority"`
	// no indexing
	Icon       string `datastore:"icon,noindex"`
	RawRequest string `datastore:"rawRequest,noindex"`
}
