package storage

// Mock DB behavior

import (
	"github.com/suhelz/loan-processing-system/model"
)

func GetAllAccountingProviders() ([]*model.AccountingProvider, error) {
	accountingProviders := make([]*model.AccountingProvider, 0)
	for _, ap := range accountingProviderStore {
		accountingProviders = append(accountingProviders, ap)
	}
	return accountingProviders, nil
}
