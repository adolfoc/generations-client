package generational_landscape

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func CreateIntangible(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generational-landscapes", "CreateIntangible")

	schemaID, err := getUrlGenerationSchemaID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	landscapeID, err := getUrlGenerationalLandscapeID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	itRequest, err := makeIntangibleRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postIntangible(w, r, itRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(IntangibleCreateErrorsReceivedIndex))
		NewIntangibleRetry(w, r, itRequest, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		generationalLandscape, _ := getGenerationalLandscape(w, r, schemaID, landscapeID)

		handlers.WriteSessionInfoMessage(r, GetLabel(IntangibleCreatedIndex))
		url := fmt.Sprintf("/schemas/%d/generations/%d", schemaID, generationalLandscape.GenerationID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}
}


