package persons

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func GenerateLifeSegments(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-persons", "GenerateLifeSegments")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	personID, err := getUrlPersonID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	schemaID, err := getFormSchemaID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, _, err := generateLifeSegments(w, r, personID, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusOK || code == http.StatusCreated {
		handlers.WriteSessionInfoMessage(r, GetLabel(PersonUpdatedIndex))
		url := fmt.Sprintf("/persons/%d", personID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}
