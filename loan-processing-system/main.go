package main

import (
	"log"

	"github.com/suhelz/loan-processing-system/controller"
	"github.com/suhelz/loan-processing-system/repository"
	"github.com/suhelz/loan-processing-system/router"
	"github.com/suhelz/loan-processing-system/service"
	"github.com/suhelz/loan-processing-system/storage"
)

func main() {

	// initialize in memory storage for Application
	// in actual system it will have db connection here
	storage.InitStorage()

	// loan application dependencies
	applicationRepository := repository.CreateNewApplicationRepository()
	applicationService := service.CreateNewApplicationService(applicationRepository)
	applicationController := controller.NewApplicationController(applicationService)

	// accounting provider dependencies
	accountingProviderRepository := repository.CreateNewAccountingProviderRepository()
	accountingProviderService := service.CreateNewAccountingProviderService(accountingProviderRepository)
	accountingProvideController := controller.NewAccountingProviderController(accountingProviderService)

	router := router.NewRouter()
	// Adding Loan Application controller
	router.AddLoanApplicationController(applicationController)

	// Adding Accounting Provider controller
	router.AddAccountingProviderController(accountingProvideController)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
