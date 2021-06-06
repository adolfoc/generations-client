package place_types

import "github.com/adolfoc/generations-client/handlers"

const (
	PlaceTypeIndexPageTitleIndex       = 0
	PlaceTypePageTitleIndex            = 1
	PlaceTypeNewPageTitleIndex         = 2
	PlaceTypeNewSubmitLabelIndex       = 3
	PlaceTypeEditPageTitleIndex        = 4
	PlaceTypeEditSubmitLabelIndex      = 5
	PlaceTypeCreatedIndex              = 6
	PlaceTypeUpdatedIndex              = 7
	PlaceTypeCreateErrorsReceivedIndex = 8
	PlaceTypeUpdateErrorsReceivedIndex = 9

	PlaceTypeIndexPageTitleES       = "Todos los tipos de lugares"
	PlaceTypePageTitleES            = "Tipo de lugar"
	PlaceTypeNewPageTitleES         = "Nuevo Tipo de lugar"
	PlaceTypeNewSubmitLabelES       = "Crear Tipo de lugar"
	PlaceTypeEditPageTitleES        = "Editar Tipo de lugar"
	PlaceTypeEditSubmitLabelES      = "Actualizar Tipo de lugar"
	PlaceTypeCreatedES              = "El tipo de lugar fue creado satisfactoriamente"
	PlaceTypeUpdatedES              = "El tipo de lugar fue actualizado satisfactoriamente"
	PlaceTypeCreateErrorsReceivedES = "Por favor corrija los errores para poder crear el tipo de lugar"
	PlaceTypeUpdateErrorsReceivedES = "Por favor corrija los errores para poder actualizar el tipo de lugar"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[PlaceTypeIndexPageTitleIndex] = PlaceTypeIndexPageTitleES
		LangES[PlaceTypePageTitleIndex] = PlaceTypePageTitleES
		LangES[PlaceTypeNewPageTitleIndex] = PlaceTypeNewPageTitleES
		LangES[PlaceTypeNewSubmitLabelIndex] = PlaceTypeNewSubmitLabelES
		LangES[PlaceTypeEditPageTitleIndex] = PlaceTypeEditPageTitleES
		LangES[PlaceTypeEditSubmitLabelIndex] = PlaceTypeEditSubmitLabelES
		LangES[PlaceTypeCreatedIndex] = PlaceTypeCreatedES
		LangES[PlaceTypeUpdatedIndex] = PlaceTypeUpdatedES
		LangES[PlaceTypeCreateErrorsReceivedIndex] = PlaceTypeCreateErrorsReceivedES
		LangES[PlaceTypeUpdateErrorsReceivedIndex] = PlaceTypeUpdateErrorsReceivedES
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

