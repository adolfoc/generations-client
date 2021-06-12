package generations

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditGeneration(w http.ResponseWriter, r *http.Request) {
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

	generation, err := getSchemaGeneration(w, r, schemaID, generationID)
	if handlers.HandleError(w, r, err) {
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

	generationRequest := buildGenerationRequest(generation)

	url := fmt.Sprintf("/schemas/%d/generations/%d/update", schemaID, generationID)
	generationForm, err := MakeGenerationForm(w, r, url, GetLabel(GenerationEditPageTitleIndex),
		generationSchema.MakeStudyTitle(), GetLabel(GenerationEditSubmitLabelIndex),
		generation, generationRequest, generationTypes, places, moments, handlers.ResponseErrors{})
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

func EditGenerationRetry(w http.ResponseWriter, r *http.Request, generationRequest *model.GenerationRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-generations", "EditGenerationRetry")

	schemaID := generationRequest.SchemaID

	generationSchema, err := handlers.GetGenerationSchema(w, r, schemaID)
	if err != nil {
		log.FailedReturn()
		return
	}

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

	url := fmt.Sprintf("/schemas/%d/generations/%d/update", schemaID, generationRequest.ID)
	generationForm, err := MakeGenerationForm(w, r, url, GetLabel(GenerationEditPageTitleIndex),
		generationSchema.MakeStudyTitle(), GetLabel(GenerationEditSubmitLabelIndex),
		generation, generationRequest, generationTypes, places, moments, errors)
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

const (
	GenerationSchemaURL = "generation-schemas"
)

func getGenerationTypesForSchema(w http.ResponseWriter, r *http.Request, schemaID int) ([]*model.GenerationType, error) {
	url := fmt.Sprintf("%s%s/%d/generation-types", handlers.GetAPIHostURL(), GenerationSchemaURL, schemaID)
	code, body, err := handlers.GetResource(w, r, url)

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	if err != nil {
		return nil, err
	}

	var generationTypes []*model.GenerationType
	err = json.Unmarshal(body, &generationTypes)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return generationTypes, nil
}

func getPlacesForSchema(w http.ResponseWriter, r *http.Request, schemaID int) ([]*model.Place, error) {
	url := fmt.Sprintf("%splaces", handlers.GetAPIHostURL())
	code, body, err := handlers.GetResource(w, r, url)

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	if err != nil {
		return nil, err
	}

	var places *model.Places
	err = json.Unmarshal(body, &places)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return places.Places, nil
}

func getMomentsForSchema(w http.ResponseWriter, r *http.Request, schemaID int) ([]*model.HistoricalMoment, error) {
	url := fmt.Sprintf("%s%s/%d/moments", handlers.GetAPIHostURL(), GenerationSchemaURL, schemaID)
	code, body, err := handlers.GetResource(w, r, url)

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	if err != nil {
		return nil, err
	}

	var moments *model.HistoricalMoments
	err = json.Unmarshal(body, &moments)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return moments.HistoricalMoments, nil
}

func matchGenerationType(typeID int, generationTypes []*model.GenerationType) *model.GenerationType {
	for _, gt := range generationTypes {
		if gt.ID == typeID {
			return gt
		}
	}

	return nil
}

func matchPlace(placeID int, places []*model.Place) *model.Place {
	for _, place := range places {
		if place.ID == placeID {
			return place
		}
	}

	return nil
}
