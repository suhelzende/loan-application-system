package repository

import (
	accountingsystem "github.com/suhelz/loan-processing-system/accounting-system"
	"github.com/suhelz/loan-processing-system/model"
	"github.com/suhelz/loan-processing-system/storage"
)

func CreateNewAccountingProviderRepository() AccountingProviderServiceInterface {
	return AccountingProvider{}
}

type AccountingProvider struct {
	// place holder for db and third part API integration
}

// Returns All Accounting Provider available
func (ap AccountingProvider) GetAllAccountingProviders() ([]*model.AccountingProvider, error) {
	accountingProvider, err := storage.GetAllAccountingProviders()
	if err != nil {
		return nil, err
	}
	return accountingProvider, nil
}

// Returns Balance sheet for business details from accounting provider
func (ap AccountingProvider) GetBalenceSheet(businessDetails model.BusinessDetails, accountingProvider model.AccountingProvider) (*model.BalenceSheet, error) {
	balanceSheet, err := accountingsystem.GetBalanceSheet(businessDetails, accountingProvider)
	if err != nil {
		return nil, err
	}
	return balanceSheet, nil
}
