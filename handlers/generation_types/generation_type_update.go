package generation_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdateGenerationType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generation_types", "UpdateGenerationType")

	generationTypeRequest, err := makeGenerationTypeRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchGenerationType(w, r, generationTypeRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(GenerationTypeUpdateErrorsReceivedIndex))
		EditGenerationTypeRetry(w, r, generationTypeRequest, *responseErrors)
		log.NormalReturn()
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(GenerationTypeUpdatedIndex))
		url := fmt.Sprintf("/schemas/%d", generationTypeRequest.SchemaID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		log.NormalReturn()
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}
