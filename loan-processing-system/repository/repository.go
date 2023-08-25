package repository

import "github.com/suhelz/loan-processing-system/model"

// LoanApplicationRepositoryInterface - it defines behavior of LoanApplication repository
type LoanApplicationRepositoryInterface interface {
	StartNewApplication(request *model.LoanApplication) (*model.LoanApplication, error)

	UpdateApplication(request *model.LoanApplication) error

	SubmitApplication(businessDetails model.BusinessDetails, profitLossSummary []model.ProfitLossSummary, preAssessmentScore int) (string, error)

	GetApplicationByID(loanID string) (*model.LoanApplication, error)
}

// AccountingProviderServiceInterface - it defines behavior of AccountProvider Repository
type AccountingProviderServiceInterface interface {
	GetAllAccountingProviders() ([]*model.AccountingProvider, error)

	GetBalenceSheet(businessDetails model.BusinessDetails, accountingProvider model.AccountingProvider) (*model.BalenceSheet, error)
}
