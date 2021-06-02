package model

import "html/template"

// AuthenticationRequest must be kept in sync with api
type AuthenticationRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthenticationResponse must be kept in sync with api
type AuthenticationResponse struct {
	GenerationsSessionID string `json:"generations_session_id"`
	UserID               int    `json:"user_id"`
	Role                 string `json:"role"`
}

func (ar *AuthenticationRequest) EmailLabel() template.HTML {
	return BuildLabel("inputEmail", "Email")
}

func (ar *AuthenticationRequest) EmailInput(message, value string) template.HTML {
	return BuildEmailInput("inputEmail", value, message)
}

func (ar *AuthenticationRequest) PasswordLabel() template.HTML {
	return BuildLabel("inputPassword", "Contraseña")
}

func (ar *AuthenticationRequest) PasswordInput(message, value string) template.HTML {
	return BuildPasswordInput("inputPassword", value, message)
}

