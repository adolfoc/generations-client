package generation_types

import "github.com/adolfoc/generations-client/handlers"

const (
	GenerationTypeIndexPageTitleIndex       = 0
	GenerationTypePageTitleIndex            = 1
	GenerationTypeNewPageTitleIndex         = 2
	GenerationTypeNewSubmitLabelIndex       = 3
	GenerationTypeEditPageTitleIndex        = 4
	GenerationTypeEditSubmitLabelIndex      = 5
	GenerationTypeCreatedIndex              = 6
	GenerationTypeUpdatedIndex              = 7
	GenerationTypeCreateErrorsReceivedIndex = 8
	GenerationTypeUpdateErrorsReceivedIndex = 9
	GenerationTypeDeletedIndex              = 10

	GenerationTypeIndexPageTitleES       = "Todos los tipos de generación"
	GenerationTypePageTitleES            = "Tipo de Generación"
	GenerationTypeNewPageTitleES         = "Nuevo Tipo de Generación"
	GenerationTypeNewSubmitLabelES       = "Crear Tipo de Generación"
	GenerationTypeEditPageTitleES        = "Editar Tipo de Generación"
	GenerationTypeEditSubmitLabelES      = "Actualizar Tipo de Generación"
	GenerationTypeCreatedES              = "El tipo de generación fue creada satisfactoriamente"
	GenerationTypeUpdatedES              = "El tipo de generación fue actualizada satisfactoriamente"
	GenerationTypeCreateErrorsReceivedES = "Por favor corrija los errores para poder crear el tipo de generación"
	GenerationTypeUpdateErrorsReceivedES = "Por favor corrija los errores para poder actualizar el tipo generación"
	GenerationTypeDeletedES              = "El tipo de generación fue eliminado satisfactoriamente"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[GenerationTypeIndexPageTitleIndex] = GenerationTypeIndexPageTitleES
		LangES[GenerationTypePageTitleIndex] = GenerationTypePageTitleES
		LangES[GenerationTypeNewPageTitleIndex] = GenerationTypeNewPageTitleES
		LangES[GenerationTypeNewSubmitLabelIndex] = GenerationTypeNewSubmitLabelES
		LangES[GenerationTypeEditPageTitleIndex] = GenerationTypeEditPageTitleES
		LangES[GenerationTypeEditSubmitLabelIndex] = GenerationTypeEditSubmitLabelES
		LangES[GenerationTypeCreatedIndex] = GenerationTypeCreatedES
		LangES[GenerationTypeUpdatedIndex] = GenerationTypeUpdatedES
		LangES[GenerationTypeCreateErrorsReceivedIndex] = GenerationTypeCreateErrorsReceivedES
		LangES[GenerationTypeUpdateErrorsReceivedIndex] = GenerationTypeUpdateErrorsReceivedES
		LangES[GenerationTypeDeletedIndex] = GenerationTypeDeletedES
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
