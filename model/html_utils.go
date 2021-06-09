package model

import (
	"fmt"
	"html/template"
)

func BuildHiddenIDInput(inputID string, value interface{}) template.HTML {
	return template.HTML(fmt.Sprintf("<input type=%q class=%q id=%q name=%q value=%d>\n",
		"text", "form-control d-none", inputID, inputID, value))
}

func BuildTextInput(inputID string, value interface{}, errorMessage string) template.HTML {
	if len(errorMessage) == 0 {
		return template.HTML(fmt.Sprintf("<input style='background-color: cornflowerblue;' type=%q class=%q id=%q name=%q value=%q>\n",
			"text", "form-control library-control", inputID, inputID, value))
	} else {
		inputCtl := fmt.Sprintf("<input type=%q class=%q id=%q name=%q value=%q>\n",
			"text", "form-control library-control is-invalid", inputID, inputID, value)
		msgDivID := fmt.Sprintf("%s%s", inputID, "Feedback")
		msgDiv := fmt.Sprintf("<div id=%q class=%q>%s</div>\n", msgDivID, "invalid-feedback", errorMessage)
		return template.HTML(fmt.Sprintf("%s%s", inputCtl, msgDiv))
	}
}

func BuildIntegerInput(inputID string, value interface{}, errorMessage string) template.HTML {
	if len(errorMessage) == 0 {
		return template.HTML(fmt.Sprintf("<input style='background-color: cornflowerblue;' type=%q class=%q id=%q name=%q value=%d>\n",
			"text", "form-control library-control", inputID, inputID, value))
	} else {
		inputCtl := fmt.Sprintf("<input type=%q class=%q id=%q name=%q value=%d>\n",
			"text", "form-control library-control is-invalid", inputID, inputID, value)
		msgDivID := fmt.Sprintf("%s%s", inputID, "Feedback")
		msgDiv := fmt.Sprintf("<div id=%q class=%q>%s</div>\n", msgDivID, "invalid-feedback", errorMessage)
		return template.HTML(fmt.Sprintf("%s%s", inputCtl, msgDiv))
	}
}

func BuildDateInput(inputID string, value interface{}, errorMessage string) template.HTML {
	if len(errorMessage) == 0 {
		return template.HTML(fmt.Sprintf("<input style='background-color: cornflowerblue;' type=%q class=%q id=%q name=%q value=%q placeholder=%q>\n",
			"date", "form-control library-control", inputID, inputID, value, "AAAA-MM-DD"))
	} else {
		inputCtl := fmt.Sprintf("<input type=%q class=%q id=%q name=%q value=%q placeholder=%q>\n",
			"date", "form-control library-control is-invalid", inputID, inputID, value,  "AAAA-MM-DD")
		msgDivID := fmt.Sprintf("%s%s", inputID, "Feedback")
		msgDiv := fmt.Sprintf("<div id=%q class=%q>%s</div>\n", msgDivID, "invalid-feedback", errorMessage)
		return template.HTML(fmt.Sprintf("%s%s", inputCtl, msgDiv))
	}
}

func BuildCheckboxInput(inputID, label string, value interface{}, errorMessage string) template.HTML {
	if value == true {
		return template.HTML(fmt.Sprintf("<label class=%q><input type=%q checked id=%q name=%q> %s</label>",
			"checkbox", "checkbox", inputID, inputID, label))
	} else {
		return template.HTML(fmt.Sprintf("<label class=%q><input type=%q id=%q name=%q> %s</label>",
			"checkbox", "checkbox", inputID, inputID, label))
	}
}

func BuildTextAreaInput(inputID string, value interface{}, errorMessage string) template.HTML {
	textBoxStyle := "background-color: cornflowerblue; height: 100%;"
	if len(errorMessage) == 0 {
		textBoxClass := "form-control library-control"
		taStart := fmt.Sprintf("<textarea style=%q class=%q id=%q name=%q rows='4'>",
			textBoxStyle, textBoxClass, inputID, inputID)
		taMiddle := value
		taEnd := "</textarea>"
		return template.HTML(fmt.Sprintf("%s\n%s\n%s\n", taStart, taMiddle, taEnd))
	} else {
		textBoxClass := "form-control library-control is-invalid"
		taStart := fmt.Sprintf("<textarea style=%q class=%q id=%q name=%q rows='4' >",
			textBoxStyle, textBoxClass, inputID, inputID)
		taMiddle := value
		taEnd := "</textarea>"
		inputCtl := fmt.Sprintf("%s\n%s\n%s\n", taStart, taMiddle, taEnd)
		msgDivID := fmt.Sprintf("%s%s", inputID, "Feedback")
		msgDiv := fmt.Sprintf("<div id=%q class=%q>%s</div>\n", msgDivID, "invalid-feedback", errorMessage)
		return template.HTML(fmt.Sprintf("%s%s", inputCtl, msgDiv))
	}
}

func BuildFileInput(inputID string, value interface{}, errorMessage string) template.HTML {
	if len(errorMessage) == 0 {
		return template.HTML(fmt.Sprintf("<input accept='application/pdf' type=%q class=%q id=%q name=%q value=%q>\n",
			"file", "form-control library-control", inputID, inputID, value))
	} else {
		inputCtl := fmt.Sprintf("<input accept='application/pdf' type=%q class=%q id=%q name=%q value=%q>\n",
			"file", "form-control library-control is-invalid", inputID, inputID, value)
		msgDivID := fmt.Sprintf("%s%s", inputID, "Feedback")
		msgDiv := fmt.Sprintf("<div id=%q class=%q>%s</div>\n", msgDivID, "invalid-feedback", errorMessage)
		return template.HTML(fmt.Sprintf("%s%s", inputCtl, msgDiv))
	}
}

func BuildNumberInput(inputID string, value interface{}, errorMessage string) template.HTML {
	normalizedValue := fmt.Sprintf("%d", value)
	if len(errorMessage) == 0 {
		return template.HTML(fmt.Sprintf("<input type=%q class=%q id=%q name=%q value=%q>\n",
			"number", "form-control library-control", inputID, inputID, normalizedValue))
	} else {
		inputCtl := fmt.Sprintf("<input type=%q class=%q id=%q name=%q value=%q>\n",
			"number", "form-control library-control is-invalid", inputID, inputID, normalizedValue)
		msgDivID := fmt.Sprintf("%s%s", inputID, "Feedback")
		msgDiv := fmt.Sprintf("<div id=%q class=%q>%s</div>\n", msgDivID, "invalid-feedback", errorMessage)
		return template.HTML(fmt.Sprintf("%s%s", inputCtl, msgDiv))
	}
}

func BuildEmailInput(inputID string, value interface{}, errorMessage string) template.HTML {
	if len(errorMessage) == 0 {
		return template.HTML(fmt.Sprintf("<input type=%q class=%q id=%q name=%q autocomplete=%q value=%q >\n",
			"email", "form-control library-control", inputID, inputID, "email", value))
	} else {
		inputCtl := fmt.Sprintf("<input type=%q class=%q id=%q name=%q value=%q>\n",
			"email", "form-control library-control is-invalid", inputID, inputID, value)
		msgDivID := fmt.Sprintf("%s%s", inputID, "Feedback")
		msgDiv := fmt.Sprintf("<div id=%q class=%q>%s</div>\n", msgDivID, "invalid-feedback", errorMessage)
		return template.HTML(fmt.Sprintf("%s%s", inputCtl, msgDiv))
	}
}

func BuildPasswordInput(inputID string, value interface{}, errorMessage string) template.HTML {
	if len(errorMessage) == 0 {
		return template.HTML(fmt.Sprintf("<input type=%q class=%q id=%q name=%q autocomplete=%q value=%q>\n",
			"password", "form-control library-control", inputID, inputID, "contraseña", value))
	} else {
		inputCtl := fmt.Sprintf("<input type=%q class=%q id=%q name=%q value=%q>\n",
			"password", "form-control library-control is-invalid", inputID, inputID, value)
		msgDivID := fmt.Sprintf("%s%s", inputID, "Feedback")
		msgDiv := fmt.Sprintf("<div id=%q class=%q>%s</div>\n", msgDivID, "invalid-feedback", errorMessage)
		return template.HTML(fmt.Sprintf("%s%s", inputCtl, msgDiv))
	}
}

func BuildLabel(inputID, label string) template.HTML {
	cssClass := "col-sm-2 col-form-label"
	return template.HTML(fmt.Sprintf("<label for=%q class=%q>%s</label>\n", inputID, cssClass, label))
}

