package persons

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
	"strings"
)

const (
	ResourcePerson = "persons"
)

func getPersonsURL(page int, column string) string {
	return fmt.Sprintf("%s%s?page[number]=%d&sort[column]=%s", handlers.GetAPIHostURL(), ResourcePerson, page, column)
}

func getSimplePersonsURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourcePerson)
}

func getPersonURL(personID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourcePerson, personID)
}

func getGenerateLifeSegmentsURL(personID, schemaID int) string {
	return fmt.Sprintf("%s%s/%d/create-life-segments/%d", handlers.GetAPIHostURL(), ResourcePerson, personID, schemaID)
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

func convertToArray(value string) []string {
	return strings.Split(value, ",")
}

func makePersonRequest(r *http.Request) (*model.PersonRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	names := handlers.GetStringFormValue(r, "inputNames")
	aliases := handlers.GetStringFormValue(r, "inputAliases")
	knownAs := handlers.GetStringFormValue(r, "inputKnownAs")
	sex := handlers.GetIntFormValue(r, "inputSex")
	birthDate := handlers.GetStringFormValue(r, "inputBirthDate")
	birthPlaceID := handlers.GetIntFormValue(r, "inputBirthPlaceID")
	deathDate := handlers.GetStringFormValue(r, "inputDeathDate")
	deathPlaceID := handlers.GetIntFormValue(r, "inputDeathPlaceID")
	summary := handlers.GetStringFormValue(r, "inputSummary")

	er := &model.PersonRequest{
		ID:           normalizedID,
		Names:        names,
		Aliases:      convertToArray(aliases),
		KnownAs:      convertToArray(knownAs),
		Sex:          sex,
		BirthDate:    birthDate,
		BirthPlaceID: birthPlaceID,
		DeathDate:    deathDate,
		DeathPlaceID: deathPlaceID,
		Summary:      summary,
	}

	return er, nil
}

func getFormSchemaID(w http.ResponseWriter, r *http.Request) (int, error) {
	schemaID := handlers.GetIntFormValue(r, "inputSchemaID")

	return schemaID, nil
}

func postPerson(w http.ResponseWriter, r *http.Request, personRequest *model.PersonRequest) (int, []byte, error) {
	url := getSimplePersonsURL()

	payload, err := json.Marshal(personRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchPerson(w http.ResponseWriter, r *http.Request, personRequest *model.PersonRequest) (int, []byte, error) {
	url := getPersonURL(personRequest.ID)

	payload, err := json.Marshal(personRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}

	return code, body, nil
}

func generateLifeSegments(w http.ResponseWriter, r *http.Request, personID, schemaID int) (int, []byte, error) {
	url := getGenerateLifeSegmentsURL(personID, schemaID)

	code, body, err := handlers.PutResource(w, r, url, []byte{})
	if err != nil {
		return 0, nil, err
	}

	return code, body, nil
}