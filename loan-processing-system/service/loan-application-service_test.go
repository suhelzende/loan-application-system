package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	appconstants "github.com/suhelz/loan-processing-system/constants/app-constants"
	"github.com/suhelz/loan-processing-system/model"
)

func TestPreAssessment(t *testing.T) {
	testCase := []struct {
		Name string
		test func(t *testing.T)
	}{
		{
			Name: "Test 1 : Test PreAssessment 20",
			test: testPreAssessment20,
		},
		{
			Name: "Test 2 : Test PreAssessment 60",
			test: testPreAssessment60,
		},
		{
			Name: "Test 3 : Test PreAssessment 100",
			test: testPreAssessment100,
		},
	}
	for _, tc := range testCase {
		t.Run(tc.Name, tc.test)
	}
}

func testPreAssessment20(t *testing.T) {

	sheet := []model.BalanceSheetForMonth{
		{
			Year:         2023,
			Month:        8,
			ProfitOrLoss: 2000,
			AssetsValue:  1000,
		},
		{
			Year:         2023,
			Month:        7,
			ProfitOrLoss: -1200,
			AssetsValue:  1000,
		},
		{
			Year:         2023,
			Month:        6,
			ProfitOrLoss: -300,
			AssetsValue:  1000,
		},
		{
			Year:         2023,
			Month:        5,
			ProfitOrLoss: -500,
			AssetsValue:  1000,
		},
	}
	request := model.SubmitLoanApplicationRequest{
		LoanDetails: model.LoanApplication{
			LoanAmount: 5000,
		},
		BalanceSheet: model.BalenceSheet{
			Sheet: sheet,
		},
	}
	score := preAssessment(request)
	assert.Equal(t, appconstants.PreAssessmentScoreDefault, score)
}

func testPreAssessment60(t *testing.T) {

	sheet := []model.BalanceSheetForMonth{
		{
			Year:         2023,
			Month:        8,
			ProfitOrLoss: 3000,
			AssetsValue:  1000,
		},
		{
			Year:         2023,
			Month:        7,
			ProfitOrLoss: -1200,
			AssetsValue:  1000,
		},
		{
			Year:         2023,
			Month:        6,
			ProfitOrLoss: -300,
			AssetsValue:  1000,
		},
		{
			Year:         2023,
			Month:        5,
			ProfitOrLoss: -500,
			AssetsValue:  1000,
		},
	}
	request := model.SubmitLoanApplicationRequest{
		LoanDetails: model.LoanApplication{
			LoanAmount: 5000,
		},
		BalanceSheet: model.BalenceSheet{
			Sheet: sheet,
		},
	}
	score := preAssessment(request)
	assert.Equal(t, appconstants.PreAssessmentScoreProfit, score)
}

func testPreAssessment100(t *testing.T) {

	sheet := []model.BalanceSheetForMonth{
		{
			Year:         2023,
			Month:        8,
			ProfitOrLoss: 3000,
			AssetsValue:  5000,
		},
		{
			Year:         2023,
			Month:        7,
			ProfitOrLoss: -1200,
			AssetsValue:  4000,
		},
		{
			Year:         2023,
			Month:        6,
			ProfitOrLoss: -300,
			AssetsValue:  7000,
		},
		{
			Year:         2023,
			Month:        5,
			ProfitOrLoss: -500,
			AssetsValue:  6000,
		},
	}
	request := model.SubmitLoanApplicationRequest{
		LoanDetails: model.LoanApplication{
			LoanAmount: 5000,
		},
		BalanceSheet: model.BalenceSheet{
			Sheet: sheet,
		},
	}
	score := preAssessment(request)
	assert.Equal(t, appconstants.PreAssessmentScoreAsset, score)
}

func TestProfitLossSummaryByYear(t *testing.T) {
	sheet := []model.BalanceSheetForMonth{
		{
			Year:         2023,
			Month:        2,
			ProfitOrLoss: 3000,
			AssetsValue:  5000,
		},
		{
			Year:         2023,
			Month:        1,
			ProfitOrLoss: -1200,
			AssetsValue:  4000,
		},
		{
			Year:         2022,
			Month:        12,
			ProfitOrLoss: -300,
			AssetsValue:  7000,
		},
		{
			Year:         2022,
			Month:        11,
			ProfitOrLoss: -500,
			AssetsValue:  6000,
		},
	}

	expected := []model.ProfitLossSummary{
		{
			Year:         2023,
			ProfitOrLoss: 1800,
		},
		{
			Year:         2022,
			ProfitOrLoss: -800,
		},
	}

	summary := profitLossSummaryByYear(sheet)
	assert.EqualValues(t, expected, summary)
}
