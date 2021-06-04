package events

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-events", "UpdateEvent")

	eventRequest, err := makeEventRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchEvent(w, r, eventRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(EventUpdateErrorsReceivedIndex))
		EditEventRetry(w, r, eventRequest, *responseErrors)
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(EventUpdatedIndex))
		url := fmt.Sprintf("/events/%d", eventRequest.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}


