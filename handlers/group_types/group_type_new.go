package group_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func NewGroupType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-group_types", "NewGroupType")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	groupTypeRequest := &model.GroupTypeRequest{}

	url := fmt.Sprintf("/group-types/create")
	groupTypeForm, err := MakeGroupTypeForm(w, r, url, GetLabel(GroupTypeNewPageTitleIndex), "",
		GetLabel(GroupTypeNewSubmitLabelIndex), &model.GroupType{}, groupTypeRequest, handlers.ResponseErrors{})
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

func NewGroupTypeRetry(w http.ResponseWriter, r *http.Request, groupTypeRequest *model.GroupTypeRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-group_types", "NewGroupTypeRetry")

	groupType := &model.GroupType{
		ID:          groupTypeRequest.ID,
		Name:        groupTypeRequest.Name,
		Description: groupTypeRequest.Description,
	}

	url := fmt.Sprintf("/group-types/create")
	groupForm, err := MakeGroupTypeForm(w, r, url, GetLabel(GroupTypeNewPageTitleIndex), "",
		GetLabel(GroupTypeNewSubmitLabelIndex), groupType, groupTypeRequest, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(GroupTypeFormTemplate, groupForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}

