package authentication

import (
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/handlers"
	"github.com/adolfoc/generations-client/model"
	"net/http"
)

func getAuthenticateURL() string {
	return fmt.Sprintf("%s%s", handlers.GetAPIHostURL(), "auth/getToken")
}

func authenticate(r *http.Request, email, password string) (*model.AuthenticationResponse, []*http.Cookie, error) {
	log := common.StartLog("handlers-authentication", "authenticate")

	url := getAuthenticateURL()
	loginRequest := model.LoginRequest{
		Email:    email,
		Password: password,
	}

	payload, err := json.Marshal(loginRequest)
	code, body, cookies, err := handlers.PostResourceRaw(r, url, payload)

	if code != 202 {
		log.FailedReturn()
		return nil, nil, fmt.Errorf("received %d", code)
	}

	if err != nil {
		log.FailedReturn()
		return nil, nil, err
	}

	var authenticationResponse model.AuthenticationResponse
	err = json.Unmarshal(body, &authenticationResponse)
	if err != nil {
		log.FailedReturn()
		return nil, nil, fmt.Errorf("%s", err.Error())
	}

	log.NormalReturn()
	return &authenticationResponse, cookies, nil
}


