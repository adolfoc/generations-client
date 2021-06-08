package moment_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func CreateMomentType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-moment_types", "CreateMomentType")

	momentTypeRequest, err := makeMomentTypeRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postMomentType(w, r, momentTypeRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(MomentTypeCreateErrorsReceivedIndex))
		NewMomentTypeRetry(w, r, momentTypeRequest, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(MomentTypeCreatedIndex))
		url := fmt.Sprintf("/schemas/%d", momentTypeRequest.SchemaID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}
}

