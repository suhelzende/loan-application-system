package service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/suhelz/loan-processing-system/model"
	mock "github.com/suhelz/loan-processing-system/test-mock"
)

func TestGetAllAccountingProviders(t *testing.T) {
	testCase := []struct {
		Name string
		test func(t *testing.T)
	}{
		{
			Name: "Test 1 : Test Accounting Provider Error",
			test: testGetAllAccountingProvidersError,
		},
		{
			Name: "Test 2 : Test Accounting Provider Success",
			test: testGetAllAccountingProvidersPass,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.Name, tc.test)
	}
}

func testGetAllAccountingProvidersPass(t *testing.T) {
	repoMock := &mock.AccountingProviderRepositoryMock{}

	expectedAccountingProvider := []*model.AccountingProvider{
		{
			ID:   "TEST1",
			Name: "Test 1",
		},
		{
			ID:   "TEST1",
			Name: "Test 1",
		},
	}
	repoMock.On("GetAllAccountingProviders").Return(expectedAccountingProvider, nil)

	service := CreateNewAccountingProviderService(repoMock)
	result, err := service.GetAllAccountingProviders()
	assert.NoError(t, err)
	assert.EqualValues(t, expectedAccountingProvider, result)
}

func testGetAllAccountingProvidersError(t *testing.T) {
	repoMock := &mock.AccountingProviderRepositoryMock{}

	repoMock.On("GetAllAccountingProviders").Return(nil, errors.New("failed"))

	service := CreateNewAccountingProviderService(repoMock)
	result, err := service.GetAllAccountingProviders()
	assert.Nil(t, result)
	assert.Error(t, err)
}

func TestGetBalanceSheet(t *testing.T) {
	testCase := []struct {
		Name string
		test func(t *testing.T)
	}{
		{
			Name: "Test 1 : Test Get BalanceSheet Error",
			test: testGetBalanceSheetError,
		},
		{
			Name: "Test 2 : Test Get BalanceSheet Success",
			test: testGetBalanceSheetPass,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.Name, tc.test)
	}
}

func testGetBalanceSheetPass(t *testing.T) {
	repoMock := &mock.AccountingProviderRepositoryMock{}

	inputBusinessDetail := model.BusinessDetails{
		RegistrationID:  "TESTINGREGISTRATIONID",
		Name:            "Test Business",
		EstablishedYear: 2023,
	}
	inputAccountProvider := model.AccountingProvider{
		ID:   "TESTACCOUNTPROVIDER",
		Name: "test account provider",
	}
	expectedBalanceSheet := &model.BalenceSheet{
		Sheet: []model.BalanceSheetForMonth{
			{
				Year:         2023,
				Month:        8,
				ProfitOrLoss: 222081,
				AssetsValue:  21212,
			},
			{
				Year:         2023,
				Month:        7,
				ProfitOrLoss: 222081,
				AssetsValue:  21212,
			},
			{
				Year:         2023,
				Month:        6,
				ProfitOrLoss: 222081,
				AssetsValue:  21212,
			},
			{
				Year:         2023,
				Month:        5,
				ProfitOrLoss: 222081,
				AssetsValue:  21212,
			},
		},
	}
	repoMock.On("GetBalenceSheet", inputBusinessDetail, inputAccountProvider).Return(expectedBalanceSheet, nil)

	service := CreateNewAccountingProviderService(repoMock)
	result, err := service.GetBalenceSheet(inputBusinessDetail, inputAccountProvider)
	assert.NoError(t, err)
	assert.EqualValues(t, expectedBalanceSheet, result)
}

func testGetBalanceSheetError(t *testing.T) {
	repoMock := &mock.AccountingProviderRepositoryMock{}

	inputBusinessDetail := model.BusinessDetails{
		RegistrationID:  "TESTINGREGISTRATIONID",
		Name:            "Test Business",
		EstablishedYear: 2023,
	}
	inputAccountProvider := model.AccountingProvider{
		ID:   "TESTACCOUNTPROVIDER",
		Name: "test account provider",
	}

	repoMock.On("GetBalenceSheet", inputBusinessDetail, inputAccountProvider).Return(nil, errors.New("failed"))

	service := CreateNewAccountingProviderService(repoMock)
	result, err := service.GetBalenceSheet(inputBusinessDetail, inputAccountProvider)
	assert.Error(t, err, "Error expected to be not nil")
	assert.Nil(t, result, "result expected to be nil")
}
