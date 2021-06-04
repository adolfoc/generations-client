package events

import "github.com/adolfoc/generations-client/handlers"

const (
	EventIndexPageTitleIndex       = 0
	EventPageTitleIndex            = 1
	EventNewPageTitleIndex         = 2
	EventNewSubmitLabelIndex       = 3
	EventEditPageTitleIndex        = 4
	EventEditSubmitLabelIndex      = 5
	EventCreatedIndex              = 6
	EventUpdatedIndex              = 7
	EventCreateErrorsReceivedIndex = 8
	EventUpdateErrorsReceivedIndex = 9

	EventIndexPageTitleES       = "Todos los eventos"
	EventPageTitleES            = "Evento"
	EventNewPageTitleES         = "Nuevo Evento"
	EventNewSubmitLabelES       = "Crear Evento"
	EventEditPageTitleES        = "Editar Evento"
	EventEditSubmitLabelES      = "Actualizar Evento"
	EventCreatedES              = "El evento fue creado satisfactoriamente"
	EventUpdatedES              = "El evento fue actualizado satisfactoriamente"
	EventCreateErrorsReceivedES = "Por favor corrija los errores para poder crear el evento"
	EventUpdateErrorsReceivedES = "Por favor corrija los errores para poder actualizar el evento"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[EventIndexPageTitleIndex] = EventIndexPageTitleES
		LangES[EventPageTitleIndex] = EventPageTitleES
		LangES[EventNewPageTitleIndex] = EventNewPageTitleES
		LangES[EventNewSubmitLabelIndex] = EventNewSubmitLabelES
		LangES[EventEditPageTitleIndex] = EventEditPageTitleES
		LangES[EventEditSubmitLabelIndex] = EventEditSubmitLabelES
		LangES[EventCreatedIndex] = EventCreatedES
		LangES[EventUpdatedIndex] = EventUpdatedES
		LangES[EventCreateErrorsReceivedIndex] = EventCreateErrorsReceivedES
		LangES[EventUpdateErrorsReceivedIndex] = EventUpdateErrorsReceivedES
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

