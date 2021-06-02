package authentication

import (
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-authentication", "Authenticate")

	authenticationRequest := &model.AuthenticationRequest{}

	af := MakeFormAuthentication("/request-authentication", GetLabel(PageTitleIndex),
		GetLabel(SubmitLabelIndex), authenticationRequest, handlers.ResponseErrors{}, "")

	err := handlers.ExecuteView(FormAuthenticationTemplate, af, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func AuthenticateRetry(w http.ResponseWriter, authenticationRequest *model.AuthenticationRequest, errorMessage string) error {
	log := common.StartLog("handlers-authentication", "AuthenticateRetry")

	af := MakeFormAuthentication("/request-authentication", GetLabel(PageTitleIndex),
		GetLabel(SubmitLabelIndex), authenticationRequest, handlers.ResponseErrors{}, errorMessage)

	err := handlers.ExecuteView(FormAuthenticationTemplate, af, w)
	if err != nil {
		log.FailedReturn()
		return err
	}

	log.NormalReturn()
	return nil
}
