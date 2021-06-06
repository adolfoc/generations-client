package schemas

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func GenerateTemplate(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-schemas", "GenerateTemplate")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	gsID, err := getUrlGenerationSchemaID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, _, err := generateTemplate(w, r, gsID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	//if code == http.StatusUnprocessableEntity {
	//	responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(PlaceCreateErrorsReceivedIndex))
	//	NewPlaceRetry(w, r, placeRequest, *responseErrors)
	//	log.FailedReturn()
	//	return
	//}

	if code == http.StatusCreated || code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(GenerationSchemaTemplateGeneratedIndex))
		url := fmt.Sprintf("/schemas/%d", gsID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}

	log.NormalReturn()
}

