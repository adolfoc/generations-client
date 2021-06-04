package persons

import (
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	AddSegmentsFormTemplate = "add_segments_form"
)

type AddSegmentsForm struct {
	Ft      handlers.FormTemplate
	Person  *model.Person
	Schemas []*model.GenerationSchema
}

func makeAddSegmentsFormValues(schemaID int) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["SchemaID"] = schemaID

	return formValues
}

func makeAddSegmentsErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	return formErrorMessages
}

func MakeAddSegmentsForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, submitLabel string, person *model.Person,
	schemas []*model.GenerationSchema, schemaID int, errors handlers.ResponseErrors) (*AddSegmentsForm, error) {

	formValues := makeAddSegmentsFormValues(schemaID)
	formErrorMessages := makeAddSegmentsErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	lfTemplate := &AddSegmentsForm{
		Ft:      *ft,
		Person:  person,
		Schemas: schemas,
	}

	return lfTemplate, nil
}

