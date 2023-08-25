package model

import "time"

type LoanApplicationRequest struct {
	BorroweDetails Borrower  `json:"borrower"`
	Date           time.Time `json:"date"`
}

type BalenceSheetRequest struct {
	BusinessDetails    BusinessDetails    `json:"businessDetails"`
	AccountingProvider AccountingProvider `json:"accountingProvider"`
}

type SubmitLoanApplicationRequest struct {
	LoanDetails  LoanApplication `json:"loanDetails"`
	BalanceSheet BalenceSheet    `json:"balanceSheet"`
}

type ProfitLossSummary struct {
	Year         int
	ProfitOrLoss int
}
