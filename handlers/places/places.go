package places

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type PlacesTemplate struct {
	Ct               handlers.CommonTemplate
	Places           *model.Places
	Pagination       *handlers.Pagination
}

func getPaginationBaseURL() string {
	stem := fmt.Sprintf("/places/index")
	return stem + "?page=%d"
}

func MakePlacesTemplate(r *http.Request, pageTitle string, page int, places *model.Places) (*PlacesTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle)
	if err != nil {
		return nil, err
	}

	pagination := handlers.MakePagination(places.RecordCount, getPaginationBaseURL(), page)

	placesTemplate := &PlacesTemplate{
		Ct:         *ct,
		Places:     places,
		Pagination: pagination,
	}

	return placesTemplate, nil
}

func GetPlaces(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-places", "GetPlaces")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	page := handlers.GetURLPageParameter(r)
	places, err := getPlaces(w, r, page, "names")
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakePlacesTemplate(r, GetLabel(PlaceIndexPageTitleIndex), page, places)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("places", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}


