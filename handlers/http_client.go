package handlers

import (
	"github.com/adolfoc/generations-client/common"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

var (
	httpClient *http.Client
)

// SetupHttpClient is called after successful authentication to create an http client with it's cookie jar
// so we can identify ourselves to the api.
func SetupHttpClient(cookies []*http.Cookie) *http.Client {
	// Set cookiejar options
	//options := cookiejar.Options{
	//	PublicSuffixList: publicsuffix.List,
	//}

	// Create new cookiejar for holding cookies
	jar, _ := cookiejar.New(nil)
	url := MakeClientStandardURL()
	var emptyCollection []*http.Cookie
	jar.SetCookies(url, emptyCollection)

	// Create new http client with predefined options
	httpClient = &http.Client{
		Jar:     jar,
		Timeout: time.Second * 60,
	}

	return 	httpClient
}

func GetHttpClient() *http.Client {
	log := common.StartLog("utils", "GetHttpClient")

	if PrintAPITrace == true {
		//url := MakeClientStandardURL()
		//for _, cookie := range httpClient.Jar.Cookies(url) {
		//	message := fmt.Sprintf("name: %s, value %s", cookie.Name, cookie.Value)
		//	log.Log(message)
		//}
	}

	log.NormalReturn()
	return httpClient
}

func GrabAllCookies(response *http.Response) ([]*http.Cookie, error) {
	var cookies []*http.Cookie
	hCookies := response.Header.Values("Set-Cookie")
	for _, cookie := range hCookies {
		cookieType, cookieValue := detectCookie(cookie)
		if cookieType == "school_session_id" {
			cookie, _ := GrabCookie("school_session_id", cookieValue)
			cookies = append(cookies, cookie)
		} else if cookieType == "session_id" {
			cookie, _ := GrabCookie("session_id", cookieValue)
			cookies = append(cookies, cookie)
		}
	}

	return cookies, nil
}

func GrabCookie(name, value string) (*http.Cookie, error) {
	cookie := &http.Cookie{
		Name:       name,
		Value:      value,
		Path:       "",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	}

	return cookie, nil
}

func detectCookie(cookie string) (string, string) {
	terms := strings.Split(cookie, "=")
	if len(terms) != 2 {
		return "", ""
	}

	return terms[0], terms[1]
}

func getCookieJarSchoolSessionIDCookie() *http.Cookie {
	url := MakeClientStandardURL()
	httpClient := GetHttpClient()
	cookies := httpClient.Jar.Cookies(url)
	for _, c := range cookies {
		if c.Name == "school_session_id" {
			return c
		}
	}

	return nil
}

func MakeClientStandardURL() *url.URL {
	url := &url.URL{
		Scheme:     "http",
		Opaque:     "",
		User:       nil,
		Host:       "localhost",
		Path:       "/",
		RawPath:    "",
		ForceQuery: false,
		RawQuery:   "",
		Fragment:   "",
	}

	return url
}
