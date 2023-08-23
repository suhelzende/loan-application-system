package repository

import (
	dicisionengine "github.com/suhelz/loan-processing-system/dicision-engine"
	"github.com/suhelz/loan-processing-system/model"
	"github.com/suhelz/loan-processing-system/storage"
)

func CreateNewApplicationRepository(config model.Config) LoanApplicationRepositoryInterface {
	return LoanApplicationRepository{}
}

type LoanApplicationRepository struct {
}

func (lap LoanApplicationRepository) StartNewApplication(application *model.LoanApplication) (*model.LoanApplication, error) {
	newApplication := storage.CreateNewLoanApplication(application)
	return newApplication, nil
}

func (lap LoanApplicationRepository) UpdateApplication(application *model.LoanApplication) error {
	return storage.UpdateLoanApplication(application)
}

func (lap LoanApplicationRepository) SubmitApplication(businessDetails model.BusinessDetails, profitLossSummary []model.ProfitLossSummary, preAssessmentScore int) (string, error) {
	dicisionEngineRequest := model.NewDicisionEngineRequest(businessDetails, profitLossSummary, preAssessmentScore)
	return dicisionengine.SubmitApplication(dicisionEngineRequest)
}

func (lap LoanApplicationRepository) GetApplicationByID(loanID string) (*model.LoanApplication, error) {
	application := storage.GetApplicationByID(loanID)
	return application, nil
}
