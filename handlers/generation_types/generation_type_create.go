package generation_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func CreateGenerationType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generation_types", "CreateGenerationType")

	generationTypeRequest, err := makeGenerationTypeRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postGenerationType(w, r, generationTypeRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(GenerationTypeCreateErrorsReceivedIndex))
		NewGenerationTypeRetry(w, r, generationTypeRequest, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(GenerationTypeCreatedIndex))
		url := fmt.Sprintf("/schemas/%d", generationTypeRequest.SchemaID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}
}
