package groups

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-groups", "UpdateGroup")

	groupRequest, err := makeGroupRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := patchGroup(w, r, groupRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(GroupUpdateErrorsReceivedIndex))
		EditGroupRetry(w, r, groupRequest, *responseErrors)
		return
	}

	if code == http.StatusOK {
		handlers.WriteSessionInfoMessage(r, GetLabel(GroupUpdatedIndex))
		url := fmt.Sprintf("/groups/%d", groupRequest.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		log.FailedReturn()
		return
	}
}
