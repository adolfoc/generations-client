package handlers

import "net/http"

const (
	BarTitle = "Estudio Generacional"
	AppTitle= "Estudio Generacional"
	ParkName= "Instituto Tokarev"
)

type CommonTemplate struct {
	InfoMessage  string
	ErrorMessage string
	Title        string
	ParkName     string
	PageTitle    string
	BarTitle     string
	UserName     string
	UserEmail    string
	UserRole     string
}

func MakeCommonTemplate(r *http.Request, pageTitle string) (*CommonTemplate, error) {
	currentSession, err := GetCurrentSessionFromRequest(r)
	if err != nil {
		return nil, err
	}
	userName, _ := currentSession.GetUserName()
	userEmail, _ := currentSession.GetUserEmail()
	userRole, _ := currentSession.GetUserRole()

	ct := &CommonTemplate{
		InfoMessage:   currentSession.UseInfoMessage(),
		ErrorMessage:   currentSession.UseErrorMessage(),
		BarTitle:  BarTitle,
		Title:     AppTitle,
		ParkName:  ParkName,
		PageTitle: pageTitle,
		UserName:  userName,
		UserEmail: userEmail,
		UserRole:  userRole,
	}

	return ct, nil
}

func MakeSimpleTemplate(pageTitle string) *CommonTemplate {
	ct := &CommonTemplate{
		BarTitle:  BarTitle,
		Title:     AppTitle,
		ParkName:  ParkName,
		PageTitle: pageTitle,
	}

	return ct
}

func (ct *CommonTemplate) SetErrorMessage(message string) {
	ct.ErrorMessage = message
}

func (ct *CommonTemplate) SetInfoMessage(message string) {
	ct.ErrorMessage = message
}

