package schemas

import "github.com/adolfoc/generations-client/handlers"

const (
	GenerationSchemaIndexPageTitleIndex       = 0
	GenerationSchemaPageTitleIndex            = 1
	GenerationSchemaNewPageTitleIndex         = 2
	GenerationSchemaNewSubmitLabelIndex       = 3
	GenerationSchemaEditPageTitleIndex        = 4
	GenerationSchemaEditSubmitLabelIndex      = 5
	GenerationSchemaCreatedIndex              = 6
	GenerationSchemaUpdatedIndex              = 7
	GenerationSchemaCreateErrorsReceivedIndex = 8
	GenerationSchemaUpdateErrorsReceivedIndex = 9

	GenerationSchemaIndexPageTitleES       = "Todos los esquemas"
	GenerationSchemaPageTitleES            = "Esquema Generacional"
	GenerationSchemaNewPageTitleES         = "Nuevo Esquema Generacional"
	GenerationSchemaNewSubmitLabelES       = "Crear Esquema Generacional"
	GenerationSchemaEditPageTitleES        = "Editar Esquema Generacional"
	GenerationSchemaEditSubmitLabelES      = "Actualizar Esquema Generacional"
	GenerationSchemaCreatedES              = "El esquema generacional fue creado satisfactoriamente"
	GenerationSchemaUpdatedES              = "El esquema generacional fue actualizado satisfactoriamente"
	GenerationSchemaCreateErrorsReceivedES = "Por favor corrija los errores para poder crear el esquema generacional"
	GenerationSchemaUpdateErrorsReceivedES = "Por favor corrija los errores para poder actualizar el esquema generacional"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[GenerationSchemaIndexPageTitleIndex] = GenerationSchemaIndexPageTitleES
		LangES[GenerationSchemaPageTitleIndex] = GenerationSchemaPageTitleES
		LangES[GenerationSchemaNewPageTitleIndex] = GenerationSchemaNewPageTitleES
		LangES[GenerationSchemaNewSubmitLabelIndex] = GenerationSchemaNewSubmitLabelES
		LangES[GenerationSchemaEditPageTitleIndex] = GenerationSchemaEditPageTitleES
		LangES[GenerationSchemaEditSubmitLabelIndex] = GenerationSchemaEditSubmitLabelES
		LangES[GenerationSchemaCreatedIndex] = GenerationSchemaCreatedES
		LangES[GenerationSchemaUpdatedIndex] = GenerationSchemaUpdatedES
		LangES[GenerationSchemaCreateErrorsReceivedIndex] = GenerationSchemaCreateErrorsReceivedES
		LangES[GenerationSchemaUpdateErrorsReceivedIndex] = GenerationSchemaUpdateErrorsReceivedES
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

