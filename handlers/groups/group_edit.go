package groups

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func EditGroup(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-groups", "EditGroup")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	eventID, err := getUrlGroupID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	event, err := getGroup(w, r, eventID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	groupTypes, err := getGroupTypes(w, r)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	gRequest := buildGroupRequest(event)

	url := fmt.Sprintf("/groups/%d/update", eventID)
	groupForm, err := MakeGroupForm(w, r, url, GetLabel(GroupEditPageTitleIndex),
		GetLabel(GroupEditSubmitLabelIndex), event, gRequest, groupTypes, handlers.ResponseErrors{})
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

func EditGroupRetry(w http.ResponseWriter, r *http.Request, gRequest *model.GroupRequest, errors handlers.ResponseErrors) {
	log := common.StartLog("handlers-groups", "EditGroupRetry")

	groupTypes, err := getGroupTypes(w, r)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}
	groupType := matchGroupType(gRequest.GroupTypeID, groupTypes)

	group := &model.Group{
		ID:            gRequest.ID,
		Name:          gRequest.Name,
		GroupTypeID:   groupType.ID,
		ParentGroupID: gRequest.ParentGroupID,
		Start:         gRequest.Start,
		End:           gRequest.End,
		Summary:       gRequest.Summary,
		Members:       nil,
		Description:   gRequest.Description,
		GroupType:     groupType,
		GroupName:     "",
	}

	url := fmt.Sprintf("/group-types/%d/update", gRequest.ID)
	groupForm, err := MakeGroupForm(w, r, url, GetLabel(GroupEditPageTitleIndex),
		GetLabel(GroupEditSubmitLabelIndex), group, gRequest, groupTypes, errors)
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

func getGroupTypes(w http.ResponseWriter, r *http.Request) ([]*model.GroupType, error) {
	url := fmt.Sprintf("%sgroup-types", handlers.GetAPIHostURL())
	code, body, err := handlers.GetResource(w, r, url)

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	if err != nil {
		return nil, err
	}

	var groupTypes *model.GroupTypes
	err = json.Unmarshal(body, &groupTypes)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return groupTypes.GroupTypes, nil
}

func matchGroupType(typeID int, groupTypes []*model.GroupType) *model.GroupType {
	for _, et := range groupTypes {
		if et.ID == typeID {
			return et
		}
	}

	return nil
}


