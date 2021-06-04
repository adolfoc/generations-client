package persons

import "github.com/adolfoc/generations-client/handlers"

const (
	PersonIndexPageTitleIndex        = 0
	PersonPageTitleIndex             = 1
	PersonNewPageTitleIndex          = 2
	PersonNewSubmitLabelIndex        = 3
	PersonEditPageTitleIndex         = 4
	PersonEditSubmitLabelIndex       = 5
	PersonCreatedIndex               = 6
	PersonUpdatedIndex               = 7
	PersonCreateErrorsReceivedIndex  = 8
	PersonUpdateErrorsReceivedIndex  = 9
	PersonAddSegmentsPageTitleIndex  = 10
	PersonAddSegmentSubmitLabelIndex = 11

	PersonIndexPageTitleES        = "Todas las personas"
	PersonPageTitleES             = "Persona"
	PersonNewPageTitleES          = "Nueva Persona"
	PersonNewSubmitLabelES        = "Crear Persona"
	PersonEditPageTitleES         = "Editar Persona"
	PersonEditSubmitLabelES       = "Actualizar Persona"
	PersonCreatedES               = "La persona fue creada satisfactoriamente"
	PersonUpdatedES               = "La persona fue actualizada satisfactoriamente"
	PersonCreateErrorsReceivedES  = "Por favor corrija los errores para poder crear la persona"
	PersonUpdateErrorsReceivedES  = "Por favor corrija los errores para poder actualizar la persona"
	PersonAddSegmentsPageTitleES  = "Agregar segmentos"
	PersonAddSegmentSubmitLabelES = "Generar segmentos"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[PersonIndexPageTitleIndex] = PersonIndexPageTitleES
		LangES[PersonPageTitleIndex] = PersonPageTitleES
		LangES[PersonNewPageTitleIndex] = PersonNewPageTitleES
		LangES[PersonNewSubmitLabelIndex] = PersonNewSubmitLabelES
		LangES[PersonEditPageTitleIndex] = PersonEditPageTitleES
		LangES[PersonEditSubmitLabelIndex] = PersonEditSubmitLabelES
		LangES[PersonCreatedIndex] = PersonCreatedES
		LangES[PersonUpdatedIndex] = PersonUpdatedES
		LangES[PersonCreateErrorsReceivedIndex] = PersonCreateErrorsReceivedES
		LangES[PersonUpdateErrorsReceivedIndex] = PersonUpdateErrorsReceivedES
		LangES[PersonAddSegmentsPageTitleIndex] = PersonAddSegmentsPageTitleES
		LangES[PersonAddSegmentSubmitLabelIndex] = PersonAddSegmentSubmitLabelES
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

