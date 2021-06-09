package group_types

import "github.com/adolfoc/generations-client/handlers"

const (
	GroupTypeIndexPageTitleIndex       = 0
	GroupTypePageTitleIndex            = 1
	GroupTypeNewPageTitleIndex         = 2
	GroupTypeNewSubmitLabelIndex       = 3
	GroupTypeEditPageTitleIndex        = 4
	GroupTypeEditSubmitLabelIndex      = 5
	GroupTypeCreatedIndex              = 6
	GroupTypeUpdatedIndex              = 7
	GroupTypeCreateErrorsReceivedIndex = 8
	GroupTypeUpdateErrorsReceivedIndex = 9

	GroupTypeIndexPageTitleES       = "Todos los tipos de grupo"
	GroupTypePageTitleES            = "Tipo de Grupo"
	GroupTypeNewPageTitleES         = "Nuevo Tipo de Grupo"
	GroupTypeNewSubmitLabelES       = "Crear Tipo de Grupo"
	GroupTypeEditPageTitleES        = "Editar Tipo de Grupo"
	GroupTypeEditSubmitLabelES      = "Actualizar Tipo de Grupo"
	GroupTypeCreatedES              = "El tipo de grupo fue creado satisfactoriamente"
	GroupTypeUpdatedES              = "El tipo de grupo fue actualizado satisfactoriamente"
	GroupTypeCreateErrorsReceivedES = "Por favor corrija los errores para poder crear el tipo de grupo"
	GroupTypeUpdateErrorsReceivedES = "Por favor corrija los errores para poder actualizar el tipo de grupo"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[GroupTypeIndexPageTitleIndex] = GroupTypeIndexPageTitleES
		LangES[GroupTypePageTitleIndex] = GroupTypePageTitleES
		LangES[GroupTypeNewPageTitleIndex] = GroupTypeNewPageTitleES
		LangES[GroupTypeNewSubmitLabelIndex] = GroupTypeNewSubmitLabelES
		LangES[GroupTypeEditPageTitleIndex] = GroupTypeEditPageTitleES
		LangES[GroupTypeEditSubmitLabelIndex] = GroupTypeEditSubmitLabelES
		LangES[GroupTypeCreatedIndex] = GroupTypeCreatedES
		LangES[GroupTypeUpdatedIndex] = GroupTypeUpdatedES
		LangES[GroupTypeCreateErrorsReceivedIndex] = GroupTypeCreateErrorsReceivedES
		LangES[GroupTypeUpdateErrorsReceivedIndex] = GroupTypeUpdateErrorsReceivedES
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


