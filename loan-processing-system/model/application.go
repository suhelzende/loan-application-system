package model

type LoanApplication struct {
	ID              string
	Requester       Borrowe
	BusinessDetails *BusinessDetails
	Status          string
}

type BusinessDetails struct {
	RegistrationID  string
	Name            string
	EstablishedYear int
}

type Borrowe struct {
	Name  string
	Email string
}
