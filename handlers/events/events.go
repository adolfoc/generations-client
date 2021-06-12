package events

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type EventsTemplate struct {
	Ct         handlers.CommonTemplate
	Events     *model.Events
	Pagination *handlers.Pagination
}

func getPaginationBaseURL() string {
	stem := fmt.Sprintf("/events/index")
	return stem + "?page=%d"
}

func MakeEventsTemplate(r *http.Request, pageTitle, studyTitle string, page int, events *model.Events) (*EventsTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle, studyTitle)
	if err != nil {
		return nil, err
	}

	pagination := handlers.MakePagination(events.RecordCount, getPaginationBaseURL(), page)

	eventsTemplate := &EventsTemplate{
		Ct:         *ct,
		Events:     events,
		Pagination: pagination,
	}

	return eventsTemplate, nil
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-events", "GetEvents")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	page := handlers.GetURLPageParameter(r)
	events, err := getEvents(w, r, page)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakeEventsTemplate(r, GetLabel(EventIndexPageTitleIndex), "", page, events)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("events", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}
