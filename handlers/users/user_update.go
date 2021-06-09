package users

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-users", "UpdateUser")
	updateUserRequest, err := makeUpdateUserRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchUser(w, r, updateUserRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(UserUpdateErrorsReceivedIndex))
		EditUserRetry(w, r, updateUserRequest, *responseErrors)
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(UserUpdatedIndex))
		url := fmt.Sprintf("/users/%d", updateUserRequest.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}
