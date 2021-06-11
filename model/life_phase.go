package model

import (
	"fmt"
	"html/template"
)

type LifePhase struct {
	ID        int    `json:"id"`
	SchemaID  int    `json:"schema_id"`
	Name      string `json:"name"`
	StartYear int    `json:"start_year"`
	EndYear   int    `json:"end_year"`
	Role      string `json:"role"`
}

func (lp *LifePhase) Span() template.HTML {
	return template.HTML(fmt.Sprintf("%d&mdash;%d", lp.StartYear, lp.EndYear))
}

func (lp *LifePhase) CalendarSpan(generation *Generation, moment *HistoricalMoment) template.HTML {
	fromCalendar := generation.StartYear + lp.StartYear
	toCalendar := generation.EndYear + lp.StartYear

	return template.HTML(fmt.Sprintf("%d&mdash;%d (%d&mdash;%d)", lp.StartYear, lp.EndYear, fromCalendar, toCalendar))
}


func (lp *LifePhase) IDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (lp *LifePhase) SchemaIDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputSchemaID", value)
}

func (lp *LifePhase) NameLabel() template.HTML {
	return BuildLabel("inputName", "Nombre")
}

func (lp *LifePhase) NameInput(message, value string) template.HTML {
	return BuildTextInput("inputName", value, message)
}

func (lp *LifePhase) StartYearLabel() template.HTML {
	return BuildLabel("inputStartYearName", "Año inicial")
}

func (lp *LifePhase) StartYearInput(message string, value int) template.HTML {
	return BuildIntegerInput("inputStartYear", value, message)
}

func (lp *LifePhase) EndYearLabel() template.HTML {
	return BuildLabel("inputEndYear", "Año de término")
}

func (lp *LifePhase) EndYearInput(message string, value int) template.HTML {
	return BuildIntegerInput("inputEndYear", value, message)
}

func (lp *LifePhase) RoleLabel() template.HTML {
	return BuildLabel("inputRole", "Rol")
}

func (lp *LifePhase) RoleInput(message, value string) template.HTML {
	return BuildTextInput("inputRole", value, message)
}

type LifePhaseRequest struct {
	ID        int    `json:"id"`
	SchemaID  int    `json:"schema_id"`
	Name      string `json:"name"`
	StartYear int    `json:"start_year"`
	EndYear   int    `json:"end_year"`
	Role      string `json:"role"`
}
