package events

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func NewEvent(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-events", "NewEvent")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	eventTypes, err := getEventTypes(w, r)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	generationRequest := &model.EventRequest{}

	url := fmt.Sprintf("/events/create")
	eventForm, err := MakeEventForm(w, r, url, GetLabel(EventNewPageTitleIndex),
		GetLabel(EventNewSubmitLabelIndex), &model.Event{}, generationRequest, eventTypes, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(EventFormTemplate, eventForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func NewEventRetry(w http.ResponseWriter, r *http.Request, eventRequest *model.EventRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-events", "NewEventRetry")

	eventTypes, err := getEventTypes(w, r)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}
	eventType := matchEventType(eventRequest.TypeID, eventTypes)

	event := &model.Event{
		ID:          eventRequest.ID,
		EventType:   eventType,
		Name:        eventRequest.Name,
		Start:       eventRequest.Start,
		End:         eventRequest.End,
		Summary:     eventRequest.Summary,
		Description: eventRequest.Description,
	}

	url := fmt.Sprintf("/events/create")
	eventForm, err := MakeEventForm(w, r, url, GetLabel(EventNewPageTitleIndex),
		GetLabel(EventNewSubmitLabelIndex), event, eventRequest, eventTypes, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(EventFormTemplate, eventForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}
