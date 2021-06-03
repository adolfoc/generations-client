package moments

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditMoment(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-moments", "EditMoment")

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

	momentID, err := getUrlMomentID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	moment, err := getSchemaMoment(w, r, schemaID, momentID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	momentTypes, err := getMomentTypesForSchema(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	momentRequest := buildMomentRequest(moment)

	url := fmt.Sprintf("/schemas/%d/moment/%d/update", schemaID, momentID)
	generationForm, err := MakeMomentForm(w, r, url, GetLabel(MomentEditPageTitleIndex),
		GetLabel(MomentEditSubmitLabelIndex), moment, momentRequest, momentTypes, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(MomentFormTemplate, generationForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func EditMomentRetry(w http.ResponseWriter, r *http.Request, momentRequest *model.HistoricalMomentRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-moment", "EditMomentRetry")

	schemaID := momentRequest.SchemaID
	momentTypes, err := getMomentTypesForSchema(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}
	momentType := matchMomentType(momentRequest.TypeID, momentTypes)

	moment := &model.HistoricalMoment{
		ID:              momentRequest.ID,
		Name:            momentRequest.Name,
		SchemaID:        momentRequest.SchemaID,
		Type:            momentType,
		Start:           momentRequest.Start,
		End:             momentRequest.End,
		Summary:         momentRequest.Summary,
		Description:     momentRequest.Description,
	}

	url := fmt.Sprintf("/schemas/%d/moments/%d/update", schemaID, momentRequest.ID)
	generationForm, err := MakeMomentForm(w, r, url, GetLabel(MomentEditPageTitleIndex),
		GetLabel(MomentEditSubmitLabelIndex), moment, momentRequest, momentTypes, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(MomentFormTemplate, generationForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}


