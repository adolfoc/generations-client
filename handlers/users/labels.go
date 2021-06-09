package users

import "github.com/adolfoc/generations-client/handlers"

const (
	UserIndexPageTitleIndex            = 0
	UserPageTitleIndex                 = 1
	UserNewPageTitleIndex              = 2
	UserNewSubmitLabelIndex            = 3
	UserEditPageTitleIndex             = 4
	UserEditSubmitLabelIndex           = 5
	UserCreatedIndex                   = 6
	UserUpdatedIndex                   = 7
	UserCreateErrorsReceivedIndex      = 8
	UserUpdateErrorsReceivedIndex      = 9
	UserChangePasswordTitleIndex       = 10
	UserChangePasswordSubmitLabelIndex = 11

	UserIndexPageTitleES            = "Todos los usuarios"
	UserPageTitleES                 = "Usuario"
	UserNewPageTitleES              = "Nuevo Usuario"
	UserNewSubmitLabelES            = "Crear Usuario"
	UserEditPageTitleES             = "Editar Usuario"
	UserEditSubmitLabelES           = "Actualizar Usuario"
	UserCreatedES                   = "El usuario fue creado satisfactoriamente"
	UserUpdatedES                   = "El usuario fue actualizado satisfactoriamente"
	UserCreateErrorsReceivedES      = "Por favor corrija los errores para poder crear el usuario"
	UserUpdateErrorsReceivedES      = "Por favor corrija los errores para poder actualizar el usuario"
	UserChangePasswordTitleES       = "Cambiar contraseña"
	UserChangePasswordSubmitLabelES = "Actualizar"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[UserIndexPageTitleIndex] = UserIndexPageTitleES
		LangES[UserPageTitleIndex] = UserPageTitleES
		LangES[UserNewPageTitleIndex] = UserNewPageTitleES
		LangES[UserNewSubmitLabelIndex] = UserNewSubmitLabelES
		LangES[UserEditPageTitleIndex] = UserEditPageTitleES
		LangES[UserEditSubmitLabelIndex] = UserEditSubmitLabelES
		LangES[UserCreatedIndex] = UserCreatedES
		LangES[UserUpdatedIndex] = UserUpdatedES
		LangES[UserCreateErrorsReceivedIndex] = UserCreateErrorsReceivedES
		LangES[UserUpdateErrorsReceivedIndex] = UserUpdateErrorsReceivedES
		LangES[UserChangePasswordTitleIndex] = UserChangePasswordTitleES
		LangES[UserChangePasswordSubmitLabelIndex] = UserChangePasswordSubmitLabelES
	}
}

func GetLabel(labelIndex int) string {
	initializeMaps()
	locale := handlers.GetLocale()
	if locale == "es" {
		return LangES[labelIndex]
	}

	return ""
}
