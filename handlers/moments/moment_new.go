package moments

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func NewMoment(w http.ResponseWriter, r *http.Request) {
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

	generationSchema, err := handlers.GetGenerationSchema(w, r, schemaID)
	if err != nil {
		log.FailedReturn()
		return
	}

	momentTypes, err := getMomentTypesForSchema(w, r, schemaID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	momentRequest := newMomentRequest(schemaID)

	url := fmt.Sprintf("/schemas/%d/moments/create", schemaID)
	momentForm, err := MakeMomentForm(w, r, url, GetLabel(MomentNewPageTitleIndex), generationSchema.MakeStudyTitle(),
		GetLabel(MomentNewSubmitLabelIndex), &model.HistoricalMoment{}, momentRequest, momentTypes, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(MomentFormTemplate, momentForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func NewMomentRetry(w http.ResponseWriter, r *http.Request, momentRequest *model.HistoricalMomentRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-moments", "NewMomentRetry")

	schemaID := momentRequest.SchemaID

	generationSchema, err := handlers.GetGenerationSchema(w, r, schemaID)
	if err != nil {
		log.FailedReturn()
		return
	}

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

	url := fmt.Sprintf("/schemas/%d/moments/create", schemaID)
	momentForm, err := MakeMomentForm(w, r, url, GetLabel(MomentNewPageTitleIndex), generationSchema.MakeStudyTitle(),
		GetLabel(MomentNewSubmitLabelIndex), moment, momentRequest, momentTypes, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(MomentFormTemplate, momentForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}

const (
	GenerationSchemaURL = "generation-schemas"
)

func getMomentTypesForSchema(w http.ResponseWriter, r *http.Request, schemaID int) ([]*model.MomentType, error) {
	url := fmt.Sprintf("%s%s/%d/moment-types", handlers.GetAPIHostURL(), GenerationSchemaURL, schemaID)
	code, body, err := handlers.GetResource(w, r, url)

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	if err != nil {
		return nil, err
	}

	var momentTypes []*model.MomentType
	err = json.Unmarshal(body, &momentTypes)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return momentTypes, nil
}

func matchMomentType(typeID int, momentTypes []*model.MomentType) *model.MomentType {
	for _, mt := range momentTypes {
		if mt.ID == typeID {
			return mt
		}
	}

	return nil
}
