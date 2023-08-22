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
	cfg := model.Config{}

	// create repository
	// TODO: pass actual config
	loanRepository := repository.CreateNewApplicationRepository(cfg)

	loanService := service.CreateNewApplicationService(loanRepository)

	applicationController := controller.CreateNewController(loanService)

	router := router.CreateApplicationRouter(cfg, applicationController)

	err := router.run()
	if err != nil {
		log.Fatal("Failedd to start the server: ",err.Error())
	}
}
