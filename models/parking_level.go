package models

// ParkingLevel struct
type ParkingLevel struct {
	ID    int
	Spots []*ParkingSpot
}

// FindAvailableSpot finds an available spot for a vehicle
func (pl *ParkingLevel) FindAvailableSpot(vehicleType VehicleType) *ParkingSpot {
	for _, spot := range pl.Spots {
		if !spot.IsOccupied && canVehicleFit(vehicleType, spot.SpotType) {
			return spot
		}
	}
	return nil
}

// Helper function to check if a vehicle can fit in a spot
func canVehicleFit(vehicleType VehicleType, spotType SpotType) bool {
	switch vehicleType {
	case Bike:
		return true
	case Car:
		return spotType == Medium || spotType == Large
	case Truck:
		return spotType == Large
	default:
		return false
	}
}
