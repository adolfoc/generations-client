package authentication

import (
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func grabGenerationsSessionIDCookie(cookies []*http.Cookie) *http.Cookie {
	for _, cookie := range cookies {
		if cookie.Name == "generations_session_id" {
			return cookie
		}
	}

	return nil
}

func PerformAuthentication(w http.ResponseWriter, r *http.Request) {
	log := common.StartLog("handlers-authentication", "PerformAuthentication")

	email := r.FormValue("inputEmail")
	password := r.FormValue("inputPassword")

	authenticationResponse, cookies, err := authenticate(r, email, password)
	if err != nil {
		if err.Error() == "received 401" {
			ar := &model.AuthenticationRequest{
				Email:    email,
				Password: password,
			}
			AuthenticateRetry(w, ar, "No se encontraron esas credenciales")
			log.FailedReturn()
			return
		}
		log.FailedReturn()
		return
	}

	schoolSessionIDCookie := grabGenerationsSessionIDCookie(cookies)
	if schoolSessionIDCookie != nil {
		//http.SetCookie(w, schoolSessionIDCookie)
	}

	handlers.CreateSession(w, authenticationResponse.UserID, authenticationResponse.Role, authenticationResponse.GenerationsSessionID)

	sessionIDCookie, _ := r.Cookie("session_id")
	if sessionIDCookie != nil {
		cookies = append(cookies, sessionIDCookie)
	}

	client := handlers.SetupHttpClient(cookies)
	fmt.Printf("client: %+v\n", client)

	log.NormalReturn()
	http.Redirect(w, r, "/schemas/index", http.StatusMovedPermanently)
}

