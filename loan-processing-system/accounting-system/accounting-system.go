package accountingsystem

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/suhelz/loan-processing-system/model"
)

// From Problem statment it is not clear that Accounting System is third party SDK ot API
// Also it does not give much details about Accounting System
// So, this acts as Accounting system mock
// we can integrate actual accounting system here
func GetBalenceSheet(businessDetails model.BusinessDetails, accountingProvider model.AccountingProvider) (*model.BalenceSheet, error) {
	file, err := os.Open("./accounting-system/mock/balencesheet.json")
	if err != nil {
		return nil, err
	}

	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	sheet := &model.BalenceSheet{}
	err = json.Unmarshal(byteData, sheet)
	if err != nil {
		return nil, err
	}

	return sheet, nil
}
