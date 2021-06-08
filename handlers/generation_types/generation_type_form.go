package generation_types

import (
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	GenerationTypeFormTemplate = "generation_type_form"
)

type GenerationTypeForm struct {
	Ft                  handlers.FormTemplate
	SchemaID            int
	GenerationType      *model.GenerationType
}

func makeGenerationTypeFormValues(generationTypeRequest *model.GenerationTypeRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = generationTypeRequest.ID
	formValues["SchemaID"] = generationTypeRequest.SchemaID
	formValues["Archetype"] = generationTypeRequest.Archetype
	formValues["Description"] = generationTypeRequest.Description

	return formValues
}

func makeGenerationTypeErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/schema_id" {
			formErrorMessages["SchemaID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/archetype" {
			formErrorMessages["Archetype"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/description" {
			formErrorMessages["Description"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakeGenerationTypeForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, submitLabel string,
	generationType *model.GenerationType, generationTypeRequest *model.GenerationTypeRequest,
	errors handlers.ResponseErrors) (*GenerationTypeForm, error) {

	formValues := makeGenerationTypeFormValues(generationTypeRequest)
	formErrorMessages := makeGenerationTypeErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	generationTypeTemplate := &GenerationTypeForm{
		Ft:             *ft,
		SchemaID:       generationTypeRequest.SchemaID,
		GenerationType: generationType,
	}

	return generationTypeTemplate, nil
}
