package generations

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func CreateGeneration(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-authors", "CreateGeneration")

	generationRequest, err := makeGenerationRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postGeneration(w, r, generationRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(GenerationCreateErrorsReceivedIndex))
		NewGenerationRetry(w, r, generationRequest, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		var generation *model.Generation
		_ = json.Unmarshal(body, &generation)

		handlers.WriteSessionInfoMessage(r, GetLabel(GenerationCreatedIndex))
		url := fmt.Sprintf("/schemas/%d/generations/%d", generation.SchemaID, generation.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}
}
