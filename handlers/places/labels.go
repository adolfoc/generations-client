package places

import "github.com/adolfoc/generations-client/handlers"

const (
	PlaceIndexPageTitleIndex        = 0
	PlacePageTitleIndex             = 1
	PlaceNewPageTitleIndex          = 2
	PlaceNewSubmitPlaceIndex        = 3
	PlaceEditPageTitleIndex         = 4
	PlaceEditSubmitPlaceIndex       = 5
	PlaceCreatedIndex               = 6
	PlaceUpdatedIndex               = 7
	PlaceCreateErrorsReceivedIndex  = 8
	PlaceUpdateErrorsReceivedIndex  = 9
	PlaceAddSegmentsPageTitleIndex  = 10
	PlaceAddSegmentSubmitPlaceIndex = 11

	PlaceIndexPageTitleES        = "Todos los lugares"
	PlacePageTitleES             = "Lugar"
	PlaceNewPageTitleES          = "Nuevo Lugar"
	PlaceNewSubmitPlaceES        = "Crear Lugar"
	PlaceEditPageTitleES         = "Editar Lugar"
	PlaceEditSubmitPlaceES       = "Actualizar Lugar"
	PlaceCreatedES               = "El lugar fue creado satisfactoriamente"
	PlaceUpdatedES               = "El lugar fue actualizado satisfactoriamente"
	PlaceCreateErrorsReceivedES  = "Por favor corrija los errores para poder crear el lugar"
	PlaceUpdateErrorsReceivedES  = "Por favor corrija los errores para poder actualizar el lugar"
	PlaceAddSegmentsPageTitleES  = "Agregar lugar"
	PlaceAddSegmentSubmitPlaceES = "Generar lugar"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[PlaceIndexPageTitleIndex] = PlaceIndexPageTitleES
		LangES[PlacePageTitleIndex] = PlacePageTitleES
		LangES[PlaceNewPageTitleIndex] = PlaceNewPageTitleES
		LangES[PlaceNewSubmitPlaceIndex] = PlaceNewSubmitPlaceES
		LangES[PlaceEditPageTitleIndex] = PlaceEditPageTitleES
		LangES[PlaceEditSubmitPlaceIndex] = PlaceEditSubmitPlaceES
		LangES[PlaceCreatedIndex] = PlaceCreatedES
		LangES[PlaceUpdatedIndex] = PlaceUpdatedES
		LangES[PlaceCreateErrorsReceivedIndex] = PlaceCreateErrorsReceivedES
		LangES[PlaceUpdateErrorsReceivedIndex] = PlaceUpdateErrorsReceivedES
		LangES[PlaceAddSegmentsPageTitleIndex] = PlaceAddSegmentsPageTitleES
		LangES[PlaceAddSegmentSubmitPlaceIndex] = PlaceAddSegmentSubmitPlaceES
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


