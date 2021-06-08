package life_phases

import (
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	LifePhaseFormTemplate = "life_phase_form"
)

type LifePhaseForm struct {
	Ft         handlers.FormTemplate
	SchemaID   int
	LifePhase *model.LifePhase
}

func makeLifePhaseFormValues(generationTypeRequest *model.LifePhaseRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = generationTypeRequest.ID
	formValues["SchemaID"] = generationTypeRequest.SchemaID
	formValues["Name"] = generationTypeRequest.Name
	formValues["StartYear"] = generationTypeRequest.StartYear
	formValues["EndYear"] = generationTypeRequest.EndYear
	formValues["Role"] = generationTypeRequest.Role

	return formValues
}

func makeLifePhaseErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/schema_id" {
			formErrorMessages["SchemaID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/name" {
			formErrorMessages["Name"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/start_year" {
			formErrorMessages["StartYear"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/end_year" {
			formErrorMessages["EndYear"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/role" {
			formErrorMessages["Role"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakeLifePhaseForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, submitLabel string,
	lifePhase *model.LifePhase, lifePhaseRequest *model.LifePhaseRequest, errors handlers.ResponseErrors) (*LifePhaseForm, error) {

	formValues := makeLifePhaseFormValues(lifePhaseRequest)
	formErrorMessages := makeLifePhaseErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	lifePhaseTemplate := &LifePhaseForm{
		Ft:        *ft,
		SchemaID:  lifePhaseRequest.SchemaID,
		LifePhase: lifePhase,
	}

	return lifePhaseTemplate, nil
}
