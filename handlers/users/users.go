package users

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type UsersTemplate struct {
	Ct         handlers.CommonTemplate
	Users      *model.Users
	Pagination *handlers.Pagination
}

func getPaginationBaseURL() string {
	stem := fmt.Sprintf("/users/index")
	return stem + "?page=%d"
}

func MakeUsersTemplate(r *http.Request, pageTitle string, page int, users *model.Users) (*UsersTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle)
	if err != nil {
		return nil, err
	}

	pagination := handlers.MakePagination(users.RecordCount, getPaginationBaseURL(), page)

	usersTemplate := &UsersTemplate{
		Ct:         *ct,
		Users:      users,
		Pagination: pagination,
	}

	return usersTemplate, nil
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-users", "GetUsers")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	page := handlers.GetURLPageParameter(r)
	users, err := getUsers(w, r, page)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakeUsersTemplate(r, GetLabel(UserIndexPageTitleIndex), page, users)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("users", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

