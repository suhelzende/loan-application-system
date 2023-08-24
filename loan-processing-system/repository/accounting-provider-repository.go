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
}

func (ap AccountingProvider) GetAllAccountingProviders() ([]*model.AccountingProvider, error) {
	accountingProvider, err := storage.GetAllAccountingProviders()
	if err != nil {
		return nil, err
	}
	return accountingProvider, nil
}

func (ap AccountingProvider) GetBalenceSheet(businessDetails model.BusinessDetails, accountingProvider model.AccountingProvider) (*model.BalenceSheet, error) {
	balenceSheet, err := accountingsystem.GetBalenceSheet(businessDetails, accountingProvider)
	if err != nil {
		return nil, err
	}
	return balenceSheet, nil
}
