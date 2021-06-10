package generation_positions

import "github.com/adolfoc/generations-client/handlers"

const (
	GenerationPositionIndexPageTitleIndex       = 0
	GenerationPositionPageTitleIndex            = 1
	GenerationPositionNewPageTitleIndex         = 2
	GenerationPositionNewSubmitLabelIndex       = 3
	GenerationPositionEditPageTitleIndex        = 4
	GenerationPositionEditSubmitLabelIndex      = 5
	GenerationPositionCreatedIndex              = 6
	GenerationPositionUpdatedIndex              = 7
	GenerationPositionCreateErrorsReceivedIndex = 8
	GenerationPositionUpdateErrorsReceivedIndex = 9

	GenerationPositionIndexPageTitleES       = "Todas las posiciones generacionales"
	GenerationPositionPageTitleES            = "Posición Generacional"
	GenerationPositionNewPageTitleES         = "Nueva Posición Generacional"
	GenerationPositionNewSubmitLabelES       = "Crear Posición Generacional"
	GenerationPositionEditPageTitleES        = "Editar Posición Generacional"
	GenerationPositionEditSubmitLabelES      = "Actualizar Posición Generacional"
	GenerationPositionCreatedES              = "La posición generacional fue creada satisfactoriamente"
	GenerationPositionUpdatedES              = "La posición generacional fue actualizada satisfactoriamente"
	GenerationPositionCreateErrorsReceivedES = "Por favor corrija los errores para poder crear la posición generacional"
	GenerationPositionUpdateErrorsReceivedES = "Por favor corrija los errores para poder actualizar la posición generacional"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[GenerationPositionIndexPageTitleIndex] = GenerationPositionIndexPageTitleES
		LangES[GenerationPositionPageTitleIndex] = GenerationPositionPageTitleES
		LangES[GenerationPositionNewPageTitleIndex] = GenerationPositionNewPageTitleES
		LangES[GenerationPositionNewSubmitLabelIndex] = GenerationPositionNewSubmitLabelES
		LangES[GenerationPositionEditPageTitleIndex] = GenerationPositionEditPageTitleES
		LangES[GenerationPositionEditSubmitLabelIndex] = GenerationPositionEditSubmitLabelES
		LangES[GenerationPositionCreatedIndex] = GenerationPositionCreatedES
		LangES[GenerationPositionUpdatedIndex] = GenerationPositionUpdatedES
		LangES[GenerationPositionCreateErrorsReceivedIndex] = GenerationPositionCreateErrorsReceivedES
		LangES[GenerationPositionUpdateErrorsReceivedIndex] = GenerationPositionUpdateErrorsReceivedES
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


