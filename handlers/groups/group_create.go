package groups

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-groups", "CreateGroup")

	groupRequest, err := makeGroupRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postGroup(w, r, groupRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(GroupCreateErrorsReceivedIndex))
		NewGroupRetry(w, r, groupRequest, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		var group *model.Group
		_ = json.Unmarshal(body, &group)

		handlers.WriteSessionInfoMessage(r, GetLabel(GroupCreatedIndex))
		url := fmt.Sprintf("/groups/%d", group.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}

	log.NormalReturn()
}


