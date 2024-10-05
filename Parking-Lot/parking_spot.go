package main

import "fmt"

type ParkingSpot struct {
	number          int
	parkingSpotType ParkingSpotType
	free            bool
	vehicle         interface{}
}

func NewParkingSpot(number int, parkingSpotType ParkingSpotType) *ParkingSpot {
	return &ParkingSpot{
		number:          number,
		parkingSpotType: parkingSpotType,
		free:            true,
		vehicle:         nil,
	}
}

func (ps *ParkingSpot) IsFree() bool {
	return ps.free
}

func (ps *ParkingSpot) GetNumber() string {
	return fmt.Sprintf("%d", ps.number)
}
func (p *ParkingSpot) AssignVehicle(vehicle interface{}) {
	p.vehicle = vehicle
	p.free = false
}

func (p *ParkingSpot) RemoveVehicle() {
	p.vehicle = nil
	p.free = true
}

type HandicappedSpot struct {
	*ParkingSpot
}

func NewHandicappedSpot(number int) *HandicappedSpot {
	return &HandicappedSpot{
		ParkingSpot: NewParkingSpot(number, HANDICAPPED),
	}
}

type CompactSpot struct {
	*ParkingSpot
}

func NewCompactSpot(number int) *CompactSpot {
	return &CompactSpot{
		ParkingSpot: NewParkingSpot(number, COMPACT),
	}
}

type LargeSpot struct {
	*ParkingSpot
}

// NewLargeSpot creates a new large parking spot.
func NewLargeSpot(number int) *LargeSpot {
	return &LargeSpot{
		ParkingSpot: NewParkingSpot(number, LARGE),
	}
}

// MotorbikeSpot represents a motorbike parking spot.
type MotorbikeSpot struct {
	*ParkingSpot
}

func NewMotorbikeSpot(number int) *MotorbikeSpot {
	return &MotorbikeSpot{
		ParkingSpot: NewParkingSpot(number, MOTORBIKE_SPOT),
	}
}

type ElectricSpot struct {
	*ParkingSpot
}

func NewElectricSpot(number int) *ElectricSpot {
	return &ElectricSpot{
		ParkingSpot: NewParkingSpot(number, ELECTRIC_SPOT),
	}
}
