package model

import (
	"fmt"
	"html/template"
	"strings"
)

type User struct {
	ID         int    `json:"id"`
	UserName   string `json:"user_name"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	FirstNames string `json:"first_names"`
	LastNames  string `json:"last_names"`
	IsActive   bool   `json:"is_active"`
}

type Users struct {
	Users     []*User     `json:"users"`
	RecordCount int       `json:"record_count"`
}

type LoginRequest struct {
	Email string				`json:"email"`
	Password string				`json:"password"`
}

type NewUserRequest struct {
	ID                   int    `json:"id"`
	UserName             string `json:"user_name"`
	Role                 string `json:"role"`
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	FirstNames           string `json:"first_names"`
	LastNames            string `json:"last_names"`
	IsActive             bool   `json:"is_active"`
}

type UpdateUserRequest struct {
	ID         int    `json:"id"`
	UserName   string `json:"user_name"`
	Role       string `json:"role"`
	Email      string `json:"email"`
	FirstNames string `json:"first_names"`
	LastNames  string `json:"last_names"`
	IsActive   bool   `json:"is_active"`
}

type ChangePasswordRequest struct {
	ID                      int    `json:"id"`
	OldPassword             string `json:"old_password"`
	NewPassword             string `json:"new_password"`
	NewPasswordConfirmation string `json:"new_password_confirmation"`
}

func (u *User) IDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (u *User) UserNameLabel() template.HTML {
	return BuildLabel("inputUserName", "Nombre de usuario")
}

func (u *User) UserNameInput(message, value string) template.HTML {
	return BuildTextInput("inputUserName", value, message)
}

func (u *User) RoleLabel() template.HTML {
	return BuildLabel("inputRole", "Rol")
}

func (u *User) RoleSelectBox(message, selectedRole string) template.HTML {
	var selectBox []string
	startSelect := fmt.Sprintf("<select class='form-select library-control' id='inputRole' name='inputRole'>")
	selectBox = append(selectBox, startSelect)

	roles := []string{"lector", "editor", "admin"}
	for _, role := range roles {
		var option string
		if role == selectedRole {
			option = fmt.Sprintf("<option value='%s' selected>%s</option>", role, role)
		} else {
			option = fmt.Sprintf("<option value='%s'>%s</option>", role, role)
		}
		selectBox = append(selectBox, option)
	}

	endSelect := fmt.Sprintf("</select>")
	selectBox = append(selectBox, endSelect)

	return template.HTML(strings.Join(selectBox, "\n"))
}

func (u *User) EmailLabel() template.HTML {
	return BuildLabel("inputEmail", "Email")
}

func (u *User) EmailInput(message, value string) template.HTML {
	return BuildEmailInput("inputEmail", value, message)
}

func (u *User) FirstNamesLabel() template.HTML {
	return BuildLabel("inputFirstNames", "Nombres")
}

func (u *User) FirstNamesInput(message, value string) template.HTML {
	return BuildTextInput("inputFirstNames", value, message)
}

func (u *User) LastNamesLabel() template.HTML {
	return BuildLabel("inputLastNames", "Apellidos")
}

func (u *User) LastNamesInput(message, value string) template.HTML {
	return BuildTextInput("inputLastNames", value, message)
}

func (u *User) IsActiveInput(message string, value bool) template.HTML {
	return BuildCheckboxInput("inputIsActive", "Activo", value, message)
}

func (u *User) PasswordLabel() template.HTML {
	return BuildLabel("inputPassword", "Contrase??a")
}

func (u *User) PasswordInput(message, value string) template.HTML {
	return BuildPasswordInput("inputPassword", value, message)
}

func (u *User) PasswordConfirmationLabel() template.HTML {
	return BuildLabel("inputPasswordConfirmation", "Confirme Contrase??a")
}

func (u *User) PasswordConfirmationInput(message, value string) template.HTML {
	return BuildPasswordInput("inputPasswordConfirmation", value, message)
}


func (u *User) OldPasswordLabel() template.HTML {
	return BuildLabel("inputOldPassword", "Contrase??a actual")
}

func (u *User) OldPasswordInput(message, value string) template.HTML {
	return BuildPasswordInput("inputOldPassword", value, message)
}

func (u *User) NewPasswordInput(message, value string) template.HTML {
	return BuildPasswordInput("inputNewPassword", value, message)
}

func (u *User) NewPasswordLabel() template.HTML {
	return BuildLabel("inputNewPassword", "Contrase??a nueva")
}

func (u *User) NewPasswordConfirmationLabel() template.HTML {
	return BuildLabel("inputNewPasswordConfirmation", "Confirme Contrase??a")
}

func (u *User) NewPasswordConfirmationInput(message, value string) template.HTML {
	return BuildPasswordInput("inputNewPasswordConfirmation", value, message)
}
