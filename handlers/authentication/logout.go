package authentication

import (
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-authentication", "Logout")

	session, _ := handlers.GetCurrentSessionFromRequest(r)
	if session != nil {
		// FIXME: Notify the api
		session.Destroy()
	}

	deleteCookies(w, []string{ "session_id" })

	log.NormalReturn()
	handlers.RedirectToLogin(w, r)
}

func deleteCookies(w http.ResponseWriter, cookieNames []string) {
	for _, cn := range cookieNames {
		c := &http.Cookie{
			Name:   cn,
			MaxAge: -1,
		}
		http.SetCookie(w, c)
	}
}

