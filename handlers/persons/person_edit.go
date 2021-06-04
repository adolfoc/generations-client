package persons

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditPerson(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-persons", "EditPerson")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	personID, err := getUrlPersonID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	person, err := getPerson(w, r, personID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	places, err := getPlaces(w, r)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	personRequest := buildPersonRequest(person)

	url := fmt.Sprintf("/persons/%d/update", personID)
	personForm, err := MakePersonForm(w, r, url, GetLabel(PersonEditPageTitleIndex),
		GetLabel(PersonEditSubmitLabelIndex), person, personRequest, places, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(PersonFormTemplate, personForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func EditPersonRetry(w http.ResponseWriter, r *http.Request, personRequest *model.PersonRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-persons", "EditPersonRetry")

	birthPlace, err := getPlace(w, r, personRequest.BirthPlaceID)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	deathPlace, err := getPlace(w, r, personRequest.DeathPlaceID)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	person := &model.Person{
		ID:           personRequest.ID,
		Names:        personRequest.Names,
		Aliases:      personRequest.Aliases,
		KnownAs:      personRequest.KnownAs,
		Sex:          personRequest.Sex,
		BirthDate:    personRequest.BirthDate,
		BirthPlace:   birthPlace,
		DeathDate:    personRequest.DeathDate,
		DeathPlace:   deathPlace,
		Summary:      personRequest.Summary,
	}

	places, err := getPlaces(w, r)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	url := fmt.Sprintf("/persons/%d/update", personRequest.ID)
	personForm, err := MakePersonForm(w, r, url, GetLabel(PersonEditPageTitleIndex),
		GetLabel(PersonEditSubmitLabelIndex), person, personRequest, places, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(PersonFormTemplate, personForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}

