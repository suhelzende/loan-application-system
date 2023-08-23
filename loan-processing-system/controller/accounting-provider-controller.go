package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suhelz/loan-processing-system/model"
	"github.com/suhelz/loan-processing-system/service"
)

type AccountingProviderController struct {
	accountingProviderService service.AccountProviderServiceInterface
}

func NewAccountingProviderController(service service.AccountProviderServiceInterface) AccountingProviderController {
	return AccountingProviderController{
		accountingProviderService: service,
	}
}

func (controller AccountingProviderController) GetAllAccountingProviders(ctx *gin.Context) {
	accountingproviders, err := controller.accountingProviderService.GetAllAccountingProviders()
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, accountingproviders)
}

func (controller AccountingProviderController) RequestBalenceSheet(ctx *gin.Context) {
	balenceSheetRequest := model.BalenceSheetRequest{}
	err := ctx.Bind(&balenceSheetRequest)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	err = validateBalenceSheetRequest(balenceSheetRequest)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}

	balenceSheet, err := controller.accountingProviderService.GetBalenceSheet(balenceSheetRequest.BusinessDetails, balenceSheetRequest.AccountingProvider)
	if err != nil {
		log.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}

	ctx.JSON(http.StatusOK, balenceSheet)
}

func validateBalenceSheetRequest(balenceSheetRequest model.BalenceSheetRequest) error {
	err := validateBusinessDetails(&balenceSheetRequest.BusinessDetails)
	if err != nil {
		return err
	}

	if balenceSheetRequest.AccountingProvider.ID == "" {
		return errors.New("invalid or missing acccountintg provider details")
	}
	return nil
}
