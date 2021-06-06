package handlers

import (
	"github.com/adolfoc/generations-client/common"
	"net/http"
)

type TemplateSessionExpired struct {
	Ct         CommonTemplate
}

func MakeSessionExpiredTemplate(r *http.Request, pageTitle string) *TemplateSessionExpired {
	ct := MakeSimpleTemplate(pageTitle)

	expiredTemplate := &TemplateSessionExpired{
		Ct:         *ct,
	}

	return expiredTemplate
}

func SessionExpiredHandler(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-cards", "SessionExpiredHandler")
	ct := MakeSessionExpiredTemplate(r, "Término de sesión")

	err := ExecuteView("session_expired", ct, w)
	if err != nil {
		log.FailedReturn()
		return
	}


	session, _ := GetCurrentSessionFromRequest(r)
	if session != nil {
		session.Destroy()
	}

	c := &http.Cookie{
		Name:   "jwt_cookie",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	log.NormalReturn()
}

