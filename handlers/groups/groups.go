package groups

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type GroupsTemplate struct {
	Ct         handlers.CommonTemplate
	Groups     *model.Groups
	Pagination *handlers.Pagination
}

func getPaginationBaseURL() string {
	stem := fmt.Sprintf("/groups/index")
	return stem + "?page=%d"
}

func MakeGroupsTemplate(r *http.Request, pageTitle string, page int, groups *model.Groups) (*GroupsTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle)
	if err != nil {
		return nil, err
	}

	pagination := handlers.MakePagination(groups.RecordCount, getPaginationBaseURL(), page)

	groupsTemplate := &GroupsTemplate{
		Ct:         *ct,
		Groups:     groups,
		Pagination: pagination,
	}

	return groupsTemplate, nil
}

func GetGroups(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-groups", "GetGroups")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	page := handlers.GetURLPageParameter(r)
	events, err := getGroups(w, r, page)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakeGroupsTemplate(r, GetLabel(GroupIndexPageTitleIndex), page, events)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("groups", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

