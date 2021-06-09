package users

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func NewUser(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-users", "NewUser")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	newUserRequest := &model.NewUserRequest{}

	url := fmt.Sprintf("/users/create")
	newUserForm, err := MakeNewUserForm(w, r, url, GetLabel(UserNewPageTitleIndex),
		GetLabel(UserNewSubmitLabelIndex), &model.User{}, newUserRequest, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(NewUserFormTemplate, newUserForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func NewUserRetry(w http.ResponseWriter, r *http.Request, newUserRequest *model.NewUserRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-users", "NewUserRetry")

	user := &model.User{
		ID:         0,
		UserName:   newUserRequest.UserName,
		Role:       newUserRequest.Role,
		Email:      newUserRequest.Email,
		Password:   newUserRequest.Password,
		FirstNames: newUserRequest.FirstNames,
		LastNames:  newUserRequest.LastNames,
		IsActive:   newUserRequest.IsActive,
	}

	url := fmt.Sprintf("/users/create")
	newUserForm, err := MakeNewUserForm(w, r, url, GetLabel(UserNewPageTitleIndex),
		GetLabel(UserNewSubmitLabelIndex), user, newUserRequest, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(NewUserFormTemplate, newUserForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}
