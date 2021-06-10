package generational_landscape

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditIntangible(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generational_landscape", "EditIntangible")

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

	intangibleID, err := getUrlIntangibleID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	intangible, err := getIntangible(w, r, intangibleID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	generationalLandscape, err := getGenerationalLandscape(w, r, schemaID, intangible.LandscapeID)
	if err != nil {
		log.FailedReturn()
		return
	}

	itRequest := buildIntangibleRequest(intangible)

	url := fmt.Sprintf("/schemas/%d/generational-landscape/%d/intangibles/%d/update", schemaID, itRequest.LandscapeID, itRequest.ID)
	intangibleForm, err := MakeIntangibleForm(w, r, url, GetLabel(IntangibleEditPageTitleIndex),
		GetLabel(IntangibleEditSubmitLabelIndex), intangible, itRequest, schemaID, generationalLandscape.GenerationID, handlers.ResponseErrors{})
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

func EditIntangibleRetry(w http.ResponseWriter, r *http.Request, itRequest *model.IntangibleRequest,
	schemaID int, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-generational_landscape", "EditIntangibleRetry")

	intangible := &model.Intangible{
		ID:          itRequest.ID,
		LandscapeID: itRequest.LandscapeID,
		Name:        itRequest.Name,
		Description: itRequest.Description,
	}

	generationalLandscape, err := getGenerationalLandscape(w, r, schemaID, intangible.LandscapeID)
	if err != nil {
		log.FailedReturn()
		return
	}

	url := fmt.Sprintf("/schemas/%d/generational-landscape/%d/intangibles/%d/update", schemaID, itRequest.LandscapeID, itRequest.ID)
	intangibleForm, err := MakeIntangibleForm(w, r, url, GetLabel(IntangibleEditPageTitleIndex),
		GetLabel(IntangibleEditSubmitLabelIndex), intangible, itRequest, schemaID, generationalLandscape.GenerationID, errors)
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


