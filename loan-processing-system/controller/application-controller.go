package controller

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suhelz/loan-processing-system/model"
	"github.com/suhelz/loan-processing-system/service"
)

type ApplicationController struct {
	LoanApplicationService service.LoanApplicationServiceInterface
}

func NewApplicationController(service service.LoanApplicationServiceInterface) ApplicationController {
	return ApplicationController{
		LoanApplicationService: service,
	}
}

func (controller ApplicationController) StartLoanApplication(ctx *gin.Context) {
	loanApplicationRequest := model.LoanApplicationRequest{}

	err := ctx.Bind(&loanApplicationRequest)
	if err != nil {
		// TODO: return custom message instead of nil
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	err = validateLoanApplicationRequest(loanApplicationRequest)
	if err != nil {
		// TODO: return custom message instead of nil
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	loanApplication, err := controller.LoanApplicationService.StartNewApplication(loanApplicationRequest)
	if err != nil {
		// TODO: return custom message instead of nil
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	if loanApplication == nil {
		// TODO: return custom message instead of nil
		log.Println("Unable to start application")
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, loanApplication)
}

func validateLoanApplicationRequest(loanApplicationRequest model.LoanApplicationRequest) error {
	if loanApplicationRequest.BorroweDetails.Name == "" || loanApplicationRequest.BorroweDetails.Email == "" || loanApplicationRequest.Date.Equal(time.Time{}) {
		return errors.New("invalid request missing required fields")
	}
	return nil
}

func (controller ApplicationController) SubmitLoanApplication(ctx *gin.Context) {}

func (controller ApplicationController) GetLoanApplication(ctx *gin.Context) {}
