package models

import (
	"sync"
)

// SpotType enum
type SpotType int

const (
	Small SpotType = iota
	Medium
	Large
)

// ParkingSpot struct
type ParkingSpot struct {
	ID         string
	SpotType   SpotType
	IsOccupied bool
	Vehicle    *Vehicle
	mu         sync.Mutex
}

// AssignVehicle assigns a vehicle to the spot
func (ps *ParkingSpot) AssignVehicle(vehicle *Vehicle) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.Vehicle = vehicle
	ps.IsOccupied = true
}

// RemoveVehicle removes a vehicle from the spot
func (ps *ParkingSpot) RemoveVehicle() {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.Vehicle = nil
	ps.IsOccupied = false
}
