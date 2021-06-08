package generational_landscape

import "github.com/adolfoc/generations-client/handlers"

const (
	GenerationalLandscapeIndexPageTitleIndex       = 0
	GenerationalLandscapePageTitleIndex            = 1
	GenerationalLandscapeNewPageTitleIndex         = 2
	GenerationalLandscapeNewSubmitLabelIndex       = 3
	GenerationalLandscapeEditPageTitleIndex        = 4
	GenerationalLandscapeEditSubmitLabelIndex      = 5
	GenerationalLandscapeCreatedIndex              = 6
	GenerationalLandscapeUpdatedIndex              = 7
	GenerationalLandscapeCreateErrorsReceivedIndex = 8
	GenerationalLandscapeUpdateErrorsReceivedIndex = 9

	GenerationalLandscapeIndexPageTitleES       = "Todos los paisajes generacionales"
	GenerationalLandscapePageTitleES            = "Paisaje generacional"
	GenerationalLandscapeNewPageTitleES         = "Nueva Paisaje generacional"
	GenerationalLandscapeNewSubmitLabelES       = "Crear Paisaje generacional"
	GenerationalLandscapeEditPageTitleES        = "Editar Paisaje generacional"
	GenerationalLandscapeEditSubmitLabelES      = "Actualizar Paisaje generacional"
	GenerationalLandscapeCreatedES              = "El paisaje generacional fue creado satisfactoriamente"
	GenerationalLandscapeUpdatedES              = "El paisaje generacional fue actualizado satisfactoriamente"
	GenerationalLandscapeCreateErrorsReceivedES = "Por favor corrija los errores para poder crear el paisaje generacional"
	GenerationalLandscapeUpdateErrorsReceivedES = "Por favor corrija los errores para poder actualizar el paisaje generacional"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[GenerationalLandscapeIndexPageTitleIndex] = GenerationalLandscapeIndexPageTitleES
		LangES[GenerationalLandscapePageTitleIndex] = GenerationalLandscapePageTitleES
		LangES[GenerationalLandscapeNewPageTitleIndex] = GenerationalLandscapeNewPageTitleES
		LangES[GenerationalLandscapeNewSubmitLabelIndex] = GenerationalLandscapeNewSubmitLabelES
		LangES[GenerationalLandscapeEditPageTitleIndex] = GenerationalLandscapeEditPageTitleES
		LangES[GenerationalLandscapeEditSubmitLabelIndex] = GenerationalLandscapeEditSubmitLabelES
		LangES[GenerationalLandscapeCreatedIndex] = GenerationalLandscapeCreatedES
		LangES[GenerationalLandscapeUpdatedIndex] = GenerationalLandscapeUpdatedES
		LangES[GenerationalLandscapeCreateErrorsReceivedIndex] = GenerationalLandscapeCreateErrorsReceivedES
		LangES[GenerationalLandscapeUpdateErrorsReceivedIndex] = GenerationalLandscapeUpdateErrorsReceivedES
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

