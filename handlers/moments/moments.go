package moments

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type MomentsTemplate struct {
	Ct         handlers.CommonTemplate
	SchemaID   int
	Moments    *model.HistoricalMoments
	Pagination *handlers.Pagination
}

func getPaginationBaseURL(generationSchemaID int) string {
	stem := fmt.Sprintf("/schemas/%d/moments", generationSchemaID)
	return stem + "?page=%d"
}

func MakeMomentsTemplate(r *http.Request, pageTitle string, generationSchema, page int, moments *model.HistoricalMoments) (*MomentsTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle)
	if err != nil {
		return nil, err
	}

	pagination := handlers.MakePagination(moments.RecordCount, getPaginationBaseURL(generationSchema), page)

	generationTemplate := &MomentsTemplate{
		Ct:         *ct,
		SchemaID:   generationSchema,
		Moments:    moments,
		Pagination: pagination,
	}

	return generationTemplate, nil
}

func GetSchemaMoments(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-schemas", "GetSchemaMoments")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	gsID, err := getUrlGenerationSchemaID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	page := handlers.GetURLPageParameter(r)
	generations, err := getSchemaMoments(w, r, gsID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakeMomentsTemplate(r, GetLabel(MomentIndexPageTitleIndex), gsID, page, generations)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("moments", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}
