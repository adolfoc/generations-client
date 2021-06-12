package handlers

import "net/http"

type FormTemplate struct {
	Ct                CommonTemplate
	URL               string
	SubmitLabel       string
	FormValues        map[string]interface{}
	FormErrorMessages map[string]string
}

func MakeFormTemplate(r *http.Request, url, pageTitle, studyTitle, submitLabel string, formValues map[string]interface{},
	errorMessages map[string]string) (*FormTemplate, error) {

	ct, err := MakeCommonTemplate(r, pageTitle, studyTitle)
	if err != nil {
		return nil, err
	}

	ft := &FormTemplate{
		Ct:                *ct,
		URL:               url,
		SubmitLabel:       submitLabel,
		FormValues:        formValues,
		FormErrorMessages: errorMessages,
	}

	return ft, nil
}

func MakeSimpleFormTemplate(url, pageTitle, submitLabel string, formValues map[string]interface{},
	errorMessages map[string]string) *FormTemplate {

	ct := MakeSimpleTemplate(pageTitle)

	ft := &FormTemplate{
		Ct:                *ct,
		URL:               url,
		SubmitLabel:       submitLabel,
		FormValues:        formValues,
		FormErrorMessages: errorMessages,
	}

	return ft
}


