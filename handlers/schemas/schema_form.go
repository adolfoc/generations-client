package schemas

import (
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	SchemaFormTemplate = "schema_form"
)

type SchemaForm struct {
	Ft     handlers.FormTemplate
	Schema *model.GenerationSchema
	Places []*model.Place
}

func makeSchemaFormValues(sRequest *model.GenerationSchemaRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = sRequest.ID
	formValues["Name"] = sRequest.Name
	formValues["Description"] = sRequest.Description
	formValues["StartYear"] = sRequest.StartYear
	formValues["EndYear"] = sRequest.EndYear
	formValues["MinimumGenerationSpan"] = sRequest.MinimumGenerationSpan
	formValues["MaximumGenerationSpan"] = sRequest.MaximumGenerationSpan
	formValues["PlaceID"] = sRequest.PlaceID

	return formValues
}

func makeSchemaErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/name" {
			formErrorMessages["Name"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/description" {
			formErrorMessages["Description"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/start_year" {
			formErrorMessages["StartYear"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/end_year" {
			formErrorMessages["EndYear"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/minimum_generation_span" {
			formErrorMessages["MinimumGenerationSpan"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/maximum_generation_span" {
			formErrorMessages["MaximumGenerationSpan"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/place_id" {
			formErrorMessages["PlaceID"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakeSchemaForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, studyTitle, submitLabel string, schema *model.GenerationSchema,
	sRequest *model.GenerationSchemaRequest, places []*model.Place, errors handlers.ResponseErrors) (*SchemaForm, error) {

	formValues := makeSchemaFormValues(sRequest)
	formErrorMessages := makeSchemaErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, studyTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	sTemplate := &SchemaForm{
		Ft:     *ft,
		Schema: schema,
		Places: places,
	}

	return sTemplate, nil
}
