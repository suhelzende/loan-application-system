package router

import (
	apiconstants "github.com/suhelz/loan-processing-system/constants/api-constants"
	"github.com/suhelz/loan-processing-system/controller"
)

// Adds AccountingProvider APIs to router
func (r Router) AddAccountingProviderController(accountProvidingController controller.AccountingProviderController) {
	accountingControllerGroup := r.router.Group(apiconstants.EndpointAccountingControllerGroup)

	accountingControllerGroup.GET(apiconstants.EndpointGetAllAccountingProviders, accountProvidingController.GetAllAccountingProviders)

	accountingControllerGroup.POST(apiconstants.EndpointRequestBalanceSheet, accountProvidingController.RequestBalenceSheet)

}
