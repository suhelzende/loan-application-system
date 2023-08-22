package service

import "github.com/suhelz/loan-processing-system/model"

type LoanApplicationServiceInterface interface {
	StartNewApplication(request model.LoanApplicationRequest) (*model.LoanApplication, error)
	SubmitApplication(request model.LoanApplication) (*model.LoanApplication, error)
	GetApplicationByID(loanID string) (*model.LoanApplication, error)
}
