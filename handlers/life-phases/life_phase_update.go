package life_phases

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdateLifePhase(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-life_phases", "UpdateLifePhase")

	lifePhaseRequest, err := makeLifePhaseRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchLifePhase(w, r, lifePhaseRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(LifePhaseUpdateErrorsReceivedIndex))
		EditLifePhaseRetry(w, r, lifePhaseRequest, *responseErrors)
		log.NormalReturn()
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(LifePhaseUpdatedIndex))
		url := fmt.Sprintf("/schemas/%d", lifePhaseRequest.SchemaID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		log.NormalReturn()
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}

