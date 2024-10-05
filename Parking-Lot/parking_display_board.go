package main

import "fmt"

type ParkingDisplayBoard struct {
	id                  string
	handicappedFreeSpot *ParkingSpot
	compactFreeSpot     *ParkingSpot
	largeFreeSpot       *ParkingSpot
	motorbikeFreeSpot   *ParkingSpot
	electricFreeSpot    *ParkingSpot
}

func NewParkingDisplayBoard(id string) *ParkingDisplayBoard {
	return &ParkingDisplayBoard{id: id}
}

func (pdb *ParkingDisplayBoard) ShowEmptySpotNumber() {
	message := ""

	if pdb.handicappedFreeSpot != nil && pdb.handicappedFreeSpot.IsFree() {
		message += "Free Handicapped: " + pdb.handicappedFreeSpot.GetNumber() + "\n"
	} else {
		message += "Handicapped is full\n"
	}

	if pdb.compactFreeSpot != nil && pdb.compactFreeSpot.IsFree() {
		message += "Free Compact: " + pdb.compactFreeSpot.GetNumber() + "\n"
	} else {
		message += "Compact is full\n"
	}

	if pdb.largeFreeSpot != nil && pdb.largeFreeSpot.IsFree() {
		message += "Free Large: " + pdb.largeFreeSpot.GetNumber() + "\n"
	} else {
		message += "Large is full\n"
	}

	if pdb.motorbikeFreeSpot != nil && pdb.motorbikeFreeSpot.IsFree() {
		message += "Free Motorbike: " + pdb.motorbikeFreeSpot.GetNumber() + "\n"
	} else {
		message += "Motorbike is full\n"
	}

	if pdb.electricFreeSpot != nil && pdb.electricFreeSpot.IsFree() {
		message += "Free Electric: " + pdb.electricFreeSpot.GetNumber() + "\n"
	} else {
		message += "Electric is full\n"
	}

	fmt.Print(message)
}
