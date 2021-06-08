package groups

import (
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	GroupFormTemplate = "group_form"
)

type GroupForm struct {
	Ft         handlers.FormTemplate
	Group      *model.Group
	GroupTypes []*model.GroupType
}

func makeGroupFormValues(gRequest *model.GroupRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = gRequest.ID
	formValues["GroupTypeID"] = gRequest.GroupTypeID
	formValues["ParentGroupID"] = gRequest.ParentGroupID
	formValues["Name"] = gRequest.Name
	formValues["Start"] = gRequest.Start
	formValues["End"] = gRequest.End
	formValues["Summary"] = gRequest.Summary
	formValues["Description"] = gRequest.Description

	return formValues
}

func makeGroupErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/name" {
			formErrorMessages["Name"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/group_type_id" {
			formErrorMessages["GroupTypeID"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/parent_group_id" {
			formErrorMessages["ParentGroupID"] = error.Detail
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

func MakeGroupForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, submitLabel string, group *model.Group,
	gRequest *model.GroupRequest, groupTypes []*model.GroupType, errors handlers.ResponseErrors) (*GroupForm, error) {

	formValues := makeGroupFormValues(gRequest)
	formErrorMessages := makeGroupErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	gTemplate := &GroupForm{
		Ft:         *ft,
		Group:      group,
		GroupTypes: groupTypes,
	}

	return gTemplate, nil
}

