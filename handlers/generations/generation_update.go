package generations

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdateGeneration(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generations", "UpdateGeneration")

	generationRequest, err := makeGenerationRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchGeneration(w, r, generationRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(GenerationUpdateErrorsReceivedIndex))
		EditGenerationRetry(w, r, generationRequest, *responseErrors)
		log.NormalReturn()
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(GenerationUpdatedIndex))
		url := fmt.Sprintf("/schemas/%d/generations/%d", generationRequest.SchemaID, generationRequest.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		log.NormalReturn()
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}
