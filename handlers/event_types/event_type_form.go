package event_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	EventTypeFormTemplate = "event_type_form"
)

type EventTypeForm struct {
	Ft        handlers.FormTemplate
	EventType *model.EventType
}

func makeEventTypeFormValues(etRequest *model.EventTypeRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = fmt.Sprintf("%d", etRequest.ID)
	formValues["IsNatural"] = etRequest.IsNatural
	formValues["Name"] = etRequest.Name
	formValues["Description"] = etRequest.Description

	return formValues
}

func makeEventTypeErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/name" {
			formErrorMessages["Name"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/is_natural" {
			formErrorMessages["IsNatural"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/description" {
			formErrorMessages["Description"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakeEventTypeForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, studyTitle, submitLabel string, eventType *model.EventType,
	etRequest *model.EventTypeRequest, errors handlers.ResponseErrors) (*EventTypeForm, error) {

	formValues := makeEventTypeFormValues(etRequest)
	formErrorMessages := makeEventTypeErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, studyTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	ptTemplate := &EventTypeForm{
		Ft:        *ft,
		EventType: eventType,
	}

	return ptTemplate, nil
}

