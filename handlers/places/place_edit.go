package places

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditPlace(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-places", "EditPlace")

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

	placeRequest := buildPlaceRequest(place)

	url := fmt.Sprintf("/places/%d/update", placeID)
	personForm, err := MakePlaceForm(w, r, url, GetLabel(PlaceEditPageTitleIndex), "",
		GetLabel(PlaceEditSubmitPlaceIndex), place, placeRequest, placeTypes, parents.Places, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(PlaceFormTemplate, personForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func EditPlaceRetry(w http.ResponseWriter, r *http.Request, placeRequest *model.PlaceRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-places", "EditPlaceRetry")

	place := &model.Place{
		ID:            placeRequest.ID,
		Name:          placeRequest.Name,
		PlaceTypeID:   placeRequest.PlaceTypeID,
		ParentPlaceID: placeRequest.ParentPlaceID,
		Start:         placeRequest.Start,
		End:           placeRequest.End,
		Summary:       placeRequest.Summary,
		Description:   placeRequest.Description,
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

	url := fmt.Sprintf("/places/%d/update", placeRequest.ID)
	placeForm, err := MakePlaceForm(w, r, url, GetLabel(PlaceEditPageTitleIndex), "",
		GetLabel(PlaceEditSubmitPlaceIndex), place, placeRequest, placeTypes, parents.Places, errors)
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
	return
}
