package event_types

import "github.com/adolfoc/generations-client/handlers"

const (
	EventTypeIndexPageTitleIndex       = 0
	EventTypePageTitleIndex            = 1
	EventTypeNewPageTitleIndex         = 2
	EventTypeNewSubmitLabelIndex       = 3
	EventTypeEditPageTitleIndex        = 4
	EventTypeEditSubmitLabelIndex      = 5
	EventTypeCreatedIndex              = 6
	EventTypeUpdatedIndex              = 7
	EventTypeCreateErrorsReceivedIndex = 8
	EventTypeUpdateErrorsReceivedIndex = 9

	EventTypeIndexPageTitleES       = "Todos los tipos de evento"
	EventTypePageTitleES            = "Tipo de Evento"
	EventTypeNewPageTitleES         = "Tipo de Evento"
	EventTypeNewSubmitLabelES       = "Crear Tipo de Evento"
	EventTypeEditPageTitleES        = "Editar Tipo de Evento"
	EventTypeEditSubmitLabelES      = "Actualizar Tipo de Evento"
	EventTypeCreatedES              = "El tipo de evento fue creado satisfactoriamente"
	EventTypeUpdatedES              = "El tipo de evento fue actualizado satisfactoriamente"
	EventTypeCreateErrorsReceivedES = "Por favor corrija los errores para poder crear el tipo de evento"
	EventTypeUpdateErrorsReceivedES = "Por favor corrija los errores para poder actualizar el tipo de evento"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[EventTypeIndexPageTitleIndex] = EventTypeIndexPageTitleES
		LangES[EventTypePageTitleIndex] = EventTypePageTitleES
		LangES[EventTypeNewPageTitleIndex] = EventTypeNewPageTitleES
		LangES[EventTypeNewSubmitLabelIndex] = EventTypeNewSubmitLabelES
		LangES[EventTypeEditPageTitleIndex] = EventTypeEditPageTitleES
		LangES[EventTypeEditSubmitLabelIndex] = EventTypeEditSubmitLabelES
		LangES[EventTypeCreatedIndex] = EventTypeCreatedES
		LangES[EventTypeUpdatedIndex] = EventTypeUpdatedES
		LangES[EventTypeCreateErrorsReceivedIndex] = EventTypeCreateErrorsReceivedES
		LangES[EventTypeUpdateErrorsReceivedIndex] = EventTypeUpdateErrorsReceivedES
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

