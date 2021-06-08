package model

import "html/template"

type MomentType struct {
	ID          int    `json:"id"`
	SchemaID    int    `json:"schema_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (mt *MomentType) IDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (mt *MomentType) SchemaIDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputSchemaID", value)
}

func (mt *MomentType) NameLabel() template.HTML {
	return BuildLabel("inputName", "Nombre")
}

func (mt *MomentType) NameInput(message, value string) template.HTML {
	return BuildTextInput("inputName", value, message)
}

func (mt *MomentType) DescriptionLabel() template.HTML {
	return BuildLabel("inputDescription", "Descripción")
}

func (mt *MomentType) DescriptionInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputDescription", value, message)
}

type MomentTypeRequest struct {
	ID          int    `json:"id"`
	SchemaID    int    `json:"schema_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
