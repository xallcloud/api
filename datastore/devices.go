package datastore

import "time"

const (
	// KindDevices represents the "table" name (Kind) to store Devices
	KindDevices = "devices"
)

// Device contains the structure to match the devices Datastore records
type Device struct {
	// key
	ID int64 `datastore:"-"` // The integer ID used in the datastore.
	// Stored data
	DvID        string    `datastore:"dvID"`
	Label       string    `datastore:"label"`
	Type        int32     `datastore:"type"`
	Icon        string    `datastore:"icon,noindex"`
	Description string    `datastore:"description"`
	IsTwoWay    bool      `datastore:"isTwoWay"`
	Category    string    `datastore:"category"`
	Destination string    `datastore:"destination"`
	Priority    int32     `datastore:"priority"`
	Created     time.Time `datastore:"created"`
	Changed     time.Time `datastore:"changed"`
	Settings    string    `datastore:"settings,noindex"`
	RawRequest  string    `datastore:"rawRequest,noindex"`
}
