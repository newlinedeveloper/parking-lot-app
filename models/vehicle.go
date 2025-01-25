package models

// VehicleType enum
type VehicleType int

const (
	Bike VehicleType = iota
	Car
	Truck
)

// Vehicle struct
type Vehicle struct {
	LicensePlate string
	VehicleType  VehicleType
}
