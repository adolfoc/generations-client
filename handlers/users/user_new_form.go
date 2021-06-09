package users

import (
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	NewUserFormTemplate = "user_new_form"
)

type NewUserForm struct {
	Ft   handlers.FormTemplate
	User *model.User
}

func makeNewUserFormValues(nuRequest *model.NewUserRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = nuRequest.ID
	formValues["UserName"] = nuRequest.UserName
	formValues["Role"] = nuRequest.Role
	formValues["Email"] = nuRequest.Email
	formValues["Password"] = nuRequest.Password
	formValues["PasswordConfirmation"] = nuRequest.PasswordConfirmation
	formValues["FirstNames"] = nuRequest.FirstNames
	formValues["LastNames"] = nuRequest.LastNames
	formValues["IsActive"] = nuRequest.IsActive

	return formValues
}

func makeNewUserErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/user_name" {
			formErrorMessages["UserName"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/role" {
			formErrorMessages["Role"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/email" {
			formErrorMessages["Email"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/password" {
			formErrorMessages["Password"] = error.Detail
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

func MakeNewUserForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, submitLabel string, user *model.User,
	nuRequest *model.NewUserRequest, errors handlers.ResponseErrors) (*NewUserForm, error) {

	formValues := makeNewUserFormValues(nuRequest)
	formErrorMessages := makeNewUserErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	nuTemplate := &NewUserForm{
		Ft:   *ft,
		User: user,
	}

	return nuTemplate, nil
}

