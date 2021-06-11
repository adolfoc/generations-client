package life_phases

import "github.com/adolfoc/generations-client/handlers"

const (
	LifePhaseIndexPageTitleIndex       = 0
	LifePhasePageTitleIndex            = 1
	LifePhaseNewPageTitleIndex         = 2
	LifePhaseNewSubmitLabelIndex       = 3
	LifePhaseEditPageTitleIndex        = 4
	LifePhaseEditSubmitLabelIndex      = 5
	LifePhaseCreatedIndex              = 6
	LifePhaseUpdatedIndex              = 7
	LifePhaseCreateErrorsReceivedIndex = 8
	LifePhaseUpdateErrorsReceivedIndex = 9
	LifePhaseDeletedIndex              = 10

	LifePhaseIndexPageTitleES       = "Todos las etapas de vida"
	LifePhasePageTitleES            = "Etapa de Vida"
	LifePhaseNewPageTitleES         = "Nueva Etapa de Vida"
	LifePhaseNewSubmitLabelES       = "Crear Etapa de Vida"
	LifePhaseEditPageTitleES        = "Editar Etapa de Vida"
	LifePhaseEditSubmitLabelES      = "Actualizar Etapa de Vida"
	LifePhaseCreatedES              = "La etapa de vida fue creada satisfactoriamente"
	LifePhaseUpdatedES              = "La etapa de vida fue actualizada satisfactoriamente"
	LifePhaseCreateErrorsReceivedES = "Por favor corrija los errores para poder crear la etapa de vida"
	LifePhaseUpdateErrorsReceivedES = "Por favor corrija los errores para poder actualizar la etapa de vida"
	LifePhaseDeletedES              = "La etapa de vida fue eliminada satisfactoriamente"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[LifePhaseIndexPageTitleIndex] = LifePhaseIndexPageTitleES
		LangES[LifePhasePageTitleIndex] = LifePhasePageTitleES
		LangES[LifePhaseNewPageTitleIndex] = LifePhaseNewPageTitleES
		LangES[LifePhaseNewSubmitLabelIndex] = LifePhaseNewSubmitLabelES
		LangES[LifePhaseEditPageTitleIndex] = LifePhaseEditPageTitleES
		LangES[LifePhaseEditSubmitLabelIndex] = LifePhaseEditSubmitLabelES
		LangES[LifePhaseCreatedIndex] = LifePhaseCreatedES
		LangES[LifePhaseUpdatedIndex] = LifePhaseUpdatedES
		LangES[LifePhaseCreateErrorsReceivedIndex] = LifePhaseCreateErrorsReceivedES
		LangES[LifePhaseUpdateErrorsReceivedIndex] = LifePhaseUpdateErrorsReceivedES
		LangES[LifePhaseDeletedIndex] = LifePhaseDeletedES
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


