package generational_landscape

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdateTangible(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generational-landscapes", "UpdateTangible")

	schemaID, err := getUrlGenerationSchemaID(w, r)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	glRequest, err := makeTangibleRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchTangible(w, r, glRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(TangibleUpdateErrorsReceivedIndex))
		EditTangibleRetry(w, r, glRequest, schemaID, *responseErrors)
		log.NormalReturn()
		return
	}

	if code == http.StatusOK || code == http.StatusCreated {
		generationalLandscape, _ := getGenerationalLandscape(w, r, schemaID, glRequest.LandscapeID)

		handlers.WriteSessionInfoMessage(r, GetLabel(TangibleUpdatedIndex))
		url := fmt.Sprintf("/schemas/%d/generations/%d", schemaID, generationalLandscape.GenerationID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		log.NormalReturn()
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}
