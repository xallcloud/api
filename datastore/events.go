package datastore

import "time"

const (
	// KindEvents represents the "table" name (Kind) to store Events
	KindEvents = "events"
)

// Event contains the structure to match the events Datastore records
type Event struct {
	// key
	ID int64 `datastore:"-"` // The integer ID used in the datastore.
	// Stored data
	EvID          string    `datastore:"evID"`
	NtID          string    `datastore:"ntID"`
	CpID          string    `datastore:"cpID"`
	DvID          string    `datastore:"dvID"`
	Visibility    string    `datastore:"visibility"`
	EvType        string    `datastore:"evType"`
	EvSubType     string    `datastore:"evSubType"`
	EvDescription string    `datastore:"evDescription"`
	Created       time.Time `datastore:"created"`
}
