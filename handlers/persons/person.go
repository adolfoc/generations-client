package persons

import (
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type PersonTemplate struct {
	Ct               handlers.CommonTemplate
	Person           *model.Person
	HaveLifeSegments bool
}

func MakePersonTemplate(r *http.Request, pageTitle, studyTitle string, person *model.Person) (*PersonTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle, studyTitle)
	if err != nil {
		return nil, err
	}

	personTemplate := &PersonTemplate{
		Ct:               *ct,
		Person:           person,
		HaveLifeSegments: len(person.LifeSegments) > 0,
	}

	return personTemplate, nil
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-persons", "GetPerson")

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

	ct, err := MakePersonTemplate(r, GetLabel(PersonPageTitleIndex), "", person)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("person", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}
