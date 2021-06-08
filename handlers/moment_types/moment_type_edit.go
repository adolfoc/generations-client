package moment_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditMomentType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-generation_types", "EditMomentType")

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

	momentTypeID, err := getUrlMomentTypeID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	momentType, err := getSchemaMomentType(w, r, schemaID, momentTypeID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	momentTypeRequest := buildMomentTypeRequest(momentType)

	url := fmt.Sprintf("/schemas/%d/moment-types/%d/update", schemaID, momentTypeID)
	momentTypeForm, err := MakeMomentTypeForm(w, r, url, GetLabel(MomentTypeEditPageTitleIndex),
		GetLabel(MomentTypeEditSubmitLabelIndex), momentType, momentTypeRequest, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(MomentTypeFormTemplate, momentTypeForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func EditMomentTypeRetry(w http.ResponseWriter, r *http.Request, momentTypeRequest *model.MomentTypeRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-generation_types", "EditMomentTypeRetry")

	schemaID := momentTypeRequest.SchemaID

	momentType := &model.MomentType{
		ID:          momentTypeRequest.ID,
		SchemaID:    momentTypeRequest.SchemaID,
		Name:        momentTypeRequest.Name,
		Description: momentTypeRequest.Description,
	}

	url := fmt.Sprintf("/schemas/%d/moment-types/%d/update", schemaID, momentTypeRequest.ID)
	momentTypeForm, err := MakeMomentTypeForm(w, r, url, GetLabel(MomentTypeEditPageTitleIndex),
		GetLabel(MomentTypeEditSubmitLabelIndex), momentType, momentTypeRequest, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(MomentTypeFormTemplate, momentTypeForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}
