package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"
)

// SessionStatus is used by the NoSession error type to provide additional
// information on why a session was not found.
type SessionStatus int

const (
	NoSessionFound SessionStatus = iota + 1
	NoSchoolSessionCookieFound
	SessionTerminated
)

// NoSession is the type of error we return when we fail to identify a session
type NoSession struct {
	Status SessionStatus
	Message string
}

func (ns *NoSession) Error() string {
	return ns.Message
}

func MakeNoSessionFoundError() *NoSession {
	return &NoSession{
		Status:  NoSessionFound,
		Message: fmt.Sprintf("no current session found"),
	}
}

func MakeNoSchoolSessionCookieFoundError() *NoSession {
	return &NoSession{
		Status:  NoSchoolSessionCookieFound,
		Message: fmt.Sprintf("no api session cookie found"),
	}
}

// SessionData represents our session structure and is used for client processing
type SessionData struct {
	SessionID            string
	GenerationsSessionID string
	Expiry               string
	UserID               int
	UserName             string
	UserEmail            string
	UserRole             string
	ErrorMessage         string
	InfoMessage          string
	GeneralError         *GeneralError
}

var SessionStore []*SessionData

func createSessionIDCookie(w http.ResponseWriter, sessionID string, expiration time.Time) *http.Cookie {
	deleteCookie(w, "session_id")
	cookie := createCookie(w, "session_id", sessionID, expiration)
	return cookie
}

func createSchoolSessionIDCookie(w http.ResponseWriter, schoolSessionID string, expiration time.Time) *http.Cookie {
	deleteCookie(w, "generations_session_id")
	cookie := createCookie(w, "generations_session_id", schoolSessionID, expiration)
	return cookie
}

const (
	SessionCookieDuration = 20		// in minutes
)

func createCookie(w http.ResponseWriter, name, value string, expiration time.Time) *http.Cookie {
	cookie := &http.Cookie{
		Name:       name,
		Value:      value,
		Path:       "/",
		Domain:     "",
		Expires:    expiration,
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	}

	http.SetCookie(w, cookie)
	return cookie
}

func deleteCookie(w http.ResponseWriter, name string) {
	cookie := &http.Cookie{
		Name:   name,
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}

func createSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return base64.URLEncoding.EncodeToString(b)
}

func CreateSession(w http.ResponseWriter, userID int, userRole, schoolSessionID string) error {
	sessionID := createSessionID()
	sd := &SessionData{
		SessionID:            sessionID,
		GenerationsSessionID: schoolSessionID,
		Expiry:               "120",
		UserID:               userID,
		//UserName:     fmt.Sprintf("%s %s", user.FirstNames, user.LastNames),
		//UserEmail:    userEmail,
		UserRole:     userRole,
	}

	SessionStore = append(SessionStore, sd)

	// FIXME: grab expiration from school session id cookie
	expiration := time.Now().Add(SessionCookieDuration * time.Minute)
	createSessionIDCookie(w, sessionID, expiration)
	//createSchoolSessionIDCookie(w, schoolSessionID, expiration)

	return nil
}

func (sd *SessionData) UpdateCookies(w http.ResponseWriter, cookies []*http.Cookie) error {
	var expiration time.Time
	for _, c := range cookies {
		if c.Name == "generations_session_id" {
			expiration = c.Expires
		}
	}

	if expiration.IsZero() {
		expiration = time.Now().Add(SessionCookieDuration * time.Minute)
	}
	createSessionIDCookie(w, sd.SessionID, expiration)

	return nil
}

func UpdateSession(w http.ResponseWriter, r *http.Request, resp *http.Response) error {
	session, err := GetCurrentSessionFromRequest(r)
	if err != nil {
		return err
	}

	cookies := collectCookies(resp)
	if session != nil {
		session.UpdateCookies(w, cookies)
	}

	return nil
}

func getSession(sessionID string) *SessionData {
	for _, session := range SessionStore {
		if session.SessionID == sessionID {
			return session
		}
	}
	return nil
}

func GetCurrentSessionFromRequest(r *http.Request) (*SessionData, error) {
	sessionCookie, err := r.Cookie("session_id")
	if err != nil ||  sessionCookie == nil {
		return nil, MakeNoSessionFoundError()
	}
	fmt.Printf("cookie: %+v\n", sessionCookie)

	sessionData := getSession(sessionCookie.Value)
	if sessionData == nil {
		return nil, MakeNoSessionFoundError()
	}

	return sessionData, nil
}

func (sd *SessionData) GetUserName() (string, error) {
	if sd == nil {
		return "", MakeNoSessionFoundError()
	}

	return sd.UserName, nil
}

func (sd *SessionData) GetUserID() (int, error) {
	if sd == nil {
		return 0, MakeNoSessionFoundError()
	}

	return sd.UserID, nil
}

func (sd *SessionData) GetUserEmail() (string, error) {
	if sd == nil {
		return "", MakeNoSessionFoundError()
	}

	return sd.UserEmail, nil
}

func (sd *SessionData) GetUserRole() (string, error) {
	if sd == nil {
		return "", MakeNoSessionFoundError()
	}

	return sd.UserRole, nil
}

func (sd *SessionData) GetSessionIndex() int {
	for i, d := range SessionStore {
		if d == sd {
			return i
		}
	}

	return -1
}

func (sd *SessionData) Destroy() {
	index := sd.GetSessionIndex()
	if index != -1 {
		SessionStore[index] = SessionStore[len(SessionStore)-1]
		// We do not need to put s[i] at the end, as it will be discarded anyway
		SessionStore = SessionStore[:len(SessionStore)-1]
	}
}

func (sd *SessionData) WriteErrorMessage(message string) {
	sd.ErrorMessage = message
}

func (sd *SessionData) PeekErrorMessage() string {
	return sd.ErrorMessage
}

func (sd *SessionData) UseErrorMessage() string {
	message := sd.ErrorMessage
	sd.ErrorMessage = ""
	return message
}

func (sd *SessionData) WriteInfoMessage(message string) {
	sd.InfoMessage = message
}

func (sd *SessionData) PeekInfoMessage() string {
	return sd.InfoMessage
}

func (sd *SessionData) UseInfoMessage() string {
	message := sd.InfoMessage
	sd.InfoMessage = ""
	return message
}

func (sd *SessionData) WriteSessionErrorMessage(r *http.Request, message string) {
	session, err := GetCurrentSessionFromRequest(r)
	if err == nil {
		session.WriteErrorMessage(message)
	}
}

func (sd *SessionData) WriteSessionInfoMessage(r *http.Request, message string) {
	session, err := GetCurrentSessionFromRequest(r)
	if err == nil {
		session.WriteInfoMessage(message)
	}
}


