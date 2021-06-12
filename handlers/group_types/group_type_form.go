package group_types

import (
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	GroupTypeFormTemplate = "group_type_form"
)

type GroupTypeForm struct {
	Ft        handlers.FormTemplate
	GroupType *model.GroupType
}

func makeGroupTypeFormValues(gtRequest *model.GroupTypeRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = gtRequest.ID
	formValues["Name"] = gtRequest.Name
	formValues["Description"] = gtRequest.Description

	return formValues
}

func makeGroupTypeErrorMessages(errors handlers.ResponseErrors) map[string]string {
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

func MakeGroupTypeForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, studyTitle, submitLabel string, groupType *model.GroupType,
	gtRequest *model.GroupTypeRequest, errors handlers.ResponseErrors) (*GroupTypeForm, error) {

	formValues := makeGroupTypeFormValues(gtRequest)
	formErrorMessages := makeGroupTypeErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, studyTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	gtTemplate := &GroupTypeForm{
		Ft:        *ft,
		GroupType: groupType,
	}

	return gtTemplate, nil
}
