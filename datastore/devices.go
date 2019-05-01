package datastore

import "time"

const (
	// KindDevices represents the "table" name (Kind) to store Devices
	KindDevices = "devices"
)

// Device contains the structure to match the devices Datastore records
type Device struct {
	// Stored data
	DvID        string    `datastore:"dvID"`
	Created     time.Time `datastore:"created"`
	Label       string    `datastore:"label"`
	Description string    `datastore:"description"`
	Type        int32     `datastore:"type"`

	Priority    int32  `datastore:"priority"`
	IsTwoWay    bool   `datastore:"isTwoWay"`
	Category    string `datastore:"category"`
	Destination string `datastore:"destination"`

	// no indexing
	Settings   string `datastore:"settings,noindex"`
	Icon       string `datastore:"icon,noindex"`
	RawRequest string `datastore:"rawRequest,noindex"`
	// key
	ID int64 `datastore:"-"` // The integer ID used in the datastore.
}
