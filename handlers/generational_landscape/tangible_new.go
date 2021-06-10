package generational_landscape

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func NewTangible(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generational-landscapes", "NewTangible")

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
	if err != nil {
		log.FailedReturn()
		return
	}

	tangible := &model.Tangible{
		ID:          0,
		LandscapeID: generationalLandscapeID,
	}

	tangibleRequest := newTangibleRequest(generationalLandscapeID)

	url := fmt.Sprintf("/schemas/%d/generational-landscape/%d/tangibles/create", schemaID, generationalLandscapeID)
	tangibleForm, err := MakeTangibleForm(w, r, url, GetLabel(TangibleNewPageTitleIndex),
		GetLabel(TangibleNewSubmitLabelIndex), tangible, tangibleRequest, schemaID, generationalLandscape.GenerationID,
		handlers.ResponseErrors{})
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

func NewTangibleRetry(w http.ResponseWriter, r *http.Request, tangibleRequest *model.TangibleRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-generational-landscapes", "NewTangibleRetry")

	schemaID, err := getUrlGenerationSchemaID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	tangible := &model.Tangible{
		ID:          0,
		LandscapeID: tangibleRequest.LandscapeID,
		Name:        tangibleRequest.Name,
		Description: tangibleRequest.Description,
	}

	generationalLandscape, err := getGenerationalLandscape(w, r, schemaID, tangibleRequest.LandscapeID)
	if err != nil {
		log.FailedReturn()
		return
	}

	url := fmt.Sprintf("/schemas/%d/generational-landscape/%d/tangibles/create", schemaID, tangibleRequest.LandscapeID)
	tangibleForm, err := MakeTangibleForm(w, r, url, GetLabel(TangibleNewPageTitleIndex),
		GetLabel(TangibleNewSubmitLabelIndex), tangible, tangibleRequest, schemaID, generationalLandscape.GenerationID, errors)
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
