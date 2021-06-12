package group_types

import (
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

type GroupTypeTemplate struct {
	Ct        handlers.CommonTemplate
	GroupType *model.GroupType
}

func MakeGroupTypeTemplate(r *http.Request, pageTitle, studyTitle string, group *model.GroupType) (*GroupTypeTemplate, error) {
	ct, err := handlers.MakeCommonTemplate(r, pageTitle, studyTitle)
	if err != nil {
		return nil, err
	}

	groupTypeTemplate := &GroupTypeTemplate{
		Ct:        *ct,
		GroupType: group,
	}

	return groupTypeTemplate, nil
}

func GetGroupType(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-group_types", "GetGroupType")

	if handlers.UserAuthenticated(w, r) == false {
		log.FailedReturn()
		handlers.RedirectToSessionExpired(w, r)
		return
	}

	groupID, err := getUrlGroupTypeID(w, r)
	if err != nil {
		log.FailedReturn()
		return
	}

	groupType, err := getGroupType(w, r, groupID)
	if handlers.HandleError(w, r, err) {
		log.FailedReturn()
		return
	}

	ct, err := MakeGroupTypeTemplate(r, GetLabel(GroupTypePageTitleIndex), "", groupType)
	if err != nil {
		log.FailedReturn()
		handlers.RedirectToErrorPage(w, r)
		return
	}

	err = handlers.ExecuteView("group_type", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}

	log.NormalReturn()
}
