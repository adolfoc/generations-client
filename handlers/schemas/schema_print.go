package schemas

import (
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func PrintSchema(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-schemas", "PrintSchema")

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

	document, err := printSchema(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	w.Header().Add("Content-type", "application/pdf")
	w.Write(document)
	log.NormalReturn()
}

