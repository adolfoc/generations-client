package generational_landscape

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func NewGenerationalLandscape(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generational-landscapes", "NewGenerationalLandscape")

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

	generationID, err := getUrlGenerationID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	generation, err := getGeneration(w, r, generationID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	moments, err := getMomentsForSchema(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	generationalLandscape := &model.GenerationalLandscape{
		GenerationID:      generationID,
	}

	generationalLandscapeRequest := newGenerationalLandscapeRequest(generationID)

	url := fmt.Sprintf("/schemas/%d/generational-landscape/create", schemaID)
	generationalLandscapeForm, err := MakeGenerationalLandscapeForm(w, r, url, GetLabel(GenerationalLandscapeNewPageTitleIndex),
		generationSchema.MakeStudyTitle(), GetLabel(GenerationalLandscapeNewSubmitLabelIndex), generationalLandscape, generationalLandscapeRequest, schemaID, moments,
		generation, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(GenerationalLandscapeFormTemplate, generationalLandscapeForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func NewGenerationalLandscapeRetry(w http.ResponseWriter, r *http.Request, glRequest *model.GenerationalLandscapeRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-generational-landscapes", "NewGenerationalLandscapeRetry")

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

	moments, err := getMomentsForSchema(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	generation, err := getGeneration(w, r, glRequest.GenerationID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	generationalLandscape := &model.GenerationalLandscape{
		ID:                0,
		GenerationID:      glRequest.GenerationID,
		FormationMomentID: glRequest.FormationMomentID,
		Description:       glRequest.Description,
	}

	url := fmt.Sprintf("/schemas/%d/generational-landscape/create", schemaID)
	generationalLandscapeForm, err := MakeGenerationalLandscapeForm(w, r, url, GetLabel(GenerationalLandscapeNewPageTitleIndex),
		generationSchema.MakeStudyTitle(), GetLabel(GenerationalLandscapeNewSubmitLabelIndex),
		generationalLandscape, glRequest, schemaID, moments, generation, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(GenerationalLandscapeFormTemplate, generationalLandscapeForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}
