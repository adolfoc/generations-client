package generational_landscape

import (
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	TangibleFormTemplate = "tangible_form"
)

type TangibleForm struct {
	Ft           handlers.FormTemplate
	SchemaID     int
	GenerationID int
	Tangible     *model.Tangible
}

func makeTangibleFormValues(tangibleRequest *model.TangibleRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = tangibleRequest.ID
	formValues["LandscapeID"] = tangibleRequest.LandscapeID
	formValues["Name"] = tangibleRequest.Name
	formValues["Description"] = tangibleRequest.Description

	return formValues
}

func makeTangibleErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/landscape_id" {
			formErrorMessages["LandscapeID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/name" {
			formErrorMessages["Name"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/description" {
			formErrorMessages["Description"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakeTangibleForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, studyTitle, submitLabel string,
	tangible *model.Tangible, tRequest *model.TangibleRequest, schemaID, generationID int,
	errors handlers.ResponseErrors) (*TangibleForm, error) {

	formValues := makeTangibleFormValues(tRequest)
	formErrorMessages := makeTangibleErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, studyTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	tangibleTemplate := &TangibleForm{
		Ft:           *ft,
		SchemaID:     schemaID,
		GenerationID: generationID,
		Tangible:     tangible,
	}

	return tangibleTemplate, nil
}
