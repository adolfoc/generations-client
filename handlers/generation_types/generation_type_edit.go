package generation_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditGenerationType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generation_types", "EditGenerationType")

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

	generationSchema, err := handlers.GetGenerationSchema(w, r, schemaID)
	if err != nil {
		log.FailedReturn()
		return
	}

	generationTypeID, err := getUrlGenerationTypeID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	generationType, err := getSchemaGenerationType(w, r, schemaID, generationTypeID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	generationTypeRequest := buildGenerationTypeRequest(generationType)

	url := fmt.Sprintf("/schemas/%d/generation-types/%d/update", schemaID, generationTypeID)
	generationTypeForm, err := MakeGenerationTypeForm(w, r, url, GetLabel(GenerationTypeEditPageTitleIndex),
		generationSchema.MakeStudyTitle(), GetLabel(GenerationTypeEditSubmitLabelIndex),
		generationType, generationTypeRequest, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(GenerationTypeFormTemplate, generationTypeForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func EditGenerationTypeRetry(w http.ResponseWriter, r *http.Request, generationTypeRequest *model.GenerationTypeRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-generation_types", "EditGenerationTypeRetry")

	schemaID := generationTypeRequest.SchemaID

	generationSchema, err := handlers.GetGenerationSchema(w, r, schemaID)
	if err != nil {
		log.FailedReturn()
		return
	}

	generationType := &model.GenerationType{
		ID:          generationTypeRequest.ID,
		SchemaID:    generationTypeRequest.SchemaID,
		Archetype:   generationTypeRequest.Archetype,
		Description: generationTypeRequest.Description,
	}

	url := fmt.Sprintf("/schemas/%d/generation-types/%d/update", schemaID, generationTypeRequest.ID)
	generationTypeForm, err := MakeGenerationTypeForm(w, r, url, GetLabel(GenerationTypeEditPageTitleIndex),
		generationSchema.MakeStudyTitle(), GetLabel(GenerationTypeEditSubmitLabelIndex),
		generationType, generationTypeRequest, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(GenerationTypeFormTemplate, generationTypeForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}
