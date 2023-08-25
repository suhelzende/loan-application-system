package storage

// Mock DB behavior

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/suhelz/loan-processing-system/model"
)

func CreateNewLoanApplication(application *model.LoanApplication) *model.LoanApplication {
	application.ID = GenerateApplicationID()
	application.LastModified = time.Now()
	loanApplicationsStore[application.ID] = application
	return application
}

func UpdateLoanApplication(application *model.LoanApplication) error {
	if _, ok := loanApplicationsStore[application.ID]; !ok {
		return nil
	}

	application.LastModified = time.Now()
	loanApplicationsStore[application.ID] = application
	return nil
}

func GetApplicationByID(id string) *model.LoanApplication {
	application := loanApplicationsStore[id]
	return application
}

func GenerateApplicationID() string {
	id := uuid.New().String()
	id = strings.ReplaceAll(id, "-", "")
	id = strings.ToUpper(id)
	return id
}
