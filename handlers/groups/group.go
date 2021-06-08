package groups

import (
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type GroupTemplate struct {
	Ct    handlers.CommonTemplate
	Group *model.Group
}

func MakeGroupTemplate(r *http.Request, pageTitle string, group *model.Group) (*GroupTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle)
	if err != nil {
		return nil, err
	}

	groupTemplate := &GroupTemplate{
		Ct:    *ct,
		Group: group,
	}

	return groupTemplate, nil
}

func GetGroup(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-groups", "GetGroup")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	groupID, err := getUrlGroupID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	group, err := getGroup(w, r, groupID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakeGroupTemplate(r, GetLabel(GroupPageTitleIndex), group)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("group", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}

