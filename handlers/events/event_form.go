package events

import (
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	EventFormTemplate = "event_form"
)

type EventForm struct {
	Ft         handlers.FormTemplate
	Event      *model.Event
	EventTypes []*model.EventType
}

func makeEventFormValues(eRequest *model.EventRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = fmt.Sprintf("%d", eRequest.ID)
	formValues["TypeID"] = eRequest.TypeID
	formValues["Name"] = eRequest.Name
	formValues["Start"] = eRequest.Start
	formValues["End"] = eRequest.End
	formValues["Summary"] = eRequest.Summary
	formValues["Description"] = eRequest.Description

	return formValues
}

func makeEventErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/event_type_id" {
			formErrorMessages["TypeID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/name" {
			formErrorMessages["Name"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/start" {
			formErrorMessages["Start"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/end" {
			formErrorMessages["End"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/summary" {
			formErrorMessages["Summary"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/description" {
			formErrorMessages["Description"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakeEventForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, submitLabel string, event *model.Event,
	eRequest *model.EventRequest, eventTypes []*model.EventType, errors handlers.ResponseErrors) (*EventForm, error) {

	formValues := makeEventFormValues(eRequest)
	formErrorMessages := makeEventErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	eTemplate := &EventForm{
		Ft:         *ft,
		Event:      event,
		EventTypes: eventTypes,
	}

	return eTemplate, nil
}
