package event_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type EventTypesTemplate struct {
	Ct         handlers.CommonTemplate
	EventTypes *model.EventTypes
	Pagination *handlers.Pagination
}

func getPaginationBaseURL() string {
	stem := fmt.Sprintf("/event-types/index")
	return stem + "?page=%d"
}

func MakeEventTypesTemplate(r *http.Request, pageTitle string, page int, eventTypes *model.EventTypes) (*EventTypesTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle)
	if err != nil {
		return nil, err
	}

	pagination := handlers.MakePagination(eventTypes.RecordCount, getPaginationBaseURL(), page)

	eventsTemplate := &EventTypesTemplate{
		Ct:         *ct,
		EventTypes: eventTypes,
		Pagination: pagination,
	}

	return eventsTemplate, nil
}

func GetEventTypes(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-event-types", "GetEventTypes")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	page := handlers.GetURLPageParameter(r)
	events, err := getEventTypes(w, r, page)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakeEventTypesTemplate(r, GetLabel(EventTypeIndexPageTitleIndex), page, events)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("event_types", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}
