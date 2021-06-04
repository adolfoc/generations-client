package handlers

import (
	"github.com/adolfoc/generations-client/common"
	"net/http"
)

type ErrorPageTemplate struct {
	Ct           CommonTemplate
	GeneralError *GeneralError
}

func MakeErrorPageTemplate(pageTitle string, generalError *GeneralError) *ErrorPageTemplate {
	ct := MakeSimpleTemplate(pageTitle)

	cardTemplate := &ErrorPageTemplate{
		Ct:           *ct,
		GeneralError: generalError,
	}

	return cardTemplate
}

func GetErrorPage(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-errors", "GetErrorPage")

	if UserAuthenticated(w, r) == false {
		log.FailedReturn()
		RedirectToLogin(w, r)
		return
	}

	session, err := GetCurrentSessionFromRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	ct := MakeErrorPageTemplate("Se ha producido un error", session.GeneralError)

	err = ExecuteView("error_page", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

