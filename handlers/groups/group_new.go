package groups

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func NewGroup(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-groups", "NewGroup")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	groupTypes, err := getGroupTypes(w, r)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	groupRequest := &model.GroupRequest{}

	url := fmt.Sprintf("/groups/create")
	groupForm, err := MakeGroupForm(w, r, url, GetLabel(GroupNewPageTitleIndex), "",
		GetLabel(GroupNewSubmitLabelIndex), &model.Group{}, groupRequest, groupTypes, handlers.ResponseErrors{})
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(GroupFormTemplate, groupForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

func NewGroupRetry(w http.ResponseWriter, r *http.Request, groupRequest *model.GroupRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-groups", "NewGroupRetry")

	groupTypes, err := getGroupTypes(w, r)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}
	groupType := matchGroupType(groupRequest.GroupTypeID, groupTypes)

	group := &model.Group{
		ID:            groupRequest.ID,
		Name:          groupRequest.Name,
		GroupTypeID:   groupRequest.GroupTypeID,
		ParentGroupID: groupRequest.ParentGroupID,
		Start:         groupRequest.Start,
		End:           groupRequest.End,
		Summary:       groupRequest.Summary,
		Description:   groupRequest.Description,
		GroupType:     groupType,
	}

	url := fmt.Sprintf("/groups/create")
	groupForm, err := MakeGroupForm(w, r, url, GetLabel(GroupNewPageTitleIndex), "",
		GetLabel(GroupNewSubmitLabelIndex), group, groupRequest, groupTypes, errors)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView(GroupFormTemplate, groupForm, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
	return
}
