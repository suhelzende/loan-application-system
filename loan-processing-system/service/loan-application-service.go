package service

import (
	appconstants "github.com/suhelz/loan-processing-system/constants/app-constants"
	"github.com/suhelz/loan-processing-system/model"
	"github.com/suhelz/loan-processing-system/repository"
)

// Create LoanApplicationService - it implements LoanApplicationServiceInterface behavior
func CreateNewApplicationService(loanApplicationRepository repository.LoanApplicationRepositoryInterface) LoanApplicationServiceInterface {
	return LoanApplicationService{
		repository: loanApplicationRepository,
	}
}

// LoanApplicationService - it implements LoanApplicationServiceInterface behavior
type LoanApplicationService struct {
	repository repository.LoanApplicationRepositoryInterface
}

// Start new application
func (las LoanApplicationService) StartNewApplication(request model.LoanApplicationRequest) (*model.LoanApplication, error) {
	application := &model.LoanApplication{
		Borrower:      request.BorroweDetails,
		Status:        "PENDING",
		DateInitiated: request.Date,
	}
	return las.repository.StartNewApplication(application)
}

// Submit application
func (las LoanApplicationService) SubmitApplication(request model.SubmitLoanApplicationRequest) (model.LoanApplication, error) {
	loanApplication := request.LoanDetails
	// calculate preAssessment
	preAssessmentScore := preAssessment(request)

	// calculate profit loss by year
	profitLossSummaryByYear := profitLossSummaryByYear(request.BalanceSheet.Sheet)

	status, err := las.repository.SubmitApplication(*request.LoanDetails.BusinessDetails, profitLossSummaryByYear, preAssessmentScore)
	if err != nil {
		return loanApplication, err
	}

	loanApplication.Status = status
	// Update application with new status
	err = las.repository.UpdateApplication(&loanApplication)
	if err != nil {
		return loanApplication, err
	}

	return loanApplication, nil
}

// Calculate preAssessment score for LoanApplication
func preAssessment(request model.SubmitLoanApplicationRequest) int {
	score := appconstants.PreAssessmentScoreDefault

	profitOrLoss := 0
	totalAssetCountFor12Month := 0
	for _, balance := range request.BalanceSheet.Sheet {
		profitOrLoss += balance.ProfitOrLoss
		totalAssetCountFor12Month += balance.AssetsValue
	}

	if profitOrLoss > 0 {
		score = appconstants.PreAssessmentScoreProfit
	}
	averageAssetCount := totalAssetCountFor12Month / len(request.BalanceSheet.Sheet)

	if averageAssetCount > request.LoanDetails.LoanAmount {
		score = appconstants.PreAssessmentScoreAsset
	}
	return score
}

// Calculate profit loss summary by year
func profitLossSummaryByYear(balanceSheet []model.BalanceSheetForMonth) []model.ProfitLossSummary {
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

// return Application for given ID, else return nil
func (las LoanApplicationService) GetApplicationByID(loanID string) (*model.LoanApplication, error) {
	application, err := las.repository.GetApplicationByID(loanID)
	return application, err
}
