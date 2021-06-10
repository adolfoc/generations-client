package generational_landscape

import "github.com/adolfoc/generations-client/handlers"

const (
	GenerationalLandscapeIndexPageTitleIndex       = 0
	GenerationalLandscapePageTitleIndex            = 1
	GenerationalLandscapeNewPageTitleIndex         = 2
	GenerationalLandscapeNewSubmitLabelIndex       = 3
	GenerationalLandscapeEditPageTitleIndex        = 4
	GenerationalLandscapeEditSubmitLabelIndex      = 5
	GenerationalLandscapeCreatedIndex              = 6
	GenerationalLandscapeUpdatedIndex              = 7
	GenerationalLandscapeCreateErrorsReceivedIndex = 8
	GenerationalLandscapeUpdateErrorsReceivedIndex = 9
	TangibleNewPageTitleIndex                      = 10
	TangibleNewSubmitLabelIndex                    = 11
	TangibleCreateErrorsReceivedIndex              = 12
	TangibleCreatedIndex                           = 13
	TangibleEditPageTitleIndex                     = 14
	TangibleEditSubmitLabelIndex                   = 15
	TangibleUpdateErrorsReceivedIndex              = 16
	TangibleUpdatedIndex                           = 17
	IntangibleNewPageTitleIndex                    = 18
	IntangibleNewSubmitLabelIndex                  = 19
	IntangibleCreateErrorsReceivedIndex            = 20
	IntangibleCreatedIndex                         = 21
	IntangibleEditPageTitleIndex                   = 22
	IntangibleEditSubmitLabelIndex                 = 23
	IntangibleUpdateErrorsReceivedIndex            = 24
	IntangibleUpdatedIndex                         = 25

	GenerationalLandscapeIndexPageTitleES       = "Todos los paisajes generacionales"
	GenerationalLandscapePageTitleES            = "Paisaje generacional"
	GenerationalLandscapeNewPageTitleES         = "Nueva Paisaje generacional"
	GenerationalLandscapeNewSubmitLabelES       = "Crear Paisaje generacional"
	GenerationalLandscapeEditPageTitleES        = "Editar Paisaje generacional"
	GenerationalLandscapeEditSubmitLabelES      = "Actualizar Paisaje generacional"
	GenerationalLandscapeCreatedES              = "El paisaje generacional fue creado satisfactoriamente"
	GenerationalLandscapeUpdatedES              = "El paisaje generacional fue actualizado satisfactoriamente"
	GenerationalLandscapeCreateErrorsReceivedES = "Por favor corrija los errores para poder crear el paisaje generacional"
	GenerationalLandscapeUpdateErrorsReceivedES = "Por favor corrija los errores para poder actualizar el paisaje generacional"
	TangibleNewPageTitleES                      = "Nuevo Tangible"
	TangibleNewSubmitLabelES                    = "Crear Tangible"
	TangibleCreateErrorsReceivedES              = "Por favor corrija los errores para poder crear el tangible"
	TangibleCreatedES                           = "El tangible fue creado satisfactoriamente"
	TangibleEditPageTitleES                     = "Editar Tangible"
	TangibleEditSubmitLabelES                   = "Actualizar Tangible"
	TangibleUpdateErrorsReceivedES              = "Por favor corrija los errores para poder actualizar el tangible"
	TangibleUpdatedES                           = "El tangible fue actualizado satisfactoriamente"
	IntangibleNewPageTitleES                    = "Nuevo Intangible"
	IntangibleNewSubmitLabelES                  = "Actualizar Intangible"
	IntangibleCreateErrorsReceivedES            = "Por favor corrija los errores para poder crear el intangible"
	IntangibleCreatedES                         = "El intangible fue creado satisfactoriamente"
	IntangibleEditPageTitleES                   = "Editar Intangible"
	IntangibleEditSubmitLabelES                 = "Actualizar Intangible"
	IntangibleUpdateErrorsReceivedES            = "Por favor corrija los errores para poder actualizar el intangible"
	IntangibleUpdatedES                         = "El intangible fue actualizado satisfactoriamente"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[GenerationalLandscapeIndexPageTitleIndex] = GenerationalLandscapeIndexPageTitleES
		LangES[GenerationalLandscapePageTitleIndex] = GenerationalLandscapePageTitleES
		LangES[GenerationalLandscapeNewPageTitleIndex] = GenerationalLandscapeNewPageTitleES
		LangES[GenerationalLandscapeNewSubmitLabelIndex] = GenerationalLandscapeNewSubmitLabelES
		LangES[GenerationalLandscapeEditPageTitleIndex] = GenerationalLandscapeEditPageTitleES
		LangES[GenerationalLandscapeEditSubmitLabelIndex] = GenerationalLandscapeEditSubmitLabelES
		LangES[GenerationalLandscapeCreatedIndex] = GenerationalLandscapeCreatedES
		LangES[GenerationalLandscapeUpdatedIndex] = GenerationalLandscapeUpdatedES
		LangES[GenerationalLandscapeCreateErrorsReceivedIndex] = GenerationalLandscapeCreateErrorsReceivedES
		LangES[GenerationalLandscapeUpdateErrorsReceivedIndex] = GenerationalLandscapeUpdateErrorsReceivedES
		LangES[TangibleNewPageTitleIndex] = TangibleNewPageTitleES
		LangES[TangibleNewSubmitLabelIndex] = TangibleNewSubmitLabelES
		LangES[TangibleCreateErrorsReceivedIndex] = TangibleCreateErrorsReceivedES
		LangES[TangibleCreatedIndex] = TangibleCreatedES
		LangES[TangibleEditPageTitleIndex] = TangibleEditPageTitleES
		LangES[TangibleEditSubmitLabelIndex] = TangibleEditSubmitLabelES
		LangES[TangibleUpdateErrorsReceivedIndex] = TangibleUpdateErrorsReceivedES
		LangES[TangibleUpdatedIndex] = TangibleUpdatedES
		LangES[IntangibleNewPageTitleIndex] = IntangibleNewPageTitleES
		LangES[IntangibleNewSubmitLabelIndex] = IntangibleNewSubmitLabelES
		LangES[IntangibleCreateErrorsReceivedIndex] = IntangibleCreateErrorsReceivedES
		LangES[IntangibleCreatedIndex] = IntangibleCreatedES
		LangES[IntangibleEditPageTitleIndex] = IntangibleEditPageTitleES
		LangES[IntangibleEditSubmitLabelIndex] = IntangibleEditSubmitLabelES
		LangES[IntangibleUpdateErrorsReceivedIndex] = IntangibleUpdateErrorsReceivedES
		LangES[IntangibleUpdatedIndex] = IntangibleUpdatedES
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

