package dicisionengine

import "github.com/suhelz/loan-processing-system/model"

// As per the given problem statement it is not clear, if dicision enginer is third party API or third party SDK
// Also problem statment does not give much details about the dicision engine system,
// So, this acts as dicision Engine mock
// we can integrate dicision engine APIs here
func SubmitApplication(request model.DicisionEngineRequest) (string, error) {
	return "ACCEPTED", nil
}
