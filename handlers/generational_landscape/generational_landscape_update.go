package generational_landscape

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdateGenerationalLandscape(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generational-landscapes", "UpdateGeneration")

	schemaID, err := getUrlGenerationSchemaID(w, r)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	glRequest, err := makeGenerationalLandscapeRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchGenerationalLandscape(w, r, glRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(GenerationalLandscapeUpdateErrorsReceivedIndex))
		EditGenerationalLandscapeRetry(w, r, glRequest, schemaID, *responseErrors)
		log.NormalReturn()
		return
	}

	if code == http.StatusOK || code == http.StatusCreated {
		handlers.WriteSessionInfoMessage(r, GetLabel(GenerationalLandscapeUpdatedIndex))
		url := fmt.Sprintf("/schemas/%d/generations/%d", schemaID, glRequest.GenerationID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		log.NormalReturn()
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}
