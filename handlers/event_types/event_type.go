package event_types

import (
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type EventTypeTemplate struct {
	Ct        handlers.CommonTemplate
	EventType *model.EventType
}

func MakeEventTypeTemplate(r *http.Request, pageTitle, studyTitle string, eventType *model.EventType) (*EventTypeTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle, studyTitle)
	if err != nil {
		return nil, err
	}

	evenTypeTemplate := &EventTypeTemplate{
		Ct:        *ct,
		EventType: eventType,
	}

	return evenTypeTemplate, nil
}

func GetEventType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-event-types", "GetEventType")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	eventID, err := getUrlEventTypeID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	person, err := getEventType(w, r, eventID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakeEventTypeTemplate(r, GetLabel(EventTypePageTitleIndex), "", person)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("event_type", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}
