package place_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func NewPlaceType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-place-types", "NewPlaceType")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	placeTypeRequest := newPlaceTypeRequest()

	url := fmt.Sprintf("/places-types/create")
	placeForm, err := MakePlaceTypeForm(w, r, url, GetLabel(PlaceTypeNewPageTitleIndex),
		GetLabel(PlaceTypeNewSubmitLabelIndex), &model.PlaceType{}, placeTypeRequest, handlers.ResponseErrors{})
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
}

func NewPlaceTypeRetry(w http.ResponseWriter, r *http.Request, ptRequest *model.PlaceTypeRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-place-types", "NewPlaceTypeRetry")

	placeType := &model.PlaceType{
		ID:          ptRequest.ID,
		Name:        ptRequest.Name,
		Description: ptRequest.Description,
	}

	url := fmt.Sprintf("/place-types/create")
	placeTypeForm, err := MakePlaceTypeForm(w, r, url, GetLabel(PlaceTypeNewPageTitleIndex),
		GetLabel(PlaceTypeNewSubmitLabelIndex), placeType, ptRequest, errors)
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
	return
}
