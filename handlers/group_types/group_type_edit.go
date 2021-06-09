package group_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditGroupType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-group_types", "EditGroupType")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	eventID, err := getUrlGroupTypeID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	event, err := getGroupType(w, r, eventID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	gtRequest := buildGroupTypeRequest(event)

	url := fmt.Sprintf("/group-types/%d/update", eventID)
	groupTypeForm, err := MakeGroupTypeForm(w, r, url, GetLabel(GroupTypeEditPageTitleIndex),
		GetLabel(GroupTypeEditSubmitLabelIndex), event, gtRequest, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(GroupTypeFormTemplate, groupTypeForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func EditGroupTypeRetry(w http.ResponseWriter, r *http.Request, gtRequest *model.GroupTypeRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-group_types", "EditGroupTypeRetry")

	group := &model.GroupType{
		ID:          gtRequest.ID,
		Name:        gtRequest.Name,
		Description: gtRequest.Description,
	}

	url := fmt.Sprintf("/group-types/%d/update", gtRequest.ID)
	groupTypeForm, err := MakeGroupTypeForm(w, r, url, GetLabel(GroupTypeEditPageTitleIndex),
		GetLabel(GroupTypeEditSubmitLabelIndex), group, gtRequest, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(GroupTypeFormTemplate, groupTypeForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}
