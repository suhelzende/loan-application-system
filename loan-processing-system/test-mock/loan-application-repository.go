package mock

import (
	"github.com/suhelz/loan-processing-system/model"
)

// it implements LoanApplicationRepositoryInterface behavior
type LoanApplicationRepositoryMock struct {
	TestHelper
}

func (lap LoanApplicationRepositoryMock) StartNewApplication(application *model.LoanApplication) (*model.LoanApplication, error) {
	args := lap.Called(application)
	arg1, err := args.Get(0), args.Error(1)
	if arg1 == nil {
		return nil, err
	}
	return arg1.(*model.LoanApplication), err
}

func (lap LoanApplicationRepositoryMock) UpdateApplication(application *model.LoanApplication) error {
	args := lap.Called(application)
	return args.Error(0)
}

func (lap LoanApplicationRepositoryMock) SubmitApplication(businessDetails model.BusinessDetails, profitLossSummary []model.ProfitLossSummary, preAssessmentScore int) (string, error) {
	args := lap.Called(businessDetails, profitLossSummary, preAssessmentScore)
	arg1, err := args.Get(0), args.Error(1)
	if arg1 == nil {
		return "", err
	}
	return arg1.(string), err
}
