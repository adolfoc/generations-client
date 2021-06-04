package persons

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type PersonsTemplate struct {
	Ct               handlers.CommonTemplate
	Persons          *model.Persons
	Pagination       *handlers.Pagination
}

func getPaginationBaseURL() string {
	stem := fmt.Sprintf("/persons/index")
	return stem + "?page=%d"
}

func MakePersonsTemplate(r *http.Request, pageTitle string, page int, persons *model.Persons) (*PersonsTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle)
	if err != nil {
		return nil, err
	}

	pagination := handlers.MakePagination(persons.RecordCount, getPaginationBaseURL(), page)

	personsTemplate := &PersonsTemplate{
		Ct:         *ct,
		Persons:    persons,
		Pagination: pagination,
	}

	return personsTemplate, nil
}

func GetPersons(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-persons", "GetPersons")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	page := handlers.GetURLPageParameter(r)
	persons, err := getPersons(w, r, page, "names")
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakePersonsTemplate(r, GetLabel(PersonIndexPageTitleIndex), page, persons)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("persons", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

