package moment_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func NewMomentType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-moment_types", "EditMomentType")

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

	momentTypeRequest := newMomentTypeRequest(schemaID)
	momentType := &model.MomentType{
		SchemaID:    schemaID,
	}

	url := fmt.Sprintf("/schemas/%d/moment-types/create", schemaID)
	momentTypeForm, err := MakeMomentTypeForm(w, r, url, GetLabel(MomentTypeNewPageTitleIndex),
		GetLabel(MomentTypeNewSubmitLabelIndex), momentType, momentTypeRequest, handlers.ResponseErrors{})
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

func NewMomentTypeRetry(w http.ResponseWriter, r *http.Request, momentTypeRequest *model.MomentTypeRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-moment_types", "NewMomentTypeRetry")

	schemaID := momentTypeRequest.SchemaID

	momentType := &model.MomentType{
		ID:          0,
		SchemaID:    momentTypeRequest.SchemaID,
		Name:        momentTypeRequest.Name,
		Description: momentTypeRequest.Description,
	}

	url := fmt.Sprintf("/schemas/%d/moment-types/create", schemaID)
	momentTypeForm, err := MakeMomentTypeForm(w, r, url, GetLabel(MomentTypeNewPageTitleIndex),
		GetLabel(MomentTypeNewSubmitLabelIndex), momentType, momentTypeRequest, errors)
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

