package apiconstants

const (
	// Application
	EndpointLoanApplicationControllerGroup = "/application"

	EndpointStartLoanApplication = "/start"

	EndpointSubmitLoanApplication = "/submit"

	EndpointGetLoanApplicationByID = "/id/:applicationID"

	PathParamApplicationID = "applicationID"

	// Accounting
	EndpointAccountingControllerGroup = "/accounting"

	EndpointGetAllAccountingProviders = "/providers"

	EndpointRequestBalenceSheet = "/balencesheet/request"

	DefaultPort = 8090
)
