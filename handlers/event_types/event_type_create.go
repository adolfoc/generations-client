package event_types

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func CreateEventType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-event_types", "CreateEventType")

	eventTypeRequest, err := makeEventTypeRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postEventType(w, r, eventTypeRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(EventTypeCreateErrorsReceivedIndex))
		NewEventTypeRetry(w, r, eventTypeRequest, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		var eventType *model.EventType
		_ = json.Unmarshal(body, &eventType)

		handlers.WriteSessionInfoMessage(r, GetLabel(EventTypeCreatedIndex))
		url := fmt.Sprintf("/event-types/%d", eventType.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}

	log.NormalReturn()
}

