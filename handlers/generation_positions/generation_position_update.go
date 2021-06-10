package generation_positions

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdateGenerationPosition(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generation_positions", "UpdateGenerationPosition")

	schemaID, err := getUrlGenerationSchemaID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	generationPositionRequest, err := makeGenerationPositionRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchGenerationPosition(w, r, generationPositionRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(GenerationPositionUpdateErrorsReceivedIndex))
		EditGenerationPositionRetry(w, r, generationPositionRequest, schemaID, *responseErrors)
		log.NormalReturn()
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(GenerationPositionUpdatedIndex))
		url := fmt.Sprintf("/schemas/%d/generations/%d", schemaID, generationPositionRequest.GenerationID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		log.NormalReturn()
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}

