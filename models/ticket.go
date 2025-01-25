package models

import "time"

// ParkingTicket struct
type ParkingTicket struct {
	ID        string
	SpotID    string
	Vehicle   *Vehicle
	EntryTime time.Time
}
