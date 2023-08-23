package storage

import (
	"github.com/suhelz/loan-processing-system/model"
)

// This package serves as data store for loan-application systems

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
