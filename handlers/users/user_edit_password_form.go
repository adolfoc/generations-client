package users

import (
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	EditPasswordFormTemplate = "edit_password_form"
)

type EditPasswordForm struct {
	Ft   handlers.FormTemplate
	User *model.User
}

func makeEditPasswordFormValues(cpRequest *model.ChangePasswordRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["ID"] = cpRequest.ID
	formValues["OldPassword"] = cpRequest.OldPassword
	formValues["NewPassword"] = cpRequest.NewPassword
	formValues["NewPasswordConfirmation"] = cpRequest.NewPasswordConfirmation

	return formValues
}

func makeEditPasswordErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/old_password" {
			formErrorMessages["OldPassword"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/new_password" {
			formErrorMessages["NewPassword"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/new_password_confirmation" {
			formErrorMessages["NewPasswordConfirmation"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakeEditPasswordForm(w http.ResponseWriter, r *http.Request, url string, pageTitle, submitLabel string, user *model.User,
	cpRequest *model.ChangePasswordRequest, errors handlers.ResponseErrors) (*EditPasswordForm, error) {

	formValues := makeEditPasswordFormValues(cpRequest)
	formErrorMessages := makeEditPasswordErrorMessages(errors)
	ft, err := handlers.MakeFormTemplate(r, url, pageTitle, submitLabel, formValues, formErrorMessages)
	if err != nil {
		return nil, err
	}

	epTemplate := &EditPasswordForm{
		Ft:   *ft,
		User: user,
	}

	return epTemplate, nil
}
