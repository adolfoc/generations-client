package generational_landscape

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func DeleteTangible(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generational_landscape", "DeleteTangible")

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

	tangibleID, err := getUrlTangibleID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := deleteTangible(w, r, tangibleID)
	if err != nil {
		log.FailedReturn()
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(TangibleDeletedIndex))
	} else {
		responseErrors, _ := handlers.OnDeleteError(r, body)
		handlers.WriteSessionErrorMessage(r, handlers.ExtractFirstError(responseErrors))
	}

	url := fmt.Sprintf("/schemas/%d/generations/%d", schemaID, generationalLandscape.GenerationID)
	http.Redirect(w, r, url, http.StatusMovedPermanently)

	log.NormalReturn()
}

