package place_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type PlaceTypesTemplate struct {
	Ct         handlers.CommonTemplate
	PlaceTypes *model.PlaceTypes
	Pagination *handlers.Pagination
}

func getPaginationBaseURL() string {
	stem := fmt.Sprintf("/place-types/index")
	return stem + "?page=%d"
}

func MakePlaceTypesTemplate(r *http.Request, pageTitle, studyTitle string, page int, placeTypes *model.PlaceTypes) (*PlaceTypesTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle, studyTitle)
	if err != nil {
		return nil, err
	}

	pagination := handlers.MakePagination(placeTypes.RecordCount, getPaginationBaseURL(), page)

	eventsTemplate := &PlaceTypesTemplate{
		Ct:         *ct,
		PlaceTypes: placeTypes,
		Pagination: pagination,
	}

	return eventsTemplate, nil
}

func GetPlaceTypes(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-place_types", "GetPlaceTypes")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	page := handlers.GetURLPageParameter(r)
	events, err := getPlaceTypes(w, r, page, "Name")
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakePlaceTypesTemplate(r, GetLabel(PlaceTypeIndexPageTitleIndex), "", page, events)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("place_types", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

