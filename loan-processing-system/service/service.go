package service

import (
	"github.com/suhelz/loan-processing-system/model"
	"github.com/suhelz/loan-processing-system/repository"
)

func CreateNewApplicationService(loanApplicationRepository repository.LoanApplicationRepositoryInterface) LoanApplicationServiceInterface {
	return LoanApplicationService{
		repository: loanApplicationRepository,
	}
}

type LoanApplicationService struct {
	repository repository.LoanApplicationRepositoryInterface
}

func (las LoanApplicationService) StartNewApplication(request model.LoanApplicationRequest) (*model.LoanApplication, error) {
	// TODO: Implement this function
	return nil, nil
}

func (las LoanApplicationService) SubmitApplication(request model.LoanApplication) (*model.LoanApplication, error) {
	// TODO: Implement this function
	return nil, nil
}

func (las LoanApplicationService) GetApplicationByID(loanID string) (*model.LoanApplication, error) {
	// TODO: Implement this function
	return nil, nil
}
