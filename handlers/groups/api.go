package groups

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourceGroup = "groups"
)

func getGroupsURL(page int) string {
	return fmt.Sprintf("%s%s?page[number]=%d", handlers.GetAPIHostURL(), ResourceGroup, page)
}

func getSimpleGroupsURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourceGroup)
}

func getGroupURL(groupID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceGroup, groupID)
}

func getGroups(w http.ResponseWriter, r *http.Request, page int) (*model.Groups, error) {
	url := getGroupsURL(page)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var groups *model.Groups
	err = json.Unmarshal(body, &groups)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return groups, nil
}

func getGroup(w http.ResponseWriter, r *http.Request, groupID int) (*model.Group, error) {
	url := getGroupURL(groupID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var group *model.Group
	err = json.Unmarshal(body, &group)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return group, nil
}

func getUrlGroupID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("group_id", w, r)
}

func buildGroupRequest(group *model.Group) *model.GroupRequest {
	mtr := &model.GroupRequest{
		ID:            group.ID,
		Name:          group.Name,
		GroupTypeID:   group.GroupTypeID,
		ParentGroupID: group.ParentGroupID,
		Start:         group.Start,
		End:           group.End,
		Summary:       group.Summary,
		Description:   group.Description,
	}

	return mtr
}

func makeGroupRequest(r *http.Request) (*model.GroupRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	name := handlers.GetStringFormValue(r, "inputName")
	normalizedGroupTypeID := handlers.GetIntFormValue(r, "inputGroupTypeID")
	normalizedParentGroupID := handlers.GetIntFormValue(r, "inputParentGroupID")
	start := handlers.GetStringFormValue(r, "inputStart")
	end := handlers.GetStringFormValue(r, "inputEnd")
	summary := handlers.GetStringFormValue(r, "inputSummary")
	description := handlers.GetStringFormValue(r, "inputDescription")

	lpr := &model.GroupRequest{
		ID:            normalizedID,
		Name:          name,
		GroupTypeID:   normalizedGroupTypeID,
		ParentGroupID: normalizedParentGroupID,
		Start:         start,
		End:           end,
		Summary:       summary,
		Description:   description,
	}

	return lpr, nil
}

func postGroup(w http.ResponseWriter, r *http.Request, groupRequest *model.GroupRequest) (int, []byte, error) {
	url := getSimpleGroupsURL()

	payload, err := json.Marshal(groupRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchGroup(w http.ResponseWriter, r *http.Request, groupRequest *model.GroupRequest) (int, []byte, error) {
	url := getGroupURL(groupRequest.ID)

	payload, err := json.Marshal(groupRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}
