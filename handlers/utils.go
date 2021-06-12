package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/adolfoc/generations-client/common"
	"github.com/adolfoc/generations-client/model"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const (
	PrintAPITrace = false
)

// GetAPIHostURL returns a connect string to kickstart the web server.
func GetAPIHostURL() string {
	root, err := os.Getwd()
	fullPath := fmt.Sprintf("%s/.env", root)
	err = godotenv.Load(fullPath)
	if err != nil {
		return ""
	}

	protocol := os.Getenv("GENERATIONS_API_PROTOCOL")
	host := os.Getenv("GENERATIONS_API_HOST")
	port := os.Getenv("GENERATIONS_API_PORT")
	return fmt.Sprintf("%s://%s:%s/", protocol, host, port)
}

func GetClientHostURL() string {
	root, err := os.Getwd()
	fullPath := fmt.Sprintf("%s/.env", root)
	err = godotenv.Load(fullPath)
	if err != nil {
		return ""
	}

	protocol := os.Getenv("GENERATIONS_CLIENT_PROTOCOL")
	host := os.Getenv("GENERATIONS_CLIENT_HOST")
	port := os.Getenv("GENERATIONS_CLIENT_PORT")
	return fmt.Sprintf("%s://%s:%s/", protocol, host, port)
}

// GetLocale will return the current locale, presently it always return the es locale.
func GetLocale() string {
	return "es"
}

func printTrace(code int, body []byte, errs []error) {
	if PrintAPITrace == true {
		fmt.Printf("code %d\n", code)
		fmt.Printf("body %v\n", string(body))
		fmt.Printf("errs %v\n", errs)
	}
}

func CookiesExpired(w http.ResponseWriter, r *http.Request) bool {
	generationsSession, _ := r.Cookie("session_id")
	if generationsSession == nil {
		return true
	}

	return false
}

const (
	ResourceGenerationSchema = "generation-schemas"
)

func GetGenerationSchema(w http.ResponseWriter, r *http.Request, schemaID int) (*model.GenerationSchema, error) {
	url := fmt.Sprintf("%s%s/%d", GetAPIHostURL(), ResourceGenerationSchema, schemaID)
	code, body, err := GetResource(w, r, url)

	if err != nil {
		return nil, err
	}

	if code != 200 {
		return nil, fmt.Errorf("received %d", code)
	}

	var generationSchema *model.GenerationSchema
	err = json.Unmarshal(body, &generationSchema)
	if err != nil {
		return nil, fmt.Errorf("%s", err.Error())
	}

	return generationSchema, nil
}

// GetUrlIntParam is meant to be used to capture a resource id named param in the router.
func GetUrlIntParam(param string, w http.ResponseWriter, r *http.Request) (int, error) {
	vars := mux.Vars(r)
	parameter := vars[param]
	if len(parameter) == 0 {
		fmt.Fprintf(w, "no %q", param)
		return 0, fmt.Errorf("no %q", param)
	}

	intParam, err := strconv.Atoi(parameter)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return 0, err
	}

	return intParam, nil
}

func GetUrlStringParam(param string, w http.ResponseWriter, r *http.Request) (string, error) {
	vars := mux.Vars(r)
	parameter := vars[param]
	if len(parameter) == 0 {
		fmt.Fprintf(w, "no %q", param)
		return "", fmt.Errorf("no %q", param)
	}

	return parameter, nil
}

func GetURLPageParameter(r *http.Request) int {
	v := r.URL.Query()
	fmt.Printf("%+v", v)
	if len(v) == 0 {
		return 1
	}

	page := v.Get("page")
	fmt.Printf("%s", page)
	if page == "" {
		return 1
	}

	pageNum, _ := strconv.Atoi(page)
	return pageNum
}

// isAuthorizationError returns true if the http code refers to an authorization error
func isAuthorizationError(code int) bool {
	return code == 401
}

// shouldHandleError returns true if this http code is an unexpected error
func shouldHandleError(code int) bool {
	if code >= 200 && code <= 300 {
		return false
	}

	return true
}

func GetIntFormValue(r *http.Request, id string) int {
	stringValue := r.FormValue(id)
	intValue, _ := strconv.Atoi(stringValue)
	return intValue
}

func GetStringFormValue(r *http.Request, id string) string {
	stringValue := r.FormValue(id)
	return stringValue
}

func GetBoolFormValue(r *http.Request, id string) bool {
	stringValue := r.FormValue(id)
	if stringValue == "on" {
		return true
	}

	return false
}

func prepareStandardHeader(request *http.Request, clientRequest *http.Request) error {
	session, err := GetCurrentSessionFromRequest(clientRequest)
	if err != nil || session == nil {
		return MakeNoSessionFoundError()
	}

	cookie := fmt.Sprintf("generations_session_id=%s", session.GenerationsSessionID)
	request.Header.Set("Set-Cookie", cookie)

	return nil
}

func prepareGetRequest(clientRequest *http.Request, url string) (*http.Request, error) {
	requestBody := bytes.NewBuffer([]byte(""))
	request, err := http.NewRequest("GET", url, requestBody)
	if err != nil {
		return nil, err
	}

	err = prepareStandardHeader(request, clientRequest)
	if err != nil {
		return nil, err
	}

	return request, nil
}

// GetResource sends an HTTP GET transaction to the api, and returns the api's status code, the body and an optional error.
func GetResource(w http.ResponseWriter, r *http.Request, url string) (int, []byte, error) {
	log := common.StartLog("handlers-utils", "GetResource")
	log.Log(fmt.Sprintf("url: %s", url))

	request, err := prepareGetRequest(r, url)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := performStandardHttpRequest(w, r, request)
	if err != nil {
		log.FailedReturn()
		return code, nil, err
	}

	log.NormalReturn()
	return code, body, nil
}

func GetResourceRaw(w http.ResponseWriter, r *http.Request, url string) (int, []byte, error) {
	log := common.StartLog("handlers-utils", "GetResourceRaw")
	log.Log(fmt.Sprintf("url: %s", url))

	response, err := http.Get(url)
	if err != nil {
		log.FailedReturn()
		return 0, nil, GenerateErrorMessage(r, response.StatusCode, "GetResourceRaw: http.Get", err)
	}
	defer response.Body.Close()

	code := response.StatusCode
	if isAuthorizationError(code) {
		err := MakeAuthenticationError(code)
		log.FailedReturn()
		return code, nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.FailedReturn()
		return 0, nil, GenerateErrorMessage(r, response.StatusCode, "GetResourceRaw: ioutil.ReadAll", err)
	}

	// Cannot call UpdateSession on raw requests on behalf os the system
	//err = UpdateSession(w, r, response)
	//if err != nil {
	//	log.FailedReturn()
	//	return 0, nil, GenerateErrorMessage(r, response.StatusCode, "GetResource: UpdateSession", err)
	//}

	log.NormalReturn()
	return code, body, nil
}

func preparePostRequest(r *http.Request, url string, payload []byte, contentType string) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	err = prepareStandardHeader(request, r)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", contentType)
	return request, nil
}

func PostResource(w http.ResponseWriter, r *http.Request, url string, payload []byte) (int, []byte, error) {
	log := common.StartLog("handlers-utils", "PostResource")
	log.Log(fmt.Sprintf("url: %s", url))

	request, err := preparePostRequest(r, url, payload, "application/json")
	if err != nil {
		return 0, nil, err
	}

	code, body, err := performStandardHttpRequest(w, r, request)
	if err != nil {
		log.FailedReturn()
		return code, nil, err
	}

	log.NormalReturn()
	return code, body, nil
}

func PostMultipartResource(w http.ResponseWriter, r *http.Request, url string, payload *bytes.Buffer, contentTypeValue string) (int, []byte, error) {
	log := common.StartLog("handlers-utils", "PostMultipartResource")
	log.Log(fmt.Sprintf("url: %s", url))

	request, err := preparePostRequest(r, url, payload.Bytes(), contentTypeValue)
	if err != nil {
		return 0, nil, err
	}

	//request, err := http.NewRequest(http.MethodPost, url, payload)
	//err = prepareStandardHeader(request, r)
	//if err != nil {
	//	return 0, nil, err
	//}
	//request.Header.Set("Content-Type", contentTypeValue)

	code, body, err := performStandardHttpRequest(w, r, request)
	if err != nil {
		log.FailedReturn()
		return code, nil, err
	}

	log.NormalReturn()
	return code, body, nil
}

func PostResourceRaw(r *http.Request, url string, payload []byte) (int, []byte, []*http.Cookie, error) {
	log := common.StartLog("handlers-utils", "PostResourceRaw")
	log.Log(fmt.Sprintf("url: %s", url))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.FailedReturn()
		return 0, nil, nil, err
	}
	defer resp.Body.Close()

	code := resp.StatusCode
	if isAuthorizationError(code) {
		err := MakeAuthenticationError(code)
		log.FailedReturn()
		return code, nil, nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.FailedReturn()
		return 0, nil, nil, GenerateErrorMessage(r, resp.StatusCode, "PostResourceRaw: ioutil.ReadAll", err)
	}

	cookies := collectCookies(resp)

	log.NormalReturn()
	return resp.StatusCode, body, cookies, nil
}

func preparePatchRequest(r *http.Request, url string, payload []byte) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	err = prepareStandardHeader(request, r)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	return request, nil
}

func PatchResource(w http.ResponseWriter, r *http.Request, url string, payload []byte) (int, []byte, error) {
	log := common.StartLog("handlers-utils", "PatchResource")
	log.Log(fmt.Sprintf("url: %s", url))

	request, err := preparePatchRequest(r, url, payload)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := performStandardHttpRequest(w, r, request)
	if err != nil {
		log.FailedReturn()
		return code, nil, err
	}

	log.NormalReturn()
	return code, body, nil
}

func PatchMultipartResource(w http.ResponseWriter, r *http.Request, url string, payload *bytes.Buffer, contentTypeValue string) (int, []byte, error) {
	log := common.StartLog("handlers-utils", "PatchMultipartResource")
	log.Log(fmt.Sprintf("url: %s", url))

	request, err := http.NewRequest(http.MethodPatch, url, payload)
	err = prepareStandardHeader(request, r)
	if err != nil {
		return 0, nil, err
	}
	request.Header.Set("Content-Type", contentTypeValue)

	code, body, err := performStandardHttpRequest(w, r, request)
	if err != nil {
		log.FailedReturn()
		return 0, nil, err
	}

	log.NormalReturn()
	return code, body, nil
}

func preparePutRequest(r *http.Request, url string, payload []byte) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	err = prepareStandardHeader(request, r)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	return request, nil
}

func PutResource(w http.ResponseWriter, r *http.Request, url string, payload []byte) (int, []byte, error) {
	log := common.StartLog("handlers-utils", "PutResource")
	log.Log(fmt.Sprintf("url: %s", url))

	request, err := preparePutRequest(r, url, payload)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := performStandardHttpRequest(w, r, request)
	if err != nil {
		log.FailedReturn()
		return code, nil, err
	}

	log.NormalReturn()
	return code, body, nil
}

func prepareDeleteRequest(r *http.Request, url string) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	err = prepareStandardHeader(request, r)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	return request, nil
}

func DeleteResource(w http.ResponseWriter, r *http.Request, url string) (int, []byte, error) {
	log := common.StartLog("handlers-utils", "DeleteResource")
	log.Log(fmt.Sprintf("url: %s", url))

	request, err := prepareDeleteRequest(r, url)
	if err != nil {
		return 0, nil, err
	}

	code, body, err := performStandardHttpRequest(w, r, request)
	if err != nil {
		log.FailedReturn()
		return code, nil, err
	}

	log.NormalReturn()
	return code, body, nil
}

// performStandardHttpRequest performs an HTTP request using our standard client and performs a series of checks
// once the transaction returns: 1. checks for general errors (such as connectivity); 2. checks for authorization
// error codes; 3. reads the response body and reports any errors; 4. updates our session with the cookies received.
func performStandardHttpRequest(w http.ResponseWriter, r *http.Request, request *http.Request) (int, []byte, error) {
	log := common.StartLog("handlers-utils", "performStandardHttpRequest")

	client := GetHttpClient()
	resp, err := client.Do(request)
	if err != nil {
		statusCode := 0
		if resp != nil {
			statusCode = resp.StatusCode
		}
		log.FailedReturn()
		return 0, nil, GenerateErrorMessage(r, statusCode, "PatchMultipartResource: client.Do", err)
	}
	defer resp.Body.Close()

	code := resp.StatusCode
	if isAuthorizationError(code) {
		err := MakeAuthenticationError(code)
		log.FailedReturn()
		return code, nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.FailedReturn()
		return 0, nil, GenerateErrorMessage(r, resp.StatusCode, "PatchMultipartResource: ioutil.ReadAll", err)
	}

	printTrace(code, body, nil)

	err = UpdateSession(w, r, resp)
	if err != nil {
		log.FailedReturn()
		return 0, nil, GenerateErrorMessage(r, resp.StatusCode, "GetResource: UpdateSession", err)
	}

	log.NormalReturn()
	return code, body, nil
}

func GenerateErrorMessage(r *http.Request, httpCode int, doingWhat string, err error) *GeneralError {
	log := common.StartLog("handlers-utils", "GenerateErrorMessage")

	genError := MakeGeneralError(httpCode, doingWhat, err)

	session, _ := GetCurrentSessionFromRequest(r)
	session.GeneralError = genError

	log.NormalReturn()
	return genError
}

func GenerateNoSessionErrorMessage(r *http.Request, httpCode int, doingWhat string, err error) *GeneralError {
	log := common.StartLog("handlers-utils", "GenerateNoSessionErrorMessage")

	genError := MakeGeneralError(httpCode, doingWhat, err)

	session, _ := GetCurrentSessionFromRequest(r)
	session.GeneralError = genError

	log.NormalReturn()
	return genError
}

func OnUpdateError(r *http.Request, body []byte, message string) (*ResponseErrors, error) {
	log := common.StartLog("handlers-utils", "OnUpdateError")

	session, err := GetCurrentSessionFromRequest(r)
	if err != nil {
		log.FailedReturn()
		return nil, err
	}
	session.WriteSessionErrorMessage(r, message)

	var responseErrors ResponseErrors
	err = json.Unmarshal(body, &responseErrors)
	if err != nil {
		log.FailedReturn()
		return nil, err
	}

	log.NormalReturn()
	return &responseErrors, err
}

func OnDeleteError(r *http.Request, body []byte) (*ResponseErrors, error) {
	log := common.StartLog("handlers-utils", "OnDeleteError")

	var responseErrors ResponseErrors
	err := json.Unmarshal(body, &responseErrors)
	if err != nil {
		log.FailedReturn()
		return nil, err
	}

	log.NormalReturn()
	return &responseErrors, err
}

func ExtractFirstError(errors *ResponseErrors) string {
	if len(errors.Errors) > 0 {
		thisError := errors.Errors[0]
		return thisError.Detail
	}

	return ""
}

func WriteSessionInfoMessage(r *http.Request, message string) {
	currentSession, _ := GetCurrentSessionFromRequest(r)
	currentSession.WriteSessionInfoMessage(r, message)
}

func WriteSessionErrorMessage(r *http.Request, message string) {
	currentSession, _ := GetCurrentSessionFromRequest(r)
	currentSession.WriteSessionErrorMessage(r, message)
}

// collectCookies is invoked after each api call to collect the cookies that the api server sends.
// At this point we only care about the jwt_cookie.
func collectCookies(response *http.Response) []*http.Cookie {
	cookies, _ := GrabAllCookies(response)

	return cookies
}

func AfterRequestHandler(w http.ResponseWriter, r *http.Request, cookies []*http.Cookie) {
	log := common.StartLog("handlers-utils", "AfterRequestHandler")
	session, err := GetCurrentSessionFromRequest(r)
	if err != nil {
		log.FailedReturn()
		return
	}

	if session != nil {
		session.UpdateCookies(w, cookies)
	}

	log.NormalReturn()
}

func UserAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	//cookie := getCookieJarJwtCookie()
	//if cookie == nil {
	//	return false
	//}

	if CookiesExpired(w, r) {
		return false
	}

	//session, err := GetCurrentSession()
	//if err != nil {
	//	return false
	//}
	//
	//cookies := collectCookies(r.Response)
	//if session != nil {
	//	session.UpdateCookies(w, cookies)
	//}

	return true
}

const AnyRole = "Any"

//func UserAuthenticatedAs(w http.ResponseWriter, r *http.Request, role string) bool {
//	if role == AnyRole {
//		return true
//	}
//
//	userRole := c.Cookies("user_role")
//	if userRole != role {
//		return false
//	}
//
//	return true
//}

func RedirectToLogin(w http.ResponseWriter, r *http.Request) error {
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	return nil
}

func RedirectToSessionExpired(w http.ResponseWriter, r *http.Request) error {
	http.Redirect(w, r, "/session-expired", http.StatusMovedPermanently)
	return nil
}

func RedirectToErrorPage(w http.ResponseWriter, r *http.Request) error {
	http.Redirect(w, r, "/general-error", http.StatusMovedPermanently)
	return nil
}

