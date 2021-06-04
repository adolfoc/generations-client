package events

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourceEvents = "events"
)

func getEventsURL(page int) string {
	return fmt.Sprintf("%s%s?page[number]=%d", handlers.GetAPIHostURL(), ResourceEvents, page)
}

func getSimpleEventsURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourceEvents)
}

func getEventURL(eventID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceEvents, eventID)
}

func getEvents(w http.ResponseWriter, r *http.Request, page int) (*model.Events, error) {
	url := getEventsURL(page)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var events *model.Events
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return events, nil
}

func getEvent(w http.ResponseWriter, r *http.Request, eventID int) (*model.Event, error) {
	url := getEventURL(eventID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var event *model.Event
	err = json.Unmarshal(body, &event)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return event, nil
}

func getUrlEventID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("event_id", w, r)
}

func buildEventRequest(event *model.Event) *model.EventRequest {
	er := &model.EventRequest{
		ID:          event.ID,
		TypeID:      event.EventType.ID,
		Name:        event.Name,
		Start:       event.Start,
		End:         event.End,
		Summary:     event.Summary,
		Description: event.Description,
	}

	return er
}

func makeEventRequest(r *http.Request) (*model.EventRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	normalizedTypeID := handlers.GetIntFormValue(r, "inputTypeID")
	name := handlers.GetStringFormValue(r, "inputName")
	start := handlers.GetStringFormValue(r, "inputStart")
	end := handlers.GetStringFormValue(r, "inputEnd")
	summary := handlers.GetStringFormValue(r, "inputSummary")
	description := handlers.GetStringFormValue(r, "inputDescription")

	er := &model.EventRequest{
		ID:          normalizedID,
		TypeID:      normalizedTypeID,
		Name:        name,
		Start:       start,
		End:         end,
		Summary:     summary,
		Description: description,
	}

	return er, nil
}

func postEvent(w http.ResponseWriter, r *http.Request, eventRequest *model.EventRequest) (int, []byte, error) {
	url := getSimpleEventsURL()

	payload, err := json.Marshal(eventRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchEvent(w http.ResponseWriter, r *http.Request, eRequest *model.EventRequest) (int, []byte, error) {
	url := getEventURL(eRequest.ID)

	payload, err := json.Marshal(eRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}


