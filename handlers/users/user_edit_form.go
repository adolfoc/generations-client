package users

import (
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	EditUserFormTemplate = "user_edit_form"
)

type EditUserForm struct {
	Ft   handlers.FormTemplate
	User *model.User
}

func makeEditUserFormValues(nuRequest *model.UpdateUserRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = nuRequest.ID
	formValues["UserName"] = nuRequest.UserName
	formValues["Role"] = nuRequest.Role
	formValues["Email"] = nuRequest.Email
	formValues["FirstNames"] = nuRequest.FirstNames
	formValues["LastNames"] = nuRequest.LastNames
	formValues["IsActive"] = nuRequest.IsActive

	return formValues
}

func makeEditUserErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/user_name" {
			formErrorMessages["UserName"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/role" {
			formErrorMessages["Role"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/email" {
			formErrorMessages["Email"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/first_names" {
			formErrorMessages["FirstNames"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/last_names" {
			formErrorMessages["LastNames"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/is_active" {
			formErrorMessages["IsActive"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakeEditUserForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, submitLabel string, user *model.User,
	nuRequest *model.UpdateUserRequest, errors handlers.ResponseErrors) (*EditUserForm, error) {

	formValues := makeEditUserFormValues(nuRequest)
	formErrorMessages := makeEditUserErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	euTemplate := &EditUserForm{
		Ft:   *ft,
		User: user,
	}

	return euTemplate, nil
}

