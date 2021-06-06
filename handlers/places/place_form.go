package places

import (
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	PlaceFormTemplate = "place_form"
)

type PlaceForm struct {
	Ft         handlers.FormTemplate
	Place      *model.Place
	Parents    []*model.Place
	PlaceTypes []*model.PlaceType
}

func makePlaceFormValues(pRequest *model.PlaceRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = fmt.Sprintf("%d", pRequest.ID)
	formValues["Name"] = pRequest.Name
	formValues["PlaceTypeID"] = pRequest.PlaceTypeID
	formValues["ParentPlaceID"] = pRequest.ParentPlaceID
	formValues["Start"] = pRequest.Start
	formValues["End"] = pRequest.End
	formValues["Summary"] = pRequest.Summary
	formValues["Description"] = pRequest.Description

	return formValues
}

func makePlaceErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/name" {
			formErrorMessages["Name"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/place_type_id" {
			formErrorMessages["PlaceTypeID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/parent_place_id" {
			formErrorMessages["ParentPlaceID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/start" {
			formErrorMessages["Start"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/end" {
			formErrorMessages["End"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/summary" {
			formErrorMessages["Summary"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/description" {
			formErrorMessages["Description"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakePlaceForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, submitLabel string, place *model.Place,
	pRequest *model.PlaceRequest, placeTypes []*model.PlaceType, parents []*model.Place, errors handlers.ResponseErrors) (*PlaceForm, error) {

	formValues := makePlaceFormValues(pRequest)
	formErrorMessages := makePlaceErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	lfTemplate := &PlaceForm{
		Ft:         *ft,
		Place:      place,
		PlaceTypes: placeTypes,
		Parents:    parents,
	}

	return lfTemplate, nil
}

