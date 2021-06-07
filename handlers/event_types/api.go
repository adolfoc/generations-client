package event_types

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourceEventTypes = "event-types"
)

func getEventTypesURL(page int) string {
	return fmt.Sprintf("%s%s?page[number]=%d", handlers.GetAPIHostURL(), ResourceEventTypes, page)
}

func getSimpleEventTypesURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourceEventTypes)
}

func getEventTypeURL(eventTypeID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceEventTypes, eventTypeID)
}

func getEventTypes(w http.ResponseWriter, r *http.Request, page int) (*model.EventTypes, error) {
	url := getEventTypesURL(page)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var eventTypes *model.EventTypes
	err = json.Unmarshal(body, &eventTypes)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return eventTypes, nil
}

func getEventType(w http.ResponseWriter, r *http.Request, eventTypeID int) (*model.EventType, error) {
	url := getEventTypeURL(eventTypeID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var eventType *model.EventType
	err = json.Unmarshal(body, &eventType)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return eventType, nil
}

func getUrlEventTypeID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("event_type_id", w, r)
}

func buildEventTypeRequest(eventType *model.EventType) *model.EventTypeRequest {
	mr := &model.EventTypeRequest{
		ID:          eventType.ID,
		IsNatural:   eventType.IsNatural,
		Name:        eventType.Name,
		Description: eventType.Description,
	}

	return mr
}

func makeEventTypeRequest(r *http.Request) (*model.EventTypeRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	isNatural := handlers.GetBoolFormValue(r, "inputIsNatural")
	name := handlers.GetStringFormValue(r, "inputName")
	description := handlers.GetStringFormValue(r, "inputDescription")

	ptr := &model.EventTypeRequest{
		ID:          normalizedID,
		IsNatural:   isNatural,
		Name:        name,
		Description: description,
	}

	return ptr, nil
}

func postEventType(w http.ResponseWriter, r *http.Request, eventTypeRequest *model.EventTypeRequest) (int, []byte, error) {
	url := getSimpleEventTypesURL()

	payload, err := json.Marshal(eventTypeRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchEventType(w http.ResponseWriter, r *http.Request, eventTypeRequest *model.EventTypeRequest) (int, []byte, error) {
	url := getEventTypeURL(eventTypeRequest.ID)

	payload, err := json.Marshal(eventTypeRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}

	return code, body, nil
}
