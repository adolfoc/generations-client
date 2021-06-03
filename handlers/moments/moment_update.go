package moments

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdateMoment(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-moments", "UpdateMoment")

	momentRequest, err := makeMomentRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchMoment(w, r, momentRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(MomentUpdateErrorsReceivedIndex))
		EditMomentRetry(w, r, momentRequest, *responseErrors)
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(MomentUpdatedIndex))
		url := fmt.Sprintf("/schemas/%d/moments/%d", momentRequest.SchemaID, momentRequest.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}
}

