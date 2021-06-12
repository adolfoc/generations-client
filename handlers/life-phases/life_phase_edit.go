package life_phases

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditLifePhase(w http.ResponseWriter, r *http.Request) {
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

	generationSchema, err := handlers.GetGenerationSchema(w, r, schemaID)
	if err != nil {
		log.FailedReturn()
		return
	}

	momentTypeID, err := getUrlLifePhaseID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	momentType, err := getSchemaLifePhase(w, r, schemaID, momentTypeID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	momentTypeRequest := buildLifePhaseRequest(momentType)

	url := fmt.Sprintf("/schemas/%d/life-phases/%d/update", schemaID, momentTypeID)
	momentTypeForm, err := MakeLifePhaseForm(w, r, url, GetLabel(LifePhaseEditPageTitleIndex), generationSchema.MakeStudyTitle(),
		GetLabel(LifePhaseEditSubmitLabelIndex), momentType, momentTypeRequest, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(LifePhaseFormTemplate, momentTypeForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func EditLifePhaseRetry(w http.ResponseWriter, r *http.Request, lifePhaseRequest *model.LifePhaseRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-life_phases", "EditLifePhaseRetry")

	schemaID := lifePhaseRequest.SchemaID

	generationSchema, err := handlers.GetGenerationSchema(w, r, schemaID)
	if err != nil {
		log.FailedReturn()
		return
	}

	momentType := &model.LifePhase{
		ID:        lifePhaseRequest.ID,
		SchemaID:  lifePhaseRequest.SchemaID,
		Name:      lifePhaseRequest.Name,
		StartYear: lifePhaseRequest.StartYear,
		EndYear:   lifePhaseRequest.EndYear,
		Role:      lifePhaseRequest.Role,
	}

	url := fmt.Sprintf("/schemas/%d/life-phases/%d/update", schemaID, lifePhaseRequest.ID)
	lifePhaseForm, err := MakeLifePhaseForm(w, r, url, GetLabel(LifePhaseEditPageTitleIndex), generationSchema.MakeStudyTitle(),
		GetLabel(LifePhaseEditSubmitLabelIndex), momentType, lifePhaseRequest, errors)
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
