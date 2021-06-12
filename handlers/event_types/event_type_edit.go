package event_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditEventType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-event-types", "EditEventType")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	eventTypeID, err := getUrlEventTypeID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	eventType, err := getEventType(w, r, eventTypeID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	eventTypeRequest := buildEventTypeRequest(eventType)

	url := fmt.Sprintf("/event-types/%d/update", eventTypeID)
	eventTypeForm, err := MakeEventTypeForm(w, r, url, GetLabel(EventTypeEditPageTitleIndex), "",
		GetLabel(EventTypeEditSubmitLabelIndex), eventType, eventTypeRequest, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(EventTypeFormTemplate, eventTypeForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func EditEventTypeRetry(w http.ResponseWriter, r *http.Request, eventTypeRequest *model.EventTypeRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-event-types", "EditEventTypeRetry")

	eventType := &model.EventType{
		ID:            eventTypeRequest.ID,
		IsNatural:     eventTypeRequest.IsNatural,
		Name:          eventTypeRequest.Name,
		Description:   eventTypeRequest.Description,
	}

	url := fmt.Sprintf("/event-types/%d/update", eventTypeRequest.ID)
	placeForm, err := MakeEventTypeForm(w, r, url, GetLabel(EventTypeEditPageTitleIndex), "",
		GetLabel(EventTypeEditSubmitLabelIndex), eventType, eventTypeRequest, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(EventTypeFormTemplate, placeForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}

