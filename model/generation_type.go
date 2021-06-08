package model

import "html/template"

type GenerationType struct {
	ID          int    `json:"id"`
	SchemaID    int    `json:"schema_id"`
	Archetype   string `json:"archetype"`
	Description string `json:"description"`
}

func (gt *GenerationType) IDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (gt *GenerationType) SchemaIDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputSchemaID", value)
}

func (gt *GenerationType) ArchetypeLabel() template.HTML {
	return BuildLabel("inputArchetype", "Nombre")
}

func (gt *GenerationType) ArchetypeInput(message, value string) template.HTML {
	return BuildTextInput("inputArchetype", value, message)
}

func (gt *GenerationType) DescriptionLabel() template.HTML {
	return BuildLabel("inputDescription", "Descripción")
}

func (gt *GenerationType) DescriptionInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputDescription", value, message)
}

type GenerationTypeRequest struct {
	ID          int    `json:"id"`
	SchemaID    int    `json:"schema_id"`
	Archetype   string `json:"archetype"`
	Description string `json:"description"`
}
