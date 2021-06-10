package generational_landscape

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditTangible(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generational_landscape", "EditTangible")

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

	tangibleID, err := getUrlTangibleID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	tangible, err := getTangible(w, r, tangibleID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	generationalLandscape, err := getGenerationalLandscape(w, r, schemaID, tangible.LandscapeID)
	if err != nil {
		log.FailedReturn()
		return
	}

	tRequest := buildTangibleRequest(tangible)

	url := fmt.Sprintf("/schemas/%d/generational-landscape/%d/tangibles/%d/update", schemaID, tRequest.LandscapeID, tRequest.ID)
	tangibleForm, err := MakeTangibleForm(w, r, url, GetLabel(TangibleEditPageTitleIndex),
		GetLabel(TangibleEditSubmitLabelIndex), tangible, tRequest, schemaID, generationalLandscape.GenerationID, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(TangibleFormTemplate, tangibleForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func EditTangibleRetry(w http.ResponseWriter, r *http.Request, tRequest *model.TangibleRequest,
	schemaID int, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-generational_landscape", "EditTangibleRetry")

	tangible := &model.Tangible{
		ID:          tRequest.ID,
		LandscapeID: tRequest.LandscapeID,
		Name:        tRequest.Name,
		Description: tRequest.Description,
	}

	generationalLandscape, err := getGenerationalLandscape(w, r, schemaID, tangible.LandscapeID)
	if err != nil {
		log.FailedReturn()
		return
	}

	url := fmt.Sprintf("/schemas/%d/generational-landscape/%d/tangibles/%d/update", schemaID, tRequest.LandscapeID, tRequest.ID)
	tangibleForm, err := MakeTangibleForm(w, r, url, GetLabel(TangibleEditPageTitleIndex),
		GetLabel(TangibleEditSubmitLabelIndex), tangible, tRequest, schemaID, generationalLandscape.GenerationID, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(TangibleFormTemplate, tangibleForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}
