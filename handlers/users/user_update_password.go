package users

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-users", "UpdatePassword")

	cpRequest, err := makeChangePasswordRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := changePassword(w, r, cpRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(UserUpdateErrorsReceivedIndex))
		EditPasswordRetry(w, r, cpRequest, *responseErrors)
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(UserUpdatedIndex))
		url := fmt.Sprintf("/users/%d", cpRequest.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}


