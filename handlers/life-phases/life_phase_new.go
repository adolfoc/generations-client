package life_phases

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func NewLifePhase(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-life_phases", "EditLifePhase")

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

	lifePhaseRequest := newLifePhaseRequest(schemaID)
	lifePhase := &model.LifePhase{
		SchemaID:    schemaID,
	}

	url := fmt.Sprintf("/schemas/%d/life-phases/create", schemaID)
	lifePhaseForm, err := MakeLifePhaseForm(w, r, url, GetLabel(LifePhaseNewPageTitleIndex),
		GetLabel(LifePhaseNewSubmitLabelIndex), lifePhase, lifePhaseRequest, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(LifePhaseFormTemplate, lifePhaseForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func NewLifePhaseRetry(w http.ResponseWriter, r *http.Request, lifePhaseRequest *model.LifePhaseRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-life_phases", "NewLifePhaseRetry")

	schemaID := lifePhaseRequest.SchemaID

	lifePhase := &model.LifePhase{
		ID:        lifePhaseRequest.ID,
		SchemaID:  lifePhaseRequest.SchemaID,
		Name:      lifePhaseRequest.Name,
		StartYear: lifePhaseRequest.StartYear,
		EndYear:   lifePhaseRequest.EndYear,
		Role:      lifePhaseRequest.Role,
	}

	url := fmt.Sprintf("/schemas/%d/life-phases/create", schemaID)
	lifePhaseForm, err := MakeLifePhaseForm(w, r, url, GetLabel(LifePhaseNewPageTitleIndex),
		GetLabel(LifePhaseNewSubmitLabelIndex), lifePhase, lifePhaseRequest, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(LifePhaseFormTemplate, lifePhaseForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}


