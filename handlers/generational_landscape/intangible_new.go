package generational_landscape

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func NewIntangible(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generational-landscapes", "NewIntangible")

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

	intangible := &model.Intangible{
		ID:          0,
		LandscapeID: generationalLandscapeID,
	}

	tangibleRequest := newIntangibleRequest(generationalLandscapeID)

	url := fmt.Sprintf("/schemas/%d/generational-landscape/%d/intangibles/create", schemaID, generationalLandscapeID)
	intangibleForm, err := MakeIntangibleForm(w, r, url, GetLabel(IntangibleNewPageTitleIndex), generationSchema.MakeStudyTitle(),
		GetLabel(IntangibleNewSubmitLabelIndex), intangible, tangibleRequest, schemaID, generationalLandscape.GenerationID,
		handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(IntangibleFormTemplate, intangibleForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func NewIntangibleRetry(w http.ResponseWriter, r *http.Request, intangibleRequest *model.IntangibleRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-generational-landscapes", "NewIntangibleRetry")

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

	intangible := &model.Intangible{
		ID:          0,
		LandscapeID: intangibleRequest.LandscapeID,
		Name:        intangibleRequest.Name,
		Description: intangibleRequest.Description,
	}

	generationalLandscape, err := getGenerationalLandscape(w, r, schemaID, intangibleRequest.LandscapeID)
	if err != nil {
		log.FailedReturn()
		return
	}

	url := fmt.Sprintf("/schemas/%d/generational-landscape/%d/intangibles/create", schemaID, intangibleRequest.LandscapeID)
	intangibleForm, err := MakeIntangibleForm(w, r, url, GetLabel(IntangibleNewPageTitleIndex), generationSchema.MakeStudyTitle(),
		GetLabel(IntangibleNewSubmitLabelIndex), intangible, intangibleRequest, schemaID, generationalLandscape.GenerationID, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(IntangibleFormTemplate, intangibleForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}


