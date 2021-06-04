package events

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditEvent(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-events", "EditEvent")

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

	event, err := getEvent(w, r, eventID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	eventTypes, err := getEventTypes(w, r)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	eRequest := buildEventRequest(event)

	url := fmt.Sprintf("/events/%d/update", eventID)
	eventForm, err := MakeEventForm(w, r, url, GetLabel(EventEditPageTitleIndex),
		GetLabel(EventEditSubmitLabelIndex), event, eRequest, eventTypes, handlers.ResponseErrors{})
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

func EditEventRetry(w http.ResponseWriter, r *http.Request, eRequest *model.EventRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-events", "EditEventRetry")

	eventTypes, err := getEventTypes(w, r)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}
	eventType := matchEventType(eRequest.TypeID, eventTypes)

	event := &model.Event{
		ID:          eRequest.ID,
		EventType:   eventType,
		Name:        eRequest.Name,
		Start:       eRequest.Start,
		End:         eRequest.End,
		Summary:     eRequest.Summary,
		Description: eRequest.Description,
	}

	url := fmt.Sprintf("/events/%d/update", eRequest.ID)
	eventForm, err := MakeEventForm(w, r, url, GetLabel(EventEditPageTitleIndex),
		GetLabel(EventEditSubmitLabelIndex), event, eRequest, eventTypes, errors)
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

func getEventTypes(w http.ResponseWriter, r *http.Request) ([]*model.EventType, error) {
	url := fmt.Sprintf("%sevent-types", handlers.GetAPIHostURL())
	code, body, err := handlers.GetResource(w, r, url)

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	if err != nil {
		return nil, err
	}

	var eventTypes []*model.EventType
	err = json.Unmarshal(body, &eventTypes)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return eventTypes, nil
}

func matchEventType(typeID int, eventTypes []*model.EventType) *model.EventType {
	for _, et := range eventTypes {
		if et.ID == typeID {
			return et
		}
	}

	return nil
}

