package generational_landscape

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdateIntangible(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generational-landscapes", "UpdateIntangible")

	schemaID, err := getUrlGenerationSchemaID(w, r)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	itRequest, err := makeIntangibleRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchIntangible(w, r, itRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(IntangibleUpdateErrorsReceivedIndex))
		EditIntangibleRetry(w, r, itRequest, schemaID, *responseErrors)
		log.NormalReturn()
		return
	}

	if code == http.StatusOK || code == http.StatusCreated {
		generationalLandscape, _ := getGenerationalLandscape(w, r, schemaID, itRequest.LandscapeID)

		handlers.WriteSessionInfoMessage(r, GetLabel(IntangibleUpdatedIndex))
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


