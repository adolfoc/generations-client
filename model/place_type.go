package model

import "html/template"

type PlaceType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (pt *PlaceType) IDInput(message, value string) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (pt *PlaceType) NameLabel() template.HTML {
	return BuildLabel("inputName", "Nombre")
}

func (pt *PlaceType) NameInput(message, value string) template.HTML {
	return BuildTextInput("inputName", value, message)
}

func (pt *PlaceType) DescriptionLabel() template.HTML {
	return BuildLabel("inputDescription", "Descripción")
}

func (pt *PlaceType) DescriptionInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputDescription", value, message)
}

type PlaceTypes struct {
	PlaceTypes  []*PlaceType `json:"place_types"`
	RecordCount int          `json:"record_count"`
}

type PlaceTypeRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
