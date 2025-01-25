package main

import (
	"fmt"
	"parking-lot/models"
	"parking-lot/services"
)

func main() {
	// Initialize the parking lot
	parkingLot := models.NewParkingLot(2, 2, 2, 2)

	// Initialize parking service
	parkingService := services.NewParkingService(parkingLot)

	// Create vehicles
	bike := &models.Vehicle{LicensePlate: "BIKE123", VehicleType: models.Bike}
	car := &models.Vehicle{LicensePlate: "CAR456", VehicleType: models.Car}
	truck := &models.Vehicle{LicensePlate: "TRUCK789", VehicleType: models.Truck}

	// Park vehicles
	ticket1 := parkingService.ParkVehicle(bike)
	_ = parkingService.ParkVehicle(car)
	_ = parkingService.ParkVehicle(truck)

	// Display parking lot status
	fmt.Println("\nParking Lot Status:")
	parkingService.GetParkingStatus()

	// Free a spot
	if ticket1 != nil {
		parkingService.FreeSpot(ticket1)
	}

	// Display parking lot status after freeing a spot
	fmt.Println("\nParking Lot Status After Freeing a Spot:")
	parkingService.GetParkingStatus()
}
