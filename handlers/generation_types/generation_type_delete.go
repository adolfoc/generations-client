package generation_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func DeleteGenerationType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generation_types", "DeleteGenerationType")

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

	generationTypeID, err := getUrlGenerationTypeID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := deleteGenerationType(w, r, schemaID, generationTypeID)
	if err != nil {
		log.FailedReturn()
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(GenerationTypeDeletedIndex))
	} else {
		responseErrors, _ := handlers.OnDeleteError(r, body)
		handlers.WriteSessionErrorMessage(r, handlers.ExtractFirstError(responseErrors))
	}

	url := fmt.Sprintf("/schemas/%d", schemaID)
	http.Redirect(w, r, url, http.StatusMovedPermanently)

	log.NormalReturn()
}

