package users

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-users", "CreateUser")

	newUserRequest, err := makeNewUserRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postUser(w, r, newUserRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(UserCreateErrorsReceivedIndex))
		NewUserRetry(w, r, newUserRequest, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		var user *model.User
		_ = json.Unmarshal(body, &user)

		handlers.WriteSessionInfoMessage(r, GetLabel(UserCreatedIndex))
		url := fmt.Sprintf("/users/%d", user.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}

	log.NormalReturn()
}
