package place_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditPlaceType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-place-types", "EditPlaceType")

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

	placeType, err := getPlaceType(w, r, placeTypeID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	placeTypeRequest := buildPlaceTypeRequest(placeType)

	url := fmt.Sprintf("/place-types/%d/update", placeTypeID)
	placeTypeForm, err := MakePlaceTypeForm(w, r, url, GetLabel(PlaceTypeEditPageTitleIndex),
		GetLabel(PlaceTypeEditSubmitLabelIndex), placeType, placeTypeRequest, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(PlaceTypeFormTemplate, placeTypeForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func EditPlaceTypeRetry(w http.ResponseWriter, r *http.Request, placeTypeRequest *model.PlaceTypeRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-place-types", "EditPlaceTypeRetry")

	placeType := &model.PlaceType{
		ID:            placeTypeRequest.ID,
		Name:          placeTypeRequest.Name,
		Description:   placeTypeRequest.Description,
	}

	url := fmt.Sprintf("/place-types/%d/update", placeTypeRequest.ID)
	placeForm, err := MakePlaceTypeForm(w, r, url, GetLabel(PlaceTypeEditPageTitleIndex),
		GetLabel(PlaceTypeEditSubmitLabelIndex), placeType, placeTypeRequest, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(PlaceTypeFormTemplate, placeForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}
