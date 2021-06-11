package life_phases

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func DeleteLifePhase(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-life_phases", "DeleteLifePhase")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	schemaID, err := getUrlGenerationSchemaID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	lifePhaseID, err := getUrlLifePhaseID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := deleteLifePhase(w, r, schemaID, lifePhaseID)
	if err != nil {
		log.FailedReturn()
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(LifePhaseDeletedIndex))
	} else {
		responseErrors, _ := handlers.OnDeleteError(r, body)
		handlers.WriteSessionErrorMessage(r, handlers.ExtractFirstError(responseErrors))
	}

	url := fmt.Sprintf("/schemas/%d", schemaID)
	http.Redirect(w, r, url, http.StatusMovedPermanently)

	log.NormalReturn()
}

