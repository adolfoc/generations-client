package groups

import "github.com/adolfoc/generations-client/handlers"

const (
	GroupIndexPageTitleIndex       = 0
	GroupPageTitleIndex            = 1
	GroupNewPageTitleIndex         = 2
	GroupNewSubmitLabelIndex       = 3
	GroupEditPageTitleIndex        = 4
	GroupEditSubmitLabelIndex      = 5
	GroupCreatedIndex              = 6
	GroupUpdatedIndex              = 7
	GroupCreateErrorsReceivedIndex = 8
	GroupUpdateErrorsReceivedIndex = 9

	GroupIndexPageTitleES       = "Todos los grupos"
	GroupPageTitleES            = "Grupo"
	GroupNewPageTitleES         = "Nuevo Grupo"
	GroupNewSubmitLabelES       = "Crear Grupo"
	GroupEditPageTitleES        = "Editar Grupo"
	GroupEditSubmitLabelES      = "Actualizar Grupo"
	GroupCreatedES              = "El grupo fue creado satisfactoriamente"
	GroupUpdatedES              = "El grupo fue actualizado satisfactoriamente"
	GroupCreateErrorsReceivedES = "Por favor corrija los errores para poder crear el grupo"
	GroupUpdateErrorsReceivedES = "Por favor corrija los errores para poder actualizar el grupo"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[GroupIndexPageTitleIndex] = GroupIndexPageTitleES
		LangES[GroupPageTitleIndex] = GroupPageTitleES
		LangES[GroupNewPageTitleIndex] = GroupNewPageTitleES
		LangES[GroupNewSubmitLabelIndex] = GroupNewSubmitLabelES
		LangES[GroupEditPageTitleIndex] = GroupEditPageTitleES
		LangES[GroupEditSubmitLabelIndex] = GroupEditSubmitLabelES
		LangES[GroupCreatedIndex] = GroupCreatedES
		LangES[GroupUpdatedIndex] = GroupUpdatedES
		LangES[GroupCreateErrorsReceivedIndex] = GroupCreateErrorsReceivedES
		LangES[GroupUpdateErrorsReceivedIndex] = GroupUpdateErrorsReceivedES
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

