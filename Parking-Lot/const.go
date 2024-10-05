package main

// VehicleType represents the type of vehicles.
type VehicleType int

const (
	CAR VehicleType = iota + 1
	TRUCK
	ELECTRIC
	VAN
	MOTORBIKE
)

// ParkingSpotType represents the type of parking spots.
type ParkingSpotType int

const (
	HANDICAPPED ParkingSpotType = iota + 1
	COMPACT
	LARGE
	MOTORBIKE_SPOT
	ELECTRIC_SPOT
)

// AccountStatus represents the status of an account.
type AccountStatus int

const (
	ACTIVE AccountStatus = iota + 1
	BLOCKED
	BANNED
	COMPROMISED
	ARCHIVED
	UNKNOWN
)

// ParkingTicketStatus represents the status of a parking ticket.
type ParkingTicketStatus int

const (
	TICKET_ACTIVE ParkingTicketStatus = iota + 1
	TICKET_PAID
	TICKET_LOST
)

// Address represents a person's address.
type Address struct {
	streetAddress string
	city          string
	state         string
	zipCode       string
	country       string
}

// NewAddress creates a new address.
func NewAddress(street, city, state, zip, country string) *Address {
	return &Address{
		streetAddress: street,
		city:          city,
		state:         state,
		zipCode:       zip,
		country:       country,
	}
}

// Person represents a person with a name, address, email, and phone.
type Person struct {
	name    string
	address *Address
	email   string
	phone   string
}

// NewPerson creates a new person.
func NewPerson(name string, address *Address, email string, phone string) *Person {
	return &Person{
		name:    name,
		address: address,
		email:   email,
		phone:   phone,
	}
}
