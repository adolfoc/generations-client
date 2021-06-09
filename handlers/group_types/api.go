package group_types


import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

const (
	ResourceGroupType = "group-types"
)

func getGroupTypesURL(page int) string {
	return fmt.Sprintf("%s%s?page[number]=%d", handlers.GetAPIHostURL(), ResourceGroupType, page)
}

func getSimpleGroupTypesURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), ResourceGroupType)
}

func getGroupTypeURL(groupTypeID int) string {
	return fmt.Sprintf("%s%s/%d", handlers.GetAPIHostURL(), ResourceGroupType, groupTypeID)
}

func getGroupTypes(w http.ResponseWriter, r *http.Request, page int) (*model.GroupTypes, error) {
	url := getGroupTypesURL(page)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var groupTypes *model.GroupTypes
	err = json.Unmarshal(body, &groupTypes)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return groupTypes, nil
}

func getGroupType(w http.ResponseWriter, r *http.Request, groupID int) (*model.GroupType, error) {
	url := getGroupTypeURL(groupID)
	code, body, err := handlers.GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var groupType *model.GroupType
	err = json.Unmarshal(body, &groupType)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return groupType, nil
}

func getUrlGroupTypeID(w http.ResponseWriter, r *http.Request) (int, error) {
	return handlers.GetUrlIntParam("group_type_id", w, r)
}

func buildGroupTypeRequest(groupType *model.GroupType) *model.GroupTypeRequest {
	mtr := &model.GroupTypeRequest{
		ID:          groupType.ID,
		Name:        groupType.Name,
		Description: groupType.Description,
	}

	return mtr
}

func makeGroupTypeRequest(r *http.Request) (*model.GroupTypeRequest, error) {
	normalizedID := handlers.GetIntFormValue(r, "inputID")
	name := handlers.GetStringFormValue(r, "inputName")
	description := handlers.GetStringFormValue(r, "inputDescription")

	lpr := &model.GroupTypeRequest{
		ID:            normalizedID,
		Name:          name,
		Description:   description,
	}

	return lpr, nil
}

func postGroupType(w http.ResponseWriter, r *http.Request, groupTypeRequest *model.GroupTypeRequest) (int, []byte, error) {
	url := getSimpleGroupTypesURL()

	payload, err := json.Marshal(groupTypeRequest)
	if err != nil {
		return 0, nil, err
	}

	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PostResource(w, r, url, payload)

	return code, body, nil
}

func patchGroupType(w http.ResponseWriter, r *http.Request, groupTypeRequest *model.GroupTypeRequest) (int, []byte, error) {
	url := getGroupTypeURL(groupTypeRequest.ID)

	payload, err := json.Marshal(groupTypeRequest)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := handlers.PatchResource(w, r, url, payload)
	if err != nil {
		return 0, nil, err
	}


	return code, body, nil
}
