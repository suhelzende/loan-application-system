package service

import (
	"github.com/suhelz/loan-processing-system/model"
	"github.com/suhelz/loan-processing-system/repository"
)

func CreateNewAccountingProviderService(accountingProvider repository.AccountingProviderServiceInterface) AccountProviderServiceInterface {
	return AccountingProviderService{
		repository: accountingProvider,
	}
}

// AccountingProviderService - it implements AccountProviderServiceInterface behavior
type AccountingProviderService struct {
	repository repository.AccountingProviderServiceInterface
}

func (as AccountingProviderService) GetAllAccountingProviders() ([]*model.AccountingProvider, error) {
	accountingProviders, err := as.repository.GetAllAccountingProviders()
	if err != nil {
		return nil, err
	}
	return accountingProviders, nil
}

func (as AccountingProviderService) GetBalenceSheet(businessDetails model.BusinessDetails, accountingProvider model.AccountingProvider) (*model.BalenceSheet, error) {
	sheet, err := as.repository.GetBalenceSheet(businessDetails, accountingProvider)
	if err != nil {
		return nil, err
	}
	return sheet, nil
}
