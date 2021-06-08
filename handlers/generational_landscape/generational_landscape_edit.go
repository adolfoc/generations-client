package generational_landscape

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditGenerationalLandscape(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generational_landscape", "EditGenerationalLandscape")

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
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	moments, err := getMomentsForSchema(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	glRequest := buildGenerationalLandscapeRequest(generationalLandscape)

	url := fmt.Sprintf("/schemas/%d/generational-landscape/%d/update", schemaID, generationalLandscapeID)
	generationalLandscapeForm, err := MakeGenerationalLandscapeForm(w, r, url, GetLabel(GenerationalLandscapeEditPageTitleIndex),
		GetLabel(GenerationalLandscapeEditSubmitLabelIndex), generationalLandscape, glRequest, schemaID, moments, handlers.ResponseErrors{})
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

func EditGenerationalLandscapeRetry(w http.ResponseWriter, r *http.Request, glRequest *model.GenerationalLandscapeRequest,
	schemaID int, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-generational_landscape", "EditGenerationalLandscapeRetry")

	moments, err := getMomentsForSchema(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	generationalLandscape := &model.GenerationalLandscape{
		ID:                glRequest.ID,
		GenerationID:      glRequest.GenerationID,
		FormationMomentID: glRequest.FormationMomentID,
		Description:       glRequest.Description,
	}

	url := fmt.Sprintf("/schemas/%d/generational-landscape/%d/update", schemaID, glRequest.ID)
	generationalLandscapeForm, err := MakeGenerationalLandscapeForm(w, r, url, GetLabel(GenerationalLandscapeEditPageTitleIndex),
		GetLabel(GenerationalLandscapeEditSubmitLabelIndex), generationalLandscape, glRequest, schemaID, moments, errors)
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
