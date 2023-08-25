package repository

import (
	decisionengine "github.com/suhelz/loan-processing-system/decision-engine"
	"github.com/suhelz/loan-processing-system/model"
	"github.com/suhelz/loan-processing-system/storage"
)

func CreateNewApplicationRepository() LoanApplicationRepositoryInterface {
	return LoanApplicationRepository{}
}

type LoanApplicationRepository struct {
	// place holder for Decision Engine API object and DB
}

// Start New Loan Application and return LoanApplication with ApplicationID
func (lap LoanApplicationRepository) StartNewApplication(application *model.LoanApplication) (*model.LoanApplication, error) {
	newApplication := storage.CreateNewLoanApplication(application)
	return newApplication, nil
}

// Update Application in storage
func (lap LoanApplicationRepository) UpdateApplication(application *model.LoanApplication) error {
	return storage.UpdateLoanApplication(application)
}

// Submit Application to decisionEngine
func (lap LoanApplicationRepository) SubmitApplication(businessDetails model.BusinessDetails, profitLossSummary []model.ProfitLossSummary, preAssessmentScore int) (string, error) {
	decisionEngineRequest := model.NewDecisionEngineRequest(businessDetails, profitLossSummary, preAssessmentScore)
	return decisionengine.SubmitApplication(decisionEngineRequest)
}

// Return Application for given ID, else return nil
func (lap LoanApplicationRepository) GetApplicationByID(loanID string) (*model.LoanApplication, error) {
	application := storage.GetApplicationByID(loanID)
	return application, nil
}
