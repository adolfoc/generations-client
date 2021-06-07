package event_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdateEventType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-event-types", "UpdateEventType")

	eventTypeRequest, err := makeEventTypeRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchEventType(w, r, eventTypeRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(EventTypeUpdateErrorsReceivedIndex))
		EditEventTypeRetry(w, r, eventTypeRequest, *responseErrors)
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(EventTypeUpdatedIndex))
		url := fmt.Sprintf("/event-types/%d", eventTypeRequest.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}
