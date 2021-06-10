package event_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func NewEventType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-event-types", "NewEventType")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	eventTypeRequest := &model.EventTypeRequest{}

	url := fmt.Sprintf("/event-types/create")
	eventTypeForm, err := MakeEventTypeForm(w, r, url, GetLabel(EventTypeNewPageTitleIndex),
		GetLabel(EventTypeNewSubmitLabelIndex), &model.EventType{}, eventTypeRequest, handlers.ResponseErrors{})
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

func NewEventTypeRetry(w http.ResponseWriter, r *http.Request, eventTypeRequest *model.EventTypeRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-event-types", "NewEventTypeRetry")

	eventType := &model.EventType{
		ID:          0,
		IsNatural:   eventTypeRequest.IsNatural,
		Name:        eventTypeRequest.Name,
		Description: eventTypeRequest.Description,
	}

	url := fmt.Sprintf("/event-types/create")
	eventTypeForm, err := MakeEventTypeForm(w, r, url, GetLabel(EventTypeNewPageTitleIndex),
		GetLabel(EventTypeNewSubmitLabelIndex), eventType, eventTypeRequest, errors)
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
	return
}

