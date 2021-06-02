package authentication

import "github.com/adolfoc/generations-client/handlers"

const (
	PageTitleIndex   = 0
	SubmitLabelIndex = 1

	PageTitleES   = "Ingresar a la Estudios Generacionales"
	SubmitLabelES = "Ingresar"
)

var LangES map[int]string

func initializeMaps() {
	if len(LangES) == 0 {
		LangES = make(map[int]string)
		LangES[PageTitleIndex] = PageTitleES
		LangES[SubmitLabelIndex] = SubmitLabelES
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


