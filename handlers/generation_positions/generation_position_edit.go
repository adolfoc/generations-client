package generation_positions

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditGenerationPosition(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generation_positions", "EditGenerationPosition")

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

	generationPositionID, err := getUrlGenerationPositionID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	lifePhases, err := getSchemaLifePhases(w, r, schemaID)
	if err != nil {
		log.FailedReturn()
		return
	}

	moments, err := getSchemaMoments(w, r, schemaID)
	if err != nil {
		log.FailedReturn()
		return
	}

	generationPosition, err := getSchemaGenerationPosition(w, r, schemaID, generationPositionID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	generationPositionRequest := buildGenerationPositionRequest(generationPosition)

	url := fmt.Sprintf("/schemas/%d/generations/%d/generation-positions/%d/update", schemaID, generationID, generationPositionID)
	generationPositionForm, err := MakeGenerationPositionForm(w, r, url, GetLabel(GenerationPositionEditPageTitleIndex),
		GetLabel(GenerationPositionEditSubmitLabelIndex), generationPosition, generationPositionRequest, schemaID,
		lifePhases, moments, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(GenerationPositionFormTemplate, generationPositionForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func EditGenerationPositionRetry(w http.ResponseWriter, r *http.Request, gpRequest *model.GenerationPositionRequest,
	schemaID int, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-generation_positions", "EditGenerationPositionRetry")

	generationID, err := getUrlGenerationID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	lifePhases, err := getSchemaLifePhases(w, r, schemaID)
	if err != nil {
		log.FailedReturn()
		return
	}

	moments, err := getSchemaMoments(w, r, schemaID)
	if err != nil {
		log.FailedReturn()
		return
	}

	var lifePhase *model.LifePhase
	if gpRequest.LifePhaseID > 0 {
		lifePhase, _ = getLifePhase(w, r, gpRequest.LifePhaseID)
	}

	var generation *model.Generation
	if gpRequest.GenerationID > 0 {
		generation, _ = getGeneration(w, r, gpRequest.GenerationID)
	}

	generationPosition := &model.GenerationPosition{
		ID:         gpRequest.ID,
		MomentID:   gpRequest.MomentID,
		Name:       gpRequest.Name,
		Ordinal:    gpRequest.Ordinal,
		LifePhase:  lifePhase,
		Generation: generation,
	}

	url := fmt.Sprintf("/schemas/%d/generations/%d/generation-positions/%d/update", schemaID, generationID, gpRequest.ID)
	generationPositionForm, err := MakeGenerationPositionForm(w, r, url, GetLabel(GenerationPositionEditPageTitleIndex),
		GetLabel(GenerationPositionEditSubmitLabelIndex), generationPosition, gpRequest, schemaID, lifePhases, moments, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(GenerationPositionFormTemplate, generationPositionForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}


