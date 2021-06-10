package model

import "html/template"

type Intangible struct {
	ID          int    `json:"id"`
	LandscapeID int    `json:"landscape_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (it *Intangible) IDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (it *Intangible) LandscapeIDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputLandscapeID", value)
}

func (it *Intangible) NameLabel() template.HTML {
	return BuildLabel("inputName", "Nombre")
}

func (it *Intangible) NameInput(message, value string) template.HTML {
	return BuildTextInput("inputName", value, message)
}

func (it *Intangible) DescriptionLabel() template.HTML {
	return BuildLabel("inputDescription", "Descripción")
}

func (it *Intangible) DescriptionInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputDescription", value, message)
}


type IntangibleRequest struct {
	ID          int    `json:"id"`
	LandscapeID int    `json:"landscape_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
