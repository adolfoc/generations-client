package life_segments

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdateLifeSegment(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-life_segments", "UpdateLifeSegment")

	lfRequest, err := makeLifeSegmentRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchLifeSegment(w, r, lfRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(LifeSegmentUpdateErrorsReceivedIndex))
		EditLifeSegmentRetry(w, r, lfRequest, *responseErrors)
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(LifeSegmentUpdatedIndex))
		url := fmt.Sprintf("/persons/%d", lfRequest.PersonID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}
}
