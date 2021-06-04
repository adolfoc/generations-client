package persons

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-persons", "UpdatePerson")

	personRequest, err := makePersonRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchPerson(w, r, personRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(PersonUpdateErrorsReceivedIndex))
		EditPersonRetry(w, r, personRequest, *responseErrors)
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(PersonUpdatedIndex))
		url := fmt.Sprintf("/persons/%d", personRequest.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}
