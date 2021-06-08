package moment_types

import (
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	MomentTypeFormTemplate = "moment_type_form"
)

type MomentTypeForm struct {
	Ft         handlers.FormTemplate
	SchemaID   int
	MomentType *model.MomentType
}

func makeMomentTypeFormValues(generationTypeRequest *model.MomentTypeRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = generationTypeRequest.ID
	formValues["SchemaID"] = generationTypeRequest.SchemaID
	formValues["Name"] = generationTypeRequest.Name
	formValues["Description"] = generationTypeRequest.Description

	return formValues
}

func makeMomentTypeErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/schema_id" {
			formErrorMessages["SchemaID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/name" {
			formErrorMessages["Name"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/description" {
			formErrorMessages["Description"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakeMomentTypeForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, submitLabel string,
	momentType *model.MomentType, momentTypeRequest *model.MomentTypeRequest,
	errors handlers.ResponseErrors) (*MomentTypeForm, error) {

	formValues := makeMomentTypeFormValues(momentTypeRequest)
	formErrorMessages := makeMomentTypeErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	momentTypeTemplate := &MomentTypeForm{
		Ft:         *ft,
		SchemaID:   momentTypeRequest.SchemaID,
		MomentType: momentType,
	}

	return momentTypeTemplate, nil
}

