package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// ParkingLot struct
type ParkingLot struct {
	Levels []*ParkingLevel
}

// NewParkingLot initializes a new parking lot
func NewParkingLot(numLevels, smallPerLevel, mediumPerLevel, largePerLevel int) *ParkingLot {
	levels := []*ParkingLevel{}
	for i := 0; i < numLevels; i++ {
		level := &ParkingLevel{ID: i + 1}
		for j := 0; j < smallPerLevel; j++ {
			level.Spots = append(level.Spots, &ParkingSpot{ID: fmt.Sprintf("S-%d-%d", i+1, j+1), SpotType: Small})
		}
		for j := 0; j < mediumPerLevel; j++ {
			level.Spots = append(level.Spots, &ParkingSpot{ID: fmt.Sprintf("M-%d-%d", i+1, j+1), SpotType: Medium})
		}
		for j := 0; j < largePerLevel; j++ {
			level.Spots = append(level.Spots, &ParkingSpot{ID: fmt.Sprintf("L-%d-%d", i+1, j+1), SpotType: Large})
		}
		levels = append(levels, level)
	}
	return &ParkingLot{Levels: levels}
}

// ParkVehicle parks a vehicle in the parking lot
func (pl *ParkingLot) ParkVehicle(vehicle *Vehicle) *ParkingTicket {
	for _, level := range pl.Levels {
		spot := level.FindAvailableSpot(vehicle.VehicleType)
		if spot != nil {
			spot.AssignVehicle(vehicle)
			ticket := &ParkingTicket{
				ID:        uuid.New().String(),
				SpotID:    spot.ID,
				Vehicle:   vehicle,
				EntryTime: time.Now(),
			}
			fmt.Printf("Vehicle parked: %v\n", ticket)
			return ticket
		}
	}
	fmt.Println("No available spots for this vehicle.")
	return nil
}

// FreeSpot frees a parking spot
func (pl *ParkingLot) FreeSpot(ticket *ParkingTicket) {
	for _, level := range pl.Levels {
		for _, spot := range level.Spots {
			if spot.ID == ticket.SpotID && spot.IsOccupied && spot.Vehicle == ticket.Vehicle {
				spot.RemoveVehicle()
				fmt.Printf("Spot %s freed.\n", spot.ID)
				return
			}
		}
	}
	fmt.Println("Invalid ticket or spot already free.")
}
