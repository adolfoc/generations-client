package model

import (
	"fmt"
	"html/template"
	"strings"
)

type Event struct {
	ID          int        `json:"id"`
	EventType   *EventType `json:"event_type"`
	Name        string     `json:"name"`
	Start       string     `json:"start"`
	End         string     `json:"end"`
	Summary     string     `json:"summary"`
	Description string     `json:"description"`
}

func (e *Event) IDInput(message, value string) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (e *Event) NameLabel() template.HTML {
	return BuildLabel("inputName", "Nombre")
}

func (e *Event) NameInput(message, value string) template.HTML {
	return BuildTextInput("inputName", value, message)
}

func (e *Event) TypeIDLabel() template.HTML {
	return BuildLabel("inputTypeID", "Tipo de evento")
}

func (e *Event) TypeIDSelectBox(eventTypes []*EventType, selectedID int) template.HTML {
	var selectBox []string
	startSelect := fmt.Sprintf("<select class='form-select library-control' id='inputTypeID' name='inputTypeID'>")
	selectBox = append(selectBox, startSelect)

	blankOption := "<option value='' selected disabled hidden>Elija un tipo de evento...</option>"
	selectBox = append(selectBox, blankOption)

	for _, et := range eventTypes {
		var option string
		if et.ID == selectedID {
			option = fmt.Sprintf("<option value='%d' selected>%s</option>", et.ID, et.Name)
		} else {
			option = fmt.Sprintf("<option value='%d'>%s</option>", et.ID, et.Name)
		}
		selectBox = append(selectBox, option)
	}

	endSelect := fmt.Sprintf("</select>")
	selectBox = append(selectBox, endSelect)

	return template.HTML(strings.Join(selectBox, "\n"))
}

func (e *Event) StartLabel() template.HTML {
	return BuildLabel("inputStart", "Fecha de comienzo")
}

func (e *Event) StartInput(message, value string) template.HTML {
	return BuildTextInput("inputStart", value, message)
}

func (e *Event) EndLabel() template.HTML {
	return BuildLabel("inputEnd", "Fecha de término")
}

func (e *Event) EndInput(message, value string) template.HTML {
	return BuildTextInput("inputEnd", value, message)
}

func (e *Event) SummaryLabel() template.HTML {
	return BuildLabel("inputSummary", "Resumen")
}

func (e *Event) SummaryInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputSummary", value, message)
}

func (e *Event) DescriptionLabel() template.HTML {
	return BuildLabel("inputDescription", "Descripción")
}

func (e *Event) DescriptionInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputDescription", value, message)
}


type Events struct {
	Events      []*Event `json:"events"`
	RecordCount int      `json:"record_count"`
}

type EventRequest struct {
	ID          int    `json:"id"`
	TypeID      int    `json:"event_type_id"`
	Name        string `json:"name"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}
