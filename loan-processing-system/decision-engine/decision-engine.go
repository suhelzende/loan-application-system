package decisionengine

import "github.com/suhelz/loan-processing-system/model"

// As per the given problem statement, decision engine is third party system
// So, this acts as decision Engine mock
// we can integrate decision engine APIs here
func SubmitApplication(request model.DecisionEngineRequest) (string, error) {
	if request.PreAssessmentValue <= 20 {
		return "REJECTED", nil
	}
	return "ACCEPTED", nil
}
