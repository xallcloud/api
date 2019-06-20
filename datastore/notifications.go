package datastore

import "time"

const (
	// KindNotifications represents the "table" name (Kind) to store Notifications
	KindNotifications = "notifications"
)

// Notification contains the structure to match the notifications Datastore records
type Notification struct {
	// key
	ID int64 `datastore:"-"` // The integer ID used in the datastore.
	// Stored data
	NtID          string    `datastore:"ntID"`
	AcID          string    `datastore:"acID"`
	Priority      int32     `datastore:"priority"`
	Category      string    `datastore:"category"`
	Destination   string    `datastore:"destination"`
	Message       string    `datastore:"message"`
	ResponseTitle string    `datastore:"responseTitle"`
	Options       string    `datastore:"options"`
	Created       time.Time `datastore:"created"`
	CallpointObj  Callpoint `datastore:"-"`
	DeviceObj     Device    `datastore:"-"`
}
