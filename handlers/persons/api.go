package persons

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourcePerson = "persons"
)

func getPersonsURL(page int, column string) string {
	return fmt.Sprintf("%s%s?page[number]=%d&sort[column]=%s", handlers.GetAPIHostURL(), ResourcePerson, page, column)
}

func getPersonURL(personID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourcePerson, personID)
}

func getPersons(w http.ResponseWriter, r *http.Request, page int, sortColumn string) (*model.Persons, error) {
	url := getPersonsURL(page, sortColumn)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var persons *model.Persons
	err = json.Unmarshal(body, &persons)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return persons, nil
}

func getPerson(w http.ResponseWriter, r *http.Request, personID int) (*model.Person, error) {
	url := getPersonURL(personID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var person *model.Person
	err = json.Unmarshal(body, &person)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return person, nil
}

func getUrlPersonID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("person_id", w, r)
}

func newPersonRequest() *model.PersonRequest {
	pRequest := &model.PersonRequest{}

	return pRequest
}

func buildPersonRequest(person *model.Person) *model.PersonRequest {
	personRequest := &model.PersonRequest{
		ID:           person.ID,
		Names:        person.Names,
		Aliases:      person.Aliases,
		KnownAs:      person.KnownAs,
		Sex:          person.Sex,
		BirthDate:    person.BirthDate,
		BirthPlaceID: person.BirthPlace.ID,
		DeathDate:    person.DeathDate,
		DeathPlaceID: person.DeathPlace.ID,
		Summary:      person.Summary,
	}

	return personRequest
}
