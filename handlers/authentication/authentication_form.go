package authentication

import (
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
)

const (
	FormAuthenticationTemplate = "login_form"
)

type FormAuthentication struct {
	Ft                    handlers.FormTemplate
	AuthenticationRequest *model.AuthenticationRequest
	ErrorMessage          string
}

func makeFormAuthenticationValues(authenticationRequest *model.AuthenticationRequest) map[string]interface{} {
	formValues := make(map[string]interface{})
	formValues["Email"] = authenticationRequest.Email
	formValues["Password"] = authenticationRequest.Password

	return formValues
}

func makeAuthenticationErrorMessages(errors handlers.ResponseErrors) map[string]string {
	formErrorMessages := make(map[string]string)
	for _, error := range errors.Errors {
		if error.Source.Pointer == "data/attributes/email" {
			formErrorMessages["Email"] = error.Detail
		} else if error.Source.Pointer == "data/attributes/password" {
			formErrorMessages["Password"] = error.Detail
		}
	}

	return formErrorMessages
}

func MakeFormAuthentication(url string, pageTitle, submitLabel string, authenticationRequest *model.AuthenticationRequest,
	errors handlers.ResponseErrors, errorMessage string) *FormAuthentication {

	formValues := makeFormAuthenticationValues(authenticationRequest)
	formErrorMessages := makeAuthenticationErrorMessages(errors)
	ft := handlers.MakeSimpleFormTemplate(url, pageTitle, submitLabel, formValues, formErrorMessages)

	authenticationTemplate := &FormAuthentication{
		Ft:                    *ft,
		AuthenticationRequest: authenticationRequest,
		ErrorMessage:          errorMessage,
	}

	return authenticationTemplate
}

