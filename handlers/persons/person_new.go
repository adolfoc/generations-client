package persons

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func NewPerson(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-persons", "NewPerson")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	personRequest := newPersonRequest()

	places, err := getPlaces(w, r)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	url := fmt.Sprintf("/persons/create")
	momentForm, err := MakePersonForm(w, r, url, GetLabel(PersonNewPageTitleIndex), "",
		GetLabel(PersonNewSubmitLabelIndex), &model.Person{}, personRequest, places, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(PersonFormTemplate, momentForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func NewPersonRetry(w http.ResponseWriter, r *http.Request, pRequest *model.PersonRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-persons", "NewPersonRetry")

	birthPlace, err := getPlace(w, r, pRequest.BirthPlaceID)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	deathPlace, err := getPlace(w, r, pRequest.BirthPlaceID)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	person := &model.Person{
		ID:           pRequest.ID,
		Names:        pRequest.Names,
		Aliases:      pRequest.Aliases,
		KnownAs:      pRequest.KnownAs,
		Sex:          pRequest.Sex,
		BirthDate:    pRequest.BirthDate,
		BirthPlace:   birthPlace,
		DeathDate:    pRequest.DeathDate,
		DeathPlace:   deathPlace,
		Summary:      pRequest.Summary,
	}

	places, err := getPlaces(w, r)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	url := fmt.Sprintf("/persons/create")
	momentForm, err := MakePersonForm(w, r, url, GetLabel(PersonNewPageTitleIndex), "",
		GetLabel(PersonNewSubmitLabelIndex), person, pRequest, places, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(PersonFormTemplate, momentForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}

func getPlaces(w http.ResponseWriter, r *http.Request) ([]*model.Place, error) {
	url := fmt.Sprintf("%splaces/", handlers.GetAPIHostURL())
	code, body, err := handlers.GetResource(w, r, url)

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	if err != nil {
		return nil, err
	}

	var places *model.Places
	err = json.Unmarshal(body, &places)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return places.Places, nil
}

func getPlace(w http.ResponseWriter, r *http.Request, placeID int) (*model.Place, error) {
	url := fmt.Sprintf("%s/places/%d", handlers.GetAPIHostURL(), placeID)
	code, body, err := handlers.GetResource(w, r, url)

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	if err != nil {
		return nil, err
	}

	var place *model.Place
	err = json.Unmarshal(body, &place)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return place, nil
}

