package group_types

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func CreateGroupType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-group_types", "CreateGroupType")

	groupTypeRequest, err := makeGroupTypeRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	code, body, err := postGroupType(w, r, groupTypeRequest)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	if code == http.StatusUnprocessableEntity {
		responseErrors, _ := handlers.OnUpdateError(r, body, GetLabel(GroupTypeCreateErrorsReceivedIndex))
		NewGroupTypeRetry(w, r, groupTypeRequest, *responseErrors)
		log.FailedReturn()
		return
	}

	if code == http.StatusCreated || code == http.StatusOK {
		var groupType *model.GroupType
		_ = json.Unmarshal(body, &groupType)

		handlers.WriteSessionInfoMessage(r, GetLabel(GroupTypeCreatedIndex))
		url := fmt.Sprintf("/group-types/%d", groupType.ID)
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	} else {
		fmt.Fprintf(w, "received %d", code)
		return
	}

	log.NormalReturn()
}

