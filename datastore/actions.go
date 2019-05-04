package datastore

import "time"

const (
	// KindDevices represents the "table" name (Kind) to store Devices
	KindDevices = "actions"
)

// Device contains the structure to match the devices Datastore records
type Device struct {
	// Stored data
	AcID        string    `datastore:"acID"`
	Created     time.Time `datastore:"created"`
	Description string    `datastore:"description"`

	CpID   string `datastore:"category"`
	Action string `datastore:"action"`

	// no indexing
	RawRequest string `datastore:"rawRequest,noindex"`
	// key
	ID int64 `datastore:"-"` // The integer ID used in the datastore.
}
