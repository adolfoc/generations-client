package persons

import (
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
	"strings"
)

const (
	PersonFormTemplate = "person_form"
)

type PersonForm struct {
	Ft     handlers.FormTemplate
	Person *model.Person
	Places []*model.Place
}

func convertArrayToString(array []string) string {
	return strings.Join(array, ", ")
}

func makePersonFormValues(pRequest *model.PersonRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = fmt.Sprintf("%d", pRequest.ID)
	formValues["Names"] = pRequest.Names
	formValues["Aliases"] = convertArrayToString(pRequest.Aliases)
	formValues["KnownAs"] = convertArrayToString(pRequest.KnownAs)
	formValues["Sex"] = fmt.Sprintf("%d", pRequest.Sex)
	formValues["BirthDate"] = pRequest.BirthDate
	formValues["BirthPlaceID"] = pRequest.BirthPlaceID
	formValues["DeathDate"] = pRequest.DeathDate
	formValues["DeathPlaceID"] = pRequest.DeathPlaceID
	formValues["Summary"] = pRequest.Summary

	return formValues
}

func makePersonErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/names" {
			formErrorMessages["Names"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/aliases" {
			formErrorMessages["Aliases"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/known_as" {
			formErrorMessages["KnownAs"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/sex" {
			formErrorMessages["Sex"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/birth_date" {
			formErrorMessages["BirthDate"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/birth_place_id" {
			formErrorMessages["BirthPlaceID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/death_date" {
			formErrorMessages["DeathDate"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/death_place_id" {
			formErrorMessages["DeathPlaceID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/summary" {
			formErrorMessages["Summary"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakePersonForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, studyTitle, submitLabel string, person *model.Person,
	lsRequest *model.PersonRequest, places []*model.Place, errors handlers.ResponseErrors) (*PersonForm, error) {

	formValues := makePersonFormValues(lsRequest)
	formErrorMessages := makePersonErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, studyTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	lfTemplate := &PersonForm{
		Ft:     *ft,
		Person: person,
		Places: places,
	}

	return lfTemplate, nil
}
