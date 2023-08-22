package model

import "time"

type LoanApplicationRequest struct {
	BorroweDetails Borrowe   `json:"borrower"`
	Date           time.Time `json:"date"`
}
