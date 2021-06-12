package generation_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func NewGenerationType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generation_types", "NewGenerationType")

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

	generationTypeRequest := newGenerationTypeRequest(schemaID)
	generationType := &model.GenerationType{
		SchemaID:    schemaID,
	}

	url := fmt.Sprintf("/schemas/%d/generation-types/create", schemaID)
	generationForm, err := MakeGenerationTypeForm(w, r, url, GetLabel(GenerationTypeNewPageTitleIndex),
		generationSchema.MakeStudyTitle(), GetLabel(GenerationTypeNewSubmitLabelIndex),
		generationType, generationTypeRequest, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(GenerationTypeFormTemplate, generationForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func NewGenerationTypeRetry(w http.ResponseWriter, r *http.Request, generationTypeRequest *model.GenerationTypeRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-generation_types", "NewGenerationTypeRetry")

	schemaID := generationTypeRequest.SchemaID

	generationSchema, err := handlers.GetGenerationSchema(w, r, schemaID)
	if err != nil {
		log.FailedReturn()
		return
	}

	generationType := &model.GenerationType{
		ID:          0,
		SchemaID:    generationTypeRequest.SchemaID,
		Archetype:   generationTypeRequest.Archetype,
		Description: generationTypeRequest.Description,
	}

	url := fmt.Sprintf("/schemas/%d/generation_types/create", schemaID)
	generationTypeForm, err := MakeGenerationTypeForm(w, r, url, GetLabel(GenerationTypeNewPageTitleIndex),
		generationSchema.MakeStudyTitle(), GetLabel(GenerationTypeNewSubmitLabelIndex),
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
