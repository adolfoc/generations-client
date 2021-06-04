package moments

import (
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type MomentTemplate struct {
	Ct            handlers.CommonTemplate
	SchemaID      int
	Moment        *model.HistoricalMoment
	Constellation []*model.GenerationPosition
}

func MakeMomentTemplate(r *http.Request, pageTitle string, moment *model.HistoricalMoment, positions []*model.GenerationPosition) (*MomentTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle)
	if err != nil {
		return nil, err
	}

	momentTemplate := &MomentTemplate{
		Ct:            *ct,
		SchemaID:      moment.SchemaID,
		Moment:        moment,
		Constellation: positions,
	}

	return momentTemplate, nil
}

func GetMoment(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-moments", "GetMoment")

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

	positions, err := getGenerationPositions(w, r, schemaID, momentID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakeMomentTemplate(r, GetLabel(MomentPageTitleIndex), moment, positions)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("moment", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}
