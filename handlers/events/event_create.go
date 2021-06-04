package events

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-events", "CreateEvent")

	eventRequest, err := makeEventRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postEvent(w, r, eventRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(EventCreateErrorsReceivedIndex))
		NewEventRetry(w, r, eventRequest, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		var event *model.Event
		_ = json.Unmarshal(body, &event)

		handlers.WriteSessionInfoMessage(r, GetLabel(EventCreatedIndex))
		url := fmt.Sprintf("/events/%d", event.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}

	log.NormalReturn()
}

