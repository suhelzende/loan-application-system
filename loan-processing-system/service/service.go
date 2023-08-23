package service

import "github.com/suhelz/loan-processing-system/model"

type LoanApplicationServiceInterface interface {
	StartNewApplication(request model.LoanApplicationRequest) (*model.LoanApplication, error)

	SubmitApplication(request model.SubmitLoanApplicationRequest) (model.LoanApplication, error)

	GetApplicationByID(loanID string) (*model.LoanApplication, error)
}

type AccountProviderServiceInterface interface {
	GetAllAccountingProviders() ([]*model.AccountingProvider, error)

	GetBalenceSheet(businessDetails model.BusinessDetails, accountingProvider model.AccountingProvider) (*model.BalenceSheet, error)
}
