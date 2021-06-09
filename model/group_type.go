package model

import "html/template"

type GroupType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (gt *GroupType) IDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (gt *GroupType) NameLabel() template.HTML {
	return BuildLabel("inputName", "Nombre")
}

func (gt *GroupType) NameInput(message, value string) template.HTML {
	return BuildTextInput("inputName", value, message)
}

func (gt *GroupType) DescriptionLabel() template.HTML {
	return BuildLabel("inputDescription", "Descripción")
}

func (gt *GroupType) DescriptionInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputDescription", value, message)
}

type GroupTypes struct {
	GroupTypes  []*GroupType `json:"group_types"`
	RecordCount int          `json:"record_count"`
}

type GroupTypeRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
