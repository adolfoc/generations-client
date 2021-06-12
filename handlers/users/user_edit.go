package users

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditUser(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-users", "EditUser")

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

	euRequest := buildUpdateUserRequest(user)

	url := fmt.Sprintf("/users/%d/update", userID)
	editUserForm, err := MakeEditUserForm(w, r, url, GetLabel(UserEditPageTitleIndex), "",
		GetLabel(UserEditSubmitLabelIndex), user, euRequest, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(EditUserFormTemplate, editUserForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func EditUserRetry(w http.ResponseWriter, r *http.Request, euRequest *model.UpdateUserRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-group_types", "EditUserRetry")

	user := &model.User{
		ID:         euRequest.ID,
		UserName:   euRequest.UserName,
		Role:       euRequest.Role,
		Email:      euRequest.Email,
		FirstNames: euRequest.FirstNames,
		LastNames:  euRequest.LastNames,
		IsActive:   euRequest.IsActive,
	}

	url := fmt.Sprintf("/users/%d/update", euRequest.ID)
	editUserForm, err := MakeEditUserForm(w, r, url, GetLabel(UserEditPageTitleIndex), "",
		GetLabel(UserEditSubmitLabelIndex), user, euRequest, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(EditUserFormTemplate, editUserForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}
