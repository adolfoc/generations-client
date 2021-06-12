package place_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	PlaceTypeFormTemplate = "place_type_form"
)

type PlaceTypeForm struct {
	Ft        handlers.FormTemplate
	PlaceType *model.PlaceType
}

func makePlaceTypeFormValues(ptRequest *model.PlaceTypeRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = fmt.Sprintf("%d", ptRequest.ID)
	formValues["Name"] = ptRequest.Name
	formValues["Description"] = ptRequest.Description

	return formValues
}

func makePlaceTypeErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/name" {
			formErrorMessages["Name"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/description" {
			formErrorMessages["Description"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakePlaceTypeForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, studyTitle, submitLabel string, placeType *model.PlaceType,
	ptRequest *model.PlaceTypeRequest, errors handlers.ResponseErrors) (*PlaceTypeForm, error) {

	formValues := makePlaceTypeFormValues(ptRequest)
	formErrorMessages := makePlaceTypeErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, studyTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	ptTemplate := &PlaceTypeForm{
		Ft:        *ft,
		PlaceType: placeType,
	}

	return ptTemplate, nil
}

