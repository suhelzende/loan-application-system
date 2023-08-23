package main

import (
	"log"

	"github.com/suhelz/loan-processing-system/controller"
	"github.com/suhelz/loan-processing-system/model"
	"github.com/suhelz/loan-processing-system/repository"
	"github.com/suhelz/loan-processing-system/router"
	"github.com/suhelz/loan-processing-system/service"
)

func main() {
	// read config
	cfg := model.Config{
		Service: model.Service{
			Port: 8090,
		},
	}

	// create repository
	// TODO: pass actual config
	loanRepository := repository.CreateNewApplicationRepository(cfg)

	loanService := service.CreateNewApplicationService(loanRepository)

	applicationController := controller.NewApplicationController(loanService)

	router := router.NewRouter(cfg)

	// Adding Loan Application controller
	router.AddLoanApplicationController(applicationController)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
