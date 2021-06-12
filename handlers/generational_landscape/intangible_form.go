package generational_landscape

import (
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	IntangibleFormTemplate = "intangible_form"
)

type IntangibleForm struct {
	Ft           handlers.FormTemplate
	SchemaID     int
	GenerationID int
	Intangible   *model.Intangible
}

func makeIntangibleFormValues(tangibleRequest *model.IntangibleRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = tangibleRequest.ID
	formValues["LandscapeID"] = tangibleRequest.LandscapeID
	formValues["Name"] = tangibleRequest.Name
	formValues["Description"] = tangibleRequest.Description

	return formValues
}

func makeIntangibleErrorMessages(errors handlers.ResponseErrors) map[string]string {
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

func MakeIntangibleForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, studyTitle, submitLabel string,
	tangible *model.Intangible, tRequest *model.IntangibleRequest, schemaID, generationID int,
	errors handlers.ResponseErrors) (*IntangibleForm, error) {

	formValues := makeIntangibleFormValues(tRequest)
	formErrorMessages := makeIntangibleErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, studyTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	intangibleTemplate := &IntangibleForm{
		Ft:           *ft,
		SchemaID:     schemaID,
		GenerationID: generationID,
		Intangible:   tangible,
	}

	return intangibleTemplate, nil
}
