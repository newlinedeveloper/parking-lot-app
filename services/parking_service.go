package services

import (
	"fmt"
	"parking-lot/models"
)

// ParkingService struct
type ParkingService struct {
	ParkingLot *models.ParkingLot
}

// NewParkingService initializes a new parking service
func NewParkingService(parkingLot *models.ParkingLot) *ParkingService {
	return &ParkingService{ParkingLot: parkingLot}
}

// ParkVehicle parks a vehicle and returns a ticket
func (ps *ParkingService) ParkVehicle(vehicle *models.Vehicle) *models.ParkingTicket {
	ticket := ps.ParkingLot.ParkVehicle(vehicle)
	if ticket != nil {
		fmt.Printf("Vehicle %s parked at spot %s\n", vehicle.LicensePlate, ticket.SpotID)
		return ticket
	}
	fmt.Printf("No available spot for vehicle %s\n", vehicle.LicensePlate)
	return nil
}

// FreeSpot releases a parking spot based on the ticket
func (ps *ParkingService) FreeSpot(ticket *models.ParkingTicket) {
	ps.ParkingLot.FreeSpot(ticket)
}

// GetParkingStatus retrieves the current status of the parking lot
func (ps *ParkingService) GetParkingStatus() {
	for _, level := range ps.ParkingLot.Levels {
		fmt.Printf("Level %d:\n", level.ID)
		for _, spot := range level.Spots {
			status := "Free"
			if spot.IsOccupied {
				status = fmt.Sprintf("Occupied by %s", spot.Vehicle.LicensePlate)
			}
			fmt.Printf("  Spot %s (%v): %s\n", spot.ID, spot.SpotType, status)
		}
	}
}
