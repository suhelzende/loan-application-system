package router

import (
	apiconstants "github.com/suhelz/loan-processing-system/constants/api-constants"
	"github.com/suhelz/loan-processing-system/controller"
)

// Adds Application APIs to router
func (r Router) AddLoanApplicationController(applicationController controller.ApplicationController) {
	applicationControllerGroup := r.router.Group(apiconstants.EndpointLoanApplicationControllerGroup)

	applicationControllerGroup.POST(apiconstants.EndpointStartLoanApplication, applicationController.StartLoanApplication)

	applicationControllerGroup.POST(apiconstants.EndpointSubmitLoanApplication, applicationController.SubmitLoanApplication)

	applicationControllerGroup.GET(apiconstants.EndpointGetLoanApplicationByID, applicationController.GetLoanApplication)
}
