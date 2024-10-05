package main

type Account struct {
	user_name string
	password  string
	person    string
	status    AccountStatus
}

func NewAccount(user_name, password, person string, status AccountStatus) *Account {
	return &Account{
		user_name: user_name,
		password:  password,
		person:    person,
		status:    status,
	}
}

func (a *Account) ResetPassword() {

}

type Admin struct {
	Account
}

func NewAdmin(user_name, password, person string, status AccountStatus) *Admin {
	return &Admin{
		Account: *NewAccount(user_name, password, person, status),
	}
}

func(a *Admin)AddParkingFloor(){}

func(a *Admin)AddParkingSpot(){}

func(a *Admin)AddParkingDisplayBoard(){}

func(a *Admin)AddCustomerInfoPanel(){}

func(a *Admin)AddEntrancePanel(){}

func(a *Admin)AddExitPanel(){}


type ParkingAttendant struct{
	Account
}

func NewParkingAttendant(user_name, password, person string, status AccountStatus)*ParkingAttendant{
	return &ParkingAttendant{
		Account: *NewAccount(user_name, password, person, status),
	}
}

func(p *ParkingAttendant)ProcessTicket(){}