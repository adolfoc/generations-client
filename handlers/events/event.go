package events

import (
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type EventTemplate struct {
	Ct    handlers.CommonTemplate
	Event *model.Event
}

func MakeEventTemplate(r *http.Request, pageTitle string, event *model.Event) (*EventTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle)
	if err != nil {
		return nil, err
	}

	eventTemplate := &EventTemplate{
		Ct:    *ct,
		Event: event,
	}

	return eventTemplate, nil
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-events", "GetEvent")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	eventID, err := getUrlEventID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	person, err := getEvent(w, r, eventID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakeEventTemplate(r, GetLabel(EventPageTitleIndex), person)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("event", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}
