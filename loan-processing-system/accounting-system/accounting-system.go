package accountingsystem

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/suhelz/loan-processing-system/model"
)

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
