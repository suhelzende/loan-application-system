package main

import (
	"log"

	"github.com/suhelz/loan-processing-system/controller"
	"github.com/suhelz/loan-processing-system/model"
	"github.com/suhelz/loan-processing-system/repository"
	"github.com/suhelz/loan-processing-system/router"
	"github.com/suhelz/loan-processing-system/service"
	"github.com/suhelz/loan-processing-system/storage"
)

func main() {
	// read config
	cfg := model.Config{
		Service: model.Service{
			Port: 8090,
		},
	}

	// initialize inmemory storage for Application
	// in actual system it will have db connection here
	storage.InitStorage()

	// create repository
	applicationRepository := repository.CreateNewApplicationRepository(cfg)
	applicationService := service.CreateNewApplicationService(applicationRepository)
	applicationController := controller.NewApplicationController(applicationService)

	accountingProviderRepository := repository.CreateNewAccountingProviderRepository(cfg)
	accountingProviderService := service.CreateNewAccountingProviderService(accountingProviderRepository)
	accountingProvideController := controller.NewAccountingProviderController(accountingProviderService)

	router := router.NewRouter(cfg)
	// Adding Loan Application controller
	router.AddLoanApplicationController(applicationController)

	// Adding Accounting Provider controller
	router.AddAccountingProviderController(accountingProvideController)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
