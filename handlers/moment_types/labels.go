package moment_types

import "github.com/adolfoc/generations-client/handlers"

const (
	MomentTypeIndexPageTitleIndex       = 0
	MomentTypePageTitleIndex            = 1
	MomentTypeNewPageTitleIndex         = 2
	MomentTypeNewSubmitLabelIndex       = 3
	MomentTypeEditPageTitleIndex        = 4
	MomentTypeEditSubmitLabelIndex      = 5
	MomentTypeCreatedIndex              = 6
	MomentTypeUpdatedIndex              = 7
	MomentTypeCreateErrorsReceivedIndex = 8
	MomentTypeUpdateErrorsReceivedIndex = 9

	MomentTypeIndexPageTitleES       = "Todos los tipos de momento"
	MomentTypePageTitleES            = "Tipo de Momento"
	MomentTypeNewPageTitleES         = "Nuevo Tipo de Momento"
	MomentTypeNewSubmitLabelES       = "Crear Tipo de Momento"
	MomentTypeEditPageTitleES        = "Editar Tipo de Momento"
	MomentTypeEditSubmitLabelES      = "Actualizar Tipo de Momento"
	MomentTypeCreatedES              = "El tipo de momento fue creado satisfactoriamente"
	MomentTypeUpdatedES              = "El tipo de momento fue actualizado satisfactoriamente"
	MomentTypeCreateErrorsReceivedES = "Por favor corrija los errores para poder crear el tipo de momento"
	MomentTypeUpdateErrorsReceivedES = "Por favor corrija los errores para poder actualizar el tipo momento"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[MomentTypeIndexPageTitleIndex] = MomentTypeIndexPageTitleES
		LangES[MomentTypePageTitleIndex] = MomentTypePageTitleES
		LangES[MomentTypeNewPageTitleIndex] = MomentTypeNewPageTitleES
		LangES[MomentTypeNewSubmitLabelIndex] = MomentTypeNewSubmitLabelES
		LangES[MomentTypeEditPageTitleIndex] = MomentTypeEditPageTitleES
		LangES[MomentTypeEditSubmitLabelIndex] = MomentTypeEditSubmitLabelES
		LangES[MomentTypeCreatedIndex] = MomentTypeCreatedES
		LangES[MomentTypeUpdatedIndex] = MomentTypeUpdatedES
		LangES[MomentTypeCreateErrorsReceivedIndex] = MomentTypeCreateErrorsReceivedES
		LangES[MomentTypeUpdateErrorsReceivedIndex] = MomentTypeUpdateErrorsReceivedES
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

