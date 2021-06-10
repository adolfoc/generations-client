package generation_positions

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func CreateGenerationPosition(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generation_positions", "CreateGenerationPosition")

	schemaID, err := getUrlGenerationSchemaID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	generationID, err := getUrlGenerationID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	generationPositionRequest, err := makeGenerationPositionRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postGenerationPosition(w, r, generationPositionRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(GenerationPositionCreateErrorsReceivedIndex))
		NewGenerationPositionRetry(w, r, generationPositionRequest, schemaID, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(GenerationPositionCreatedIndex))
		url := fmt.Sprintf("/schemas/%d/generations/%d", schemaID, generationID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}
}
