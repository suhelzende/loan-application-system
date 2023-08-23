package service

import (
	appconstants "github.com/suhelz/loan-processing-system/constants/app-constants"
	"github.com/suhelz/loan-processing-system/model"
	"github.com/suhelz/loan-processing-system/repository"
)

func CreateNewApplicationService(loanApplicationRepository repository.LoanApplicationRepositoryInterface) LoanApplicationServiceInterface {
	return LoanApplicationService{
		repository: loanApplicationRepository,
	}
}

type LoanApplicationService struct {
	repository repository.LoanApplicationRepositoryInterface
}

func (las LoanApplicationService) StartNewApplication(request model.LoanApplicationRequest) (*model.LoanApplication, error) {
	application := &model.LoanApplication{
		Borrower:      request.BorroweDetails,
		Status:        "PENDING",
		DateInitiated: request.Date,
	}
	return las.repository.StartNewApplication(application)
}

func (las LoanApplicationService) SubmitApplication(request model.SubmitLoanApplicationRequest) (model.LoanApplication, error) {
	loanApplication := request.LoanDetails
	preAssessmentScore := preAssessment(request)
	profitLossSummaryByYear := profitLossSummaryByYear(request.BalenceSheet.Sheet)
	status, err := las.repository.SubmitApplication(*request.LoanDetails.BusinessDetails, profitLossSummaryByYear, preAssessmentScore)
	if err != nil {
		return loanApplication, err
	}

	loanApplication.Status = status
	err = las.repository.UpdateApplication(&loanApplication)
	if err != nil {
		return loanApplication, err
	}

	return loanApplication, nil
}

func preAssessment(request model.SubmitLoanApplicationRequest) int {
	score := appconstants.PreAssessmentScoreDefault

	profitOrLoss := 0
	totalAssetCountFor12Monht := 0
	for _, balance := range request.BalenceSheet.Sheet {
		profitOrLoss += balance.ProfitOrLoss
		totalAssetCountFor12Monht += balance.AssetsValue
	}

	if profitOrLoss > 0 {
		score = appconstants.PreAssessmentScoreProfit
	}
	averageAssetCount := totalAssetCountFor12Monht / len(request.BalenceSheet.Sheet)

	if averageAssetCount > request.LoanDetails.LoanAmount {
		score = appconstants.PreAssessmentScoreAsset
	}
	return score
}

func profitLossSummaryByYear(balanceSheet []model.BalenceSheetForMonth) []model.ProfitLossSummary {
	profitLossByYear := make(map[int]*model.ProfitLossSummary)

	for _, balance := range balanceSheet {
		_, found := profitLossByYear[balance.Year]
		if !found {
			profitLossByYear[balance.Year] = &model.ProfitLossSummary{
				Year:         balance.Year,
				ProfitOrLoss: balance.ProfitOrLoss,
			}
		} else {
			profitLossByYear[balance.Year].ProfitOrLoss += balance.ProfitOrLoss
		}
	}

	finalSummary := make([]model.ProfitLossSummary, 0)
	for _, summary := range profitLossByYear {
		finalSummary = append(finalSummary, *summary)
	}

	return finalSummary
}

func (las LoanApplicationService) GetApplicationByID(loanID string) (*model.LoanApplication, error) {
	application, err := las.repository.GetApplicationByID(loanID)
	return application, err
}
