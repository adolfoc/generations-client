package moments

import "github.com/adolfoc/generations-client/handlers"

const (
	MomentIndexPageTitleIndex       = 0
	MomentPageTitleIndex            = 1
	MomentNewPageTitleIndex         = 2
	MomentNewSubmitLabelIndex       = 3
	MomentEditPageTitleIndex        = 4
	MomentEditSubmitLabelIndex      = 5
	MomentCreatedIndex              = 6
	MomentUpdatedIndex              = 7
	MomentCreateErrorsReceivedIndex = 8
	MomentUpdateErrorsReceivedIndex = 9

	MomentIndexPageTitleES       = "Todos los momentos"
	MomentPageTitleES            = "Momento"
	MomentNewPageTitleES         = "Nuevo Momento"
	MomentNewSubmitLabelES       = "Crear Momento"
	MomentEditPageTitleES        = "Editar Momento"
	MomentEditSubmitLabelES      = "Actualizar Momento"
	MomentCreatedES              = "El momento fue creado satisfactoriamente"
	MomentUpdatedES              = "El momento fue actualizado satisfactoriamente"
	MomentCreateErrorsReceivedES = "Por favor corrija los errores para poder crear el momento"
	MomentUpdateErrorsReceivedES = "Por favor corrija los errores para poder actualizar el momento"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[MomentIndexPageTitleIndex] = MomentIndexPageTitleES
		LangES[MomentPageTitleIndex] = MomentPageTitleES
		LangES[MomentNewPageTitleIndex] = MomentNewPageTitleES
		LangES[MomentNewSubmitLabelIndex] = MomentNewSubmitLabelES
		LangES[MomentEditPageTitleIndex] = MomentEditPageTitleES
		LangES[MomentEditSubmitLabelIndex] = MomentEditSubmitLabelES
		LangES[MomentCreatedIndex] = MomentCreatedES
		LangES[MomentUpdatedIndex] = MomentUpdatedES
		LangES[MomentCreateErrorsReceivedIndex] = MomentCreateErrorsReceivedES
		LangES[MomentUpdateErrorsReceivedIndex] = MomentUpdateErrorsReceivedES
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

