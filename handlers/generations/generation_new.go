package generations

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func NewGeneration(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generations", "EditGeneration")

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

	generationTypes, err := getGenerationTypesForSchema(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	places, err := getPlacesForSchema(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	moments, err := getMomentsForSchema(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	generationRequest := newGenerationRequest(schemaID)

	url := fmt.Sprintf("/schemas/%d/generations/create", schemaID)
	generationForm, err := MakeGenerationForm(w, r, url, GetLabel(GenerationNewPageTitleIndex),
		GetLabel(GenerationNewSubmitLabelIndex), &model.Generation{}, generationRequest, generationTypes, places, moments,
		handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(GenerationFormTemplate, generationForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func NewGenerationRetry(w http.ResponseWriter, r *http.Request, generationRequest *model.GenerationRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-generations", "NewGenerationRetry")

	schemaID := generationRequest.SchemaID
	generationTypes, err := getGenerationTypesForSchema(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}
	generationType := matchGenerationType(generationRequest.TypeID, generationTypes)

	places, err := getPlacesForSchema(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}
	place := matchPlace(generationRequest.PlaceID, places)

	moments, err := getMomentsForSchema(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	generation := &model.Generation{
		ID:                   generationRequest.ID,
		Name:                 generationRequest.Name,
		SchemaID:             generationRequest.SchemaID,
		Type:                 generationType,
		StartYear:            generationRequest.StartYear,
		EndYear:              generationRequest.EndYear,
		Place:                place,
		Description:          generationRequest.Description,
		FormationLandscapeID: generationRequest.FormationLandscapeID,
	}

	url := fmt.Sprintf("/schemas/%d/generations/create", schemaID)
	generationForm, err := MakeGenerationForm(w, r, url, GetLabel(GenerationNewPageTitleIndex),
		GetLabel(GenerationNewSubmitLabelIndex), generation, generationRequest, generationTypes, places, moments, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(GenerationFormTemplate, generationForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}
