package controller

import "github.com/suhelz/loan-processing-system/service"

type Controller struct {
	LoanApplicationService service.LoanApplicationServiceInterface
}

func CreateNewController(service service.LoanApplicationServiceInterface) Controller {
	return &Controller{
		LoanApplicationService: service,
	}
}
