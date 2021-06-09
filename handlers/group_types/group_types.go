package group_types

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type GroupTypesTemplate struct {
	Ct         handlers.CommonTemplate
	GroupTypes *model.GroupTypes
	Pagination *handlers.Pagination
}

func getPaginationBaseURL() string {
	stem := fmt.Sprintf("/groups-types/index")
	return stem + "?page=%d"
}

func MakeGroupTypesTemplate(r *http.Request, pageTitle string, page int, groupTypes *model.GroupTypes) (*GroupTypesTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle)
	if err != nil {
		return nil, err
	}

	pagination := handlers.MakePagination(groupTypes.RecordCount, getPaginationBaseURL(), page)

	groupsTypesTemplate := &GroupTypesTemplate{
		Ct:         *ct,
		GroupTypes: groupTypes,
		Pagination: pagination,
	}

	return groupsTypesTemplate, nil
}

func GetGroupTypes(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-group_types", "GetGroupTypes")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	page := handlers.GetURLPageParameter(r)
	events, err := getGroupTypes(w, r, page)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakeGroupTypesTemplate(r, GetLabel(GroupTypeIndexPageTitleIndex), page, events)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("group_types", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}


