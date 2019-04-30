package datastore

import "time"

//Callpoint contains the structure to work with
type Callpoint struct {
	CpID        string    `datastore:"cpId"`
	Created     time.Time `datastore:"created"`
	AbsAddress  int32     `datastore:"absAddress"`
	Label       string    `datastore:"label"`
	Description string    `datastore:"description"`
	Type        int32     `datastore:"type"`
	Priority    int32     `datastore:"priority"`

	Icon       string `datastore:"icon,noindex"`
	RawRequest string `datastore:"rawRequest,noindex"`
	id         int64  // The integer ID used in the datastore.
}
