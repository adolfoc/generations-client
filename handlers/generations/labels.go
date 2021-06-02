package generations

import "github.com/adolfoc/generations-client/handlers"

const (
	GenerationIndexPageTitleIndex       = 0
	GenerationPageTitleIndex            = 1
	GenerationNewPageTitleIndex         = 2
	GenerationNewSubmitLabelIndex       = 3
	GenerationEditPageTitleIndex        = 4
	GenerationEditSubmitLabelIndex      = 5
	GenerationCreatedIndex              = 6
	GenerationUpdatedIndex              = 7
	GenerationCreateErrorsReceivedIndex = 8
	GenerationUpdateErrorsReceivedIndex = 9

	GenerationIndexPageTitleES       = "Todas las generaciones"
	GenerationPageTitleES            = "Generación"
	GenerationNewPageTitleES         = "Nueva Generación"
	GenerationNewSubmitLabelES       = "Crear Generación"
	GenerationEditPageTitleES        = "Editar Generación"
	GenerationEditSubmitLabelES      = "Actualizar Generación"
	GenerationCreatedES              = "La generación fue creada satisfactoriamente"
	GenerationUpdatedES              = "La generación fue actualizada satisfactoriamente"
	GenerationCreateErrorsReceivedES = "Por favor corrija los errores para poder crear la generación"
	GenerationUpdateErrorsReceivedES = "Por favor corrija los errores para poder actualizar la generación"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[GenerationIndexPageTitleIndex] = GenerationIndexPageTitleES
		LangES[GenerationPageTitleIndex] = GenerationPageTitleES
		LangES[GenerationNewPageTitleIndex] = GenerationNewPageTitleES
		LangES[GenerationNewSubmitLabelIndex] = GenerationNewSubmitLabelES
		LangES[GenerationEditPageTitleIndex] = GenerationEditPageTitleES
		LangES[GenerationEditSubmitLabelIndex] = GenerationEditSubmitLabelES
		LangES[GenerationCreatedIndex] = GenerationCreatedES
		LangES[GenerationUpdatedIndex] = GenerationUpdatedES
		LangES[GenerationCreateErrorsReceivedIndex] = GenerationCreateErrorsReceivedES
		LangES[GenerationUpdateErrorsReceivedIndex] = GenerationUpdateErrorsReceivedES
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
