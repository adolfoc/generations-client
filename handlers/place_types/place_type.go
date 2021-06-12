package place_types

import (
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type PlaceTypeTemplate struct {
	Ct        handlers.CommonTemplate
	PlaceType *model.PlaceType
}

func MakePlaceTypeTemplate(r *http.Request, pageTitle string, placeType *model.PlaceType) (*PlaceTypeTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle, "")
	if err != nil {
		return nil, err
	}

	placeTypeTemplate := &PlaceTypeTemplate{
		Ct:        *ct,
		PlaceType: placeType,
	}

	return placeTypeTemplate, nil
}

func GetPlaceType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-places_types", "GetPlace")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	placeTypeID, err := getUrlPlaceTypeID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	place, err := getPlaceType(w, r, placeTypeID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakePlaceTypeTemplate(r, GetLabel(PlaceTypePageTitleIndex), place)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("place_type", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

