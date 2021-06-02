package generations

import (
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type GenerationTemplate struct {
	Ct         handlers.CommonTemplate
	SchemaID   int
	Generation *model.Generation
}

func MakeGenerationTemplate(r *http.Request, pageTitle string, generation *model.Generation) (*GenerationTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle)
	if err != nil {
		return nil, err
	}

	generationTemplate := &GenerationTemplate{
		Ct:         *ct,
		SchemaID:   generation.SchemaID,
		Generation: generation,
	}

	return generationTemplate, nil
}

func GetGeneration(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-schemas", "GetGeneration")

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

	generationID, err := getUrlGenerationID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	generation, err := getSchemaGeneration(w, r, schemaID, generationID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakeGenerationTemplate(r, GetLabel(GenerationPageTitleIndex), generation)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("generation", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

