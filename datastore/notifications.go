package datastore

import "time"

const (
	// KindNotifications represents the "table" name (Kind) to store Notifications
	KindNotifications = "notifications"
)

// Notifications contains the structure to match the notifications Datastore records
type Notifications struct {
	// key
	ID int64 `datastore:"-"` // The integer ID used in the datastore.
	// Stored data
	NtID          string    `datastore:"ntID"`
	EvID          string    `datastore:"evID"`
	Priority      int32     `datastore:"priority"`
	Category      string    `datastore:"category"`
	Destination   string    `datastore:"destination"`
	Message       string    `datastore:"message"`
	ResponseTitle string    `datastore:"responseTitle"`
	Options       string    `datastore:"options"`
	Created       time.Time `datastore:"created"`
	EvSubType     string    `datastore:"evSubType"`
	EvDescription string    `datastore:"evDescription"`
	//??
	//CallpointObj  Callpoint `datastore:"-"`
	//DeviceObj     Device    `datastore:"-"`
}
