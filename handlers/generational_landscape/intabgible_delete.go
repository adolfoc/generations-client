package generational_landscape

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func DeleteIntangible(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generational_landscape", "DeleteIntangible")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	schemaID, err := getUrlGenerationSchemaID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	generationalLandscapeID, err := getUrlGenerationalLandscapeID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	generationalLandscape, err := getGenerationalLandscape(w, r, schemaID, generationalLandscapeID)
	if err != nil {
		log.FailedReturn()
		return
	}

	intangibleID, err := getUrlIntangibleID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := deleteIntangible(w, r, intangibleID)
	if err != nil {
		log.FailedReturn()
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(IntangibleDeletedIndex))
	} else {
		responseErrors, _ := handlers.OnDeleteError(r, body)
		handlers.WriteSessionErrorMessage(r, handlers.ExtractFirstError(responseErrors))
	}

	url := fmt.Sprintf("/schemas/%d/generations/%d", schemaID, generationalLandscape.GenerationID)
	http.Redirect(w, r, url, http.StatusMovedPermanently)

	log.NormalReturn()
}
