package group_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func CreateGroupType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-group_types", "CreateGroupType")

	groupRequest, err := makeGroupTypeRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postGroupType(w, r, groupRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(GroupTypeCreateErrorsReceivedIndex))
		NewGroupTypeRetry(w, r, groupRequest, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(GroupTypeCreatedIndex))
		url := fmt.Sprintf("/group-types/%d", groupRequest.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}

	log.NormalReturn()
}

