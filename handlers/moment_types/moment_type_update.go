package moment_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdateMomentType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-moment_types", "UpdateMomentType")

	momentTypeRequest, err := makeMomentTypeRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchMomentType(w, r, momentTypeRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(MomentTypeUpdateErrorsReceivedIndex))
		EditMomentTypeRetry(w, r, momentTypeRequest, *responseErrors)
		log.NormalReturn()
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(MomentTypeUpdatedIndex))
		url := fmt.Sprintf("/schemas/%d", momentTypeRequest.SchemaID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		log.NormalReturn()
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}
