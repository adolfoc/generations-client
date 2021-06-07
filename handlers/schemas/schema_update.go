package schemas

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdateSchema(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-persons", "UpdateSchema")

	schemaRequest, err := makeSchemaRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchSchema(w, r, schemaRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(GenerationSchemaUpdateErrorsReceivedIndex))
		EditSchemaRetry(w, r, schemaRequest, *responseErrors)
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(GenerationSchemaUpdatedIndex))
		url := fmt.Sprintf("/schemas/%d", schemaRequest.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}
