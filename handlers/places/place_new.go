package places

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func NewPlace(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-places", "NewPlace")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	placeTypes, err := getPlaceTypes(w, r)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	parents, err := getPlaces(w, r, 0, "Name")
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	placeRequest := newPlaceRequest()

	url := fmt.Sprintf("/places/create")
	placeForm, err := MakePlaceForm(w, r, url, GetLabel(PlaceNewPageTitleIndex), "",
		GetLabel(PlaceNewSubmitPlaceIndex), &model.Place{}, placeRequest, placeTypes, parents.Places, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(PlaceFormTemplate, placeForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func NewPlaceRetry(w http.ResponseWriter, r *http.Request, pRequest *model.PlaceRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-places", "NewPlaceRetry")

	place := &model.Place{
		ID:            pRequest.ID,
		Name:          pRequest.Name,
		PlaceTypeID:   pRequest.PlaceTypeID,
		ParentPlaceID: pRequest.ParentPlaceID,
		Start:         pRequest.Start,
		End:           pRequest.End,
		Summary:       pRequest.Summary,
		Description:   pRequest.Description,
	}

	placeTypes, err := getPlaceTypes(w, r)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	parents, err := getPlaces(w, r, 0, "Name")
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	url := fmt.Sprintf("/places/create")
	momentForm, err := MakePlaceForm(w, r, url, GetLabel(PlaceNewPageTitleIndex), "",
		GetLabel(PlaceNewSubmitPlaceIndex), place, pRequest, placeTypes, parents.Places, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(PlaceFormTemplate, momentForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}
