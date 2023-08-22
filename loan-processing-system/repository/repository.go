package repository

import "github.com/suhelz/loan-processing-system/model"

func CreateNewApplicationRepository(config model.Config) LoanApplicationRepositoryInterface {
	return nil
}

type LoanApplicationRepository struct {
}

func (lap LoanApplicationRepository) StartNewApplication(request model.LoanApplicationRequest) (*model.LoanApplication, error) {
	// TODO: Implement this function
	return nil, nil
}

func (lap LoanApplicationRepository) SubmitApplication(request model.LoanApplication) (*model.LoanApplication, error) {
	// TODO: Implement this function
	return nil, nil
}

func (lap LoanApplicationRepository) GetApplicationByID(loanID string) (*model.LoanApplication, error) {
	// TODO: Implement this function
	return nil, nil
}
