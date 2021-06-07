package model

import "html/template"

type EventType struct {
	ID          int    `json:"id"`
	IsNatural   bool   `json:"is_natural"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (et *EventType) IDInput(message, value string) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (et *EventType) NameLabel() template.HTML {
	return BuildLabel("inputName", "Nombre")
}

func (et *EventType) NameInput(message, value string) template.HTML {
	return BuildTextInput("inputName", value, message)
}

func (et *EventType) IsNaturalLabel() template.HTML {
	return BuildLabel("inputIsNatural", "Es natural")
}

func (et *EventType) IsNaturalInput(message string, value bool) template.HTML {
	return BuildCheckboxInput("inputIsNatural", "Es natural", value, message)
}

func (et *EventType) DescriptionLabel() template.HTML {
	return BuildLabel("inputDescription", "Descripción")
}

func (et *EventType) DescriptionInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputDescription", value, message)
}

type EventTypes struct {
	EventTypes  []*EventType `json:"event_types"`
	RecordCount int          `json:"record_count"`
}

type EventTypeRequest struct {
	ID int				`json:"id"`
	IsNatural bool		`json:"is_natural"`
	Name string			`json:"name"`
	Description string	`json:"description"`
}
