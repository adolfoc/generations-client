package users

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourceUser = "users"
)

func getUsersURL(page int) string {
	return fmt.Sprintf("%s%s?page[number]=%d", handlers.GetAPIHostURL(), ResourceUser, page)
}

func getSimpleUsersURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourceUser)
}

func getUserURL(userID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceUser, userID)
}
 func getChangePasswordURL(userID int) string {
	 return fmt.Sprintf("%s%s/%d/change-password", handlers.GetAPIHostURL(), ResourceUser, userID)
 }

 func getUsers(w http.ResponseWriter, r *http.Request, page int) (*model.Users, error) {
	url := getUsersURL(page)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var users *model.Users
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return users, nil
}

func getUser(w http.ResponseWriter, r *http.Request, groupID int) (*model.User, error) {
	url := getUserURL(groupID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var user *model.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return user, nil
}

func getUrlUserID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("user_id", w, r)
}

func buildNewUserRequest(user *model.User) *model.NewUserRequest {
	ur := &model.NewUserRequest{
		ID:         user.ID,
		UserName:   user.UserName,
		Role:       user.Role,
		Email:      user.Email,
		Password:   user.Password,
		FirstNames: user.FirstNames,
		LastNames:  user.LastNames,
		IsActive:   false,
	}

	return ur
}

func buildUpdateUserRequest(user *model.User) *model.UpdateUserRequest {
	ur := &model.UpdateUserRequest{
		ID:         user.ID,
		UserName:   user.UserName,
		Role:       user.Role,
		Email:      user.Email,
		FirstNames: user.FirstNames,
		LastNames:  user.LastNames,
		IsActive:   user.IsActive,
	}

	return ur
}

func buildChangePasswordRequest(user *model.User) *model.ChangePasswordRequest {
	cpr := &model.ChangePasswordRequest{
		ID:         user.ID,
	}

	return cpr
}

func makeNewUserRequest(r *http.Request) (*model.NewUserRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	userName := handlers.GetStringFormValue(r, "inputUserName")
	role := handlers.GetStringFormValue(r, "inputRole")
	email := handlers.GetStringFormValue(r, "inputEmail")
	password := handlers.GetStringFormValue(r, "inputPassword")
	passwordConfirmation := handlers.GetStringFormValue(r, "inputPasswordConfirmation")
	firstNames := handlers.GetStringFormValue(r, "inputFirstNames")
	lastNames := handlers.GetStringFormValue(r, "inputLastNames")
	isActive := handlers.GetBoolFormValue(r, "inputIsActive")

	nur := &model.NewUserRequest{
		ID:                   normalizedID,
		UserName:             userName,
		Role:                 role,
		Email:                email,
		Password:             password,
		PasswordConfirmation: passwordConfirmation,
		FirstNames:           firstNames,
		LastNames:            lastNames,
		IsActive:             isActive,
	}

	return nur, nil
}

func makeUpdateUserRequest(r *http.Request) (*model.UpdateUserRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	userName := handlers.GetStringFormValue(r, "inputUserName")
	role := handlers.GetStringFormValue(r, "inputRole")
	email := handlers.GetStringFormValue(r, "inputEmail")
	firstNames := handlers.GetStringFormValue(r, "inputFirstNames")
	lastNames := handlers.GetStringFormValue(r, "inputLastNames")
	isActive := handlers.GetBoolFormValue(r, "inputIsActive")

	nur := &model.UpdateUserRequest{
		ID:         normalizedID,
		UserName:   userName,
		Role:       role,
		Email:      email,
		FirstNames: firstNames,
		LastNames:  lastNames,
		IsActive:   isActive,
	}

	return nur, nil
}

func makeChangePasswordRequest(r *http.Request) (*model.ChangePasswordRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	oldPassword := handlers.GetStringFormValue(r, "inputOldPassword")
	newPassword := handlers.GetStringFormValue(r, "inputNewPassword")
	newPasswordConfirmation := handlers.GetStringFormValue(r, "inputNewPasswordConfirmation")

	cpr := &model.ChangePasswordRequest{
		ID:                      normalizedID,
		OldPassword:             oldPassword,
		NewPassword:             newPassword,
		NewPasswordConfirmation: newPasswordConfirmation,
	}

	return cpr, nil
}

func postUser(w http.ResponseWriter, r *http.Request, newUserRequest *model.NewUserRequest) (int, []byte, error) {
	url := getSimpleUsersURL()

	payload, err := json.Marshal(newUserRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchUser(w http.ResponseWriter, r *http.Request, userRequest *model.UpdateUserRequest) (int, []byte, error) {
	url := getUserURL(userRequest.ID)

	payload, err := json.Marshal(userRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}

func changePassword(w http.ResponseWriter, r *http.Request, cprRequest *model.ChangePasswordRequest) (int, []byte, error) {
	url := getChangePasswordURL(cprRequest.ID)

	payload, err := json.Marshal(cprRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PutResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}
