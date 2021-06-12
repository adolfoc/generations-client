package places

import (
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type PlaceTemplate struct {
	Ct    handlers.CommonTemplate
	Place *model.Place
}

func MakePlaceTemplate(r *http.Request, pageTitle, studyTitle string, place *model.Place) (*PlaceTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle, studyTitle)
	if err != nil {
		return nil, err
	}

	placeTemplate := &PlaceTemplate{
		Ct:    *ct,
		Place: place,
	}

	return placeTemplate, nil
}

func GetPlace(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-places", "GetPlace")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	placeID, err := getUrlPlaceID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	place, err := getPlace(w, r, placeID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakePlaceTemplate(r, GetLabel(PlacePageTitleIndex), "", place)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("place", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}
