package model

import "time"

type LoanApplication struct {
	ID              string
	Borrower        Borrower
	DateInitiated   time.Time
	LoanAmount      int
	BusinessDetails *BusinessDetails
	Status          string
	LastModified    time.Time
}

type BusinessDetails struct {
	RegistrationID  string
	Name            string
	EstablishedYear int
}

type Borrower struct {
	Name  string
	Email string
}

type DicisionEngineRequest struct {
	BusinessDetails    BusinessDetailsWithSummary
	PreAssessmentValue int
}

type BusinessDetailsWithSummary struct {
	Name                    string
	YearEstablished         int
	ProfitLossSummaryByYear []ProfitLossSummary
}

func NewDicisionEngineRequest(businessDetails BusinessDetails, profitLossSummary []ProfitLossSummary, preAssessmentValue int) DicisionEngineRequest {
	return DicisionEngineRequest{
		BusinessDetails: BusinessDetailsWithSummary{
			Name:                    businessDetails.Name,
			YearEstablished:         businessDetails.EstablishedYear,
			ProfitLossSummaryByYear: profitLossSummary,
		},
		PreAssessmentValue: preAssessmentValue,
	}
}
