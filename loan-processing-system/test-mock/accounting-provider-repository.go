package mock

import (
	"github.com/suhelz/loan-processing-system/model"
)

// it implements AccountingProviderServiceInterface behavior
type AccountingProviderRepositoryMock struct {
	TestHelper
}

func (lap AccountingProviderRepositoryMock) GetAllAccountingProviders() ([]*model.AccountingProvider, error) {
	args := lap.Called()
	arg1, err := args.Get(0), args.Error(1)
	if arg1 == nil {
		return nil, err
	}
	return arg1.([]*model.AccountingProvider), err
}

func (lap AccountingProviderRepositoryMock) GetBalenceSheet(businessDetails model.BusinessDetails, accountingProvider model.AccountingProvider) (*model.BalenceSheet, error) {
	args := lap.Called(businessDetails, accountingProvider)
	arg1, err := args.Get(0), args.Error(1)
	if arg1 == nil {
		return nil, err
	}
	return arg1.(*model.BalenceSheet), err
}
