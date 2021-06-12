package users

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditPassword(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-users", "EditPassword")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	userID, err := getUrlUserID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	user, err := getUser(w, r, userID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	cpRequest := buildChangePasswordRequest(user)

	url := fmt.Sprintf("/users/%d/update-password", userID)
	editPasswordForm, err := MakeEditPasswordForm(w, r, url, GetLabel(UserChangePasswordTitleIndex), "",
		GetLabel(UserChangePasswordSubmitLabelIndex), user, cpRequest, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(EditPasswordFormTemplate, editPasswordForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func EditPasswordRetry(w http.ResponseWriter, r *http.Request, cprRequest *model.ChangePasswordRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-users", "EditPasswordRetry")

	user, err := getUser(w, r, cprRequest.ID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	url := fmt.Sprintf("/users/%d/update-password", cprRequest.ID)
	editPasswordForm, err := MakeEditPasswordForm(w, r, url, GetLabel(UserChangePasswordTitleIndex), "",
		GetLabel(UserChangePasswordSubmitLabelIndex), user, cprRequest, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(EditPasswordFormTemplate, editPasswordForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}

