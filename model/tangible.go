package model

import "html/template"

type Tangible struct {
	ID          int    `json:"id"`
	LandscapeID int    `json:"landscape_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (t *Tangible) IDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (t *Tangible) LandscapeIDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputLandscapeID", value)
}

func (t *Tangible) NameLabel() template.HTML {
	return BuildLabel("inputName", "Nombre")
}

func (t *Tangible) NameInput(message, value string) template.HTML {
	return BuildTextInput("inputName", value, message)
}

func (t *Tangible) DescriptionLabel() template.HTML {
	return BuildLabel("inputDescription", "Descripción")
}

func (t *Tangible) DescriptionInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputDescription", value, message)
}


type TangibleRequest struct {
	ID          int    `json:"id"`
	LandscapeID int    `json:"landscape_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
