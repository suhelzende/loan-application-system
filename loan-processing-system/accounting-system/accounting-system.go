package accountingsystem

// This is mock placeholder which acts as Accounting provider for loan-application
import (
	"encoding/json"

	"github.com/suhelz/loan-processing-system/model"
)

const (
	mockData = `
	{
		"sheet": [
		  {
			"year": 2020,
			"month": 12,
			"profitOrLoss": 250000,
			"assetsValue": 1234
		  },
		  {
			"year": 2020,
			"month": 11,
			"profitOrLoss": 1150,
			"assetsValue": 5789
		  },
		  {
			"year": 2020,
			"month": 10,
			"profitOrLoss": 2500,
			"assetsValue": 22345
		  },
		  {
			"year": 2020,
			"month": 9,
			"profitOrLoss": -187000,
			"assetsValue": 223452
		  },
		  {
			"year": 2020,
			"month": 8,
			"profitOrLoss": 250000,
			"assetsValue": 1234
		  },
		  {
			"year": 2020,
			"month": 7,
			"profitOrLoss": 1150,
			"assetsValue": 5789
		  },
		  {
			"year": 2020,
			"month": 6,
			"profitOrLoss": 2500,
			"assetsValue": 22345
		  },
		  {
			"year": 2020,
			"month": 5,
			"profitOrLoss": -187000,
			"assetsValue": 223452
		  },
		  {
			"year": 2020,
			"month": 4,
			"profitOrLoss": 250000,
			"assetsValue": 1234
		  },
		  {
			"year": 2020,
			"month": 3,
			"profitOrLoss": 1150,
			"assetsValue": 5789
		  },
		  {
			"year": 2020,
			"month": 2,
			"profitOrLoss": 2500,
			"assetsValue": 22345
		  },
		  {
			"year": 2020,
			"month": 1,
			"profitOrLoss": -187000,
			"assetsValue": 223452
		  }
		]
	  }
	  
	`
)

// As per the given problem statement, decision engine is third party system
// So, this acts as decision Engine mock
// we can integrate decision engine APIs here
func GetBalanceSheet(businessDetails model.BusinessDetails, accountingProvider model.AccountingProvider) (*model.BalenceSheet, error) {

	sheet := &model.BalenceSheet{}
	err := json.Unmarshal([]byte(mockData), sheet)
	if err != nil {
		return nil, err
	}

	return sheet, nil
}
