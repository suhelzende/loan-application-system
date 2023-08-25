package service

import "github.com/suhelz/loan-processing-system/model"

// LoanApplicationServiceInterface - it defines behavior of LoanApplication Service
type LoanApplicationServiceInterface interface {
	StartNewApplication(request model.LoanApplicationRequest) (*model.LoanApplication, error)

	SubmitApplication(request model.SubmitLoanApplicationRequest) (model.LoanApplication, error)

	GetApplicationByID(loanID string) (*model.LoanApplication, error)
}

// AccountProviderServiceInterface - it defines behavior of AccountProvider Service
type AccountProviderServiceInterface interface {
	GetAllAccountingProviders() ([]*model.AccountingProvider, error)

	GetBalenceSheet(businessDetails model.BusinessDetails, accountingProvider model.AccountingProvider) (*model.BalenceSheet, error)
}
