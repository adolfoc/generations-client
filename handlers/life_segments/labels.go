package life_segments

import "github.com/adolfoc/generations-client/handlers"

const (
	LifeSegmentIndexPageTitleIndex       = 0
	LifeSegmentPageTitleIndex            = 1
	LifeSegmentNewPageTitleIndex         = 2
	LifeSegmentNewSubmitLabelIndex       = 3
	LifeSegmentEditPageTitleIndex        = 4
	LifeSegmentEditSubmitLabelIndex      = 5
	LifeSegmentCreatedIndex              = 6
	LifeSegmentUpdatedIndex              = 7
	LifeSegmentCreateErrorsReceivedIndex = 8
	LifeSegmentUpdateErrorsReceivedIndex = 9

	LifeSegmentIndexPageTitleES       = "Todos los segmentos"
	LifeSegmentPageTitleES            = "Segmento"
	LifeSegmentNewPageTitleES         = "Nuevo Segmento"
	LifeSegmentNewSubmitLabelES       = "Crear Segmento"
	LifeSegmentEditPageTitleES        = "Editar Segmento"
	LifeSegmentEditSubmitLabelES      = "Actualizar Segmento"
	LifeSegmentCreatedES              = "El segmento fue creado satisfactoriamente"
	LifeSegmentUpdatedES              = "El segmento fue actualizado satisfactoriamente"
	LifeSegmentCreateErrorsReceivedES = "Por favor corrija los errores para poder crear el segmento"
	LifeSegmentUpdateErrorsReceivedES = "Por favor corrija los errores para poder actualizar el segmento"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[LifeSegmentIndexPageTitleIndex] = LifeSegmentIndexPageTitleES
		LangES[LifeSegmentPageTitleIndex] = LifeSegmentPageTitleES
		LangES[LifeSegmentNewPageTitleIndex] = LifeSegmentNewPageTitleES
		LangES[LifeSegmentNewSubmitLabelIndex] = LifeSegmentNewSubmitLabelES
		LangES[LifeSegmentEditPageTitleIndex] = LifeSegmentEditPageTitleES
		LangES[LifeSegmentEditSubmitLabelIndex] = LifeSegmentEditSubmitLabelES
		LangES[LifeSegmentCreatedIndex] = LifeSegmentCreatedES
		LangES[LifeSegmentUpdatedIndex] = LifeSegmentUpdatedES
		LangES[LifeSegmentCreateErrorsReceivedIndex] = LifeSegmentCreateErrorsReceivedES
		LangES[LifeSegmentUpdateErrorsReceivedIndex] = LifeSegmentUpdateErrorsReceivedES
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

