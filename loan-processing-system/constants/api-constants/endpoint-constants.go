package apiconstants

const (
	// Application API endpoint
	EndpointLoanApplicationControllerGroup = "/application"

	EndpointStartLoanApplication = "/start"

	EndpointSubmitLoanApplication = "/submit"

	EndpointGetLoanApplicationByID = "/id/:applicationID"

	PathParamApplicationID = "applicationID"

	// Accounting API endpoint
	EndpointAccountingControllerGroup = "/accounting"

	EndpointGetAllAccountingProviders = "/providers"

	EndpointRequestBalanceSheet = "/balencesheet/request"

	// Default port to start the server
	DefaultPort = 8090
)
