package storage

import (
	"github.com/suhelz/loan-processing-system/model"
)

// This package serves as mock data store for loan-application systems
// This mock DB behavior, actual application wont need this application will directly call db query

var loanApplicationsStore map[string]*model.LoanApplication
var accountingProviderStore map[string]*model.AccountingProvider

func InitStorage() {
	loanApplicationsStore = make(map[string]*model.LoanApplication)
	accountingProviderStore = make(map[string]*model.AccountingProvider)
	accountingProviderStore["Xero"] = &model.AccountingProvider{
		Name: "Xero",
		ID:   "Xero",
	}

	accountingProviderStore["MYOB"] = &model.AccountingProvider{
		Name: "MYOB",
		ID:   "MYOB",
	}
}
