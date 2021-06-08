package generational_landscape

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func CreateGenerationalLandscape(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generational-landscapes", "CreateGenerationalLandscape")

	schemaID, err := getUrlGenerationSchemaID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	glRequest, err := makeGenerationalLandscapeRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postGenerationalLandscape(w, r, glRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(GenerationalLandscapeCreateErrorsReceivedIndex))
		NewGenerationalLandscapeRetry(w, r, glRequest, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		var generation *model.Generation
		_ = json.Unmarshal(body, &generation)

		handlers.WriteSessionInfoMessage(r, GetLabel(GenerationalLandscapeCreatedIndex))
		url := fmt.Sprintf("/schemas/%d/generations/%d", schemaID, glRequest.GenerationID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}
}
