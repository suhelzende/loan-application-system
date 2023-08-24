package controller

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	apiconstants "github.com/suhelz/loan-processing-system/constants/api-constants"
	"github.com/suhelz/loan-processing-system/model"
	"github.com/suhelz/loan-processing-system/service"
)

type ApplicationController struct {
	loanApplicationService service.LoanApplicationServiceInterface
}

func NewApplicationController(service service.LoanApplicationServiceInterface) ApplicationController {
	return ApplicationController{
		loanApplicationService: service,
	}
}

func (controller ApplicationController) StartLoanApplication(ctx *gin.Context) {
	loanApplicationRequest := model.LoanApplicationRequest{}

	err := ctx.Bind(&loanApplicationRequest)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	err = validateLoanApplicationRequest(loanApplicationRequest)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	loanApplication, err := controller.loanApplicationService.StartNewApplication(loanApplicationRequest)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	if loanApplication == nil {
		log.Println("Unable to start application")
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, loanApplication)
}

func (controller ApplicationController) SubmitLoanApplication(ctx *gin.Context) {
	loanApplicationRequest := model.SubmitLoanApplicationRequest{}

	err := ctx.Bind(&loanApplicationRequest)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	err = validateSubmitLoanApplicationRequest(loanApplicationRequest)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	details, err := controller.loanApplicationService.GetApplicationByID(loanApplicationRequest.LoanDetails.ID)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	if details == nil {
		log.Println("No application found for given id")
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	applicationResult, err := controller.loanApplicationService.SubmitApplication(loanApplicationRequest)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, applicationResult)
}

func (controller ApplicationController) GetLoanApplication(ctx *gin.Context) {
	applicationID := ctx.Param(apiconstants.PathParamApplicationID)
	if applicationID == "" {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}

	application, err := controller.loanApplicationService.GetApplicationByID(applicationID)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	if application == nil {
		log.Println("Application not found for given application ID")
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(http.StatusOK, application)
}

func validateLoanApplicationRequest(loanApplicationRequest model.LoanApplicationRequest) error {
	if loanApplicationRequest.BorroweDetails.Name == "" || loanApplicationRequest.BorroweDetails.Email == "" || loanApplicationRequest.Date.Equal(time.Time{}) {
		return errors.New("invalid request missing required fields")
	}
	return nil
}

func validateSubmitLoanApplicationRequest(loanApplicationRequest model.SubmitLoanApplicationRequest) error {
	if err := validateBusinessDetails(loanApplicationRequest.LoanDetails.BusinessDetails); err != nil {
		return err
	}

	if len(loanApplicationRequest.BalenceSheet.Sheet) == 0 {
		return errors.New("missing balence sheet")
	}

	if loanApplicationRequest.LoanDetails.Borrower.Email == "" || loanApplicationRequest.LoanDetails.Borrower.Name == "" {
		return errors.New("missing borrower details")
	}

	return nil
}

func validateBusinessDetails(businessDetails *model.BusinessDetails) error {
	if businessDetails == nil {
		return errors.New("missing required business details")
	}
	if businessDetails.Name == "" ||
		businessDetails.EstablishedYear == 0 ||
		businessDetails.EstablishedYear > time.Now().Year() ||
		businessDetails.RegistrationID == "" {
		return errors.New("invalid request missing required business details")
	}
	return nil
}
