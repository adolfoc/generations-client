package life_phases

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func CreateLifePhase(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-life_phases", "CreateLifePhase")

	momentTypeRequest, err := makeLifePhaseRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postLifePhase(w, r, momentTypeRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(LifePhaseCreateErrorsReceivedIndex))
		NewLifePhaseRetry(w, r, momentTypeRequest, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(LifePhaseCreatedIndex))
		url := fmt.Sprintf("/schemas/%d", momentTypeRequest.SchemaID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}
}


