package users

import (
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type UserTemplate struct {
	Ct   handlers.CommonTemplate
	User *model.User
}

func MakeUserTemplate(r *http.Request, pageTitle string, user *model.User) (*UserTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle)
	if err != nil {
		return nil, err
	}

	userTemplate := &UserTemplate{
		Ct:   *ct,
		User: user,
	}

	return userTemplate, nil
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-users", "GetUser")

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

	ct, err := MakeUserTemplate(r, GetLabel(UserPageTitleIndex), user)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("user", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}


