package model

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

type LifeSegment struct {
	ID          int        `json:"id"`
	PersonID    int        `json:"person_id"`
	LifePhase   *LifePhase `json:"life_phase"`
	Summary     string     `json:"summary"`
	Description string     `json:"description"`
}

func extractYear(date string) int {
	parts := strings.Split(date, "-")
	if len(parts) > 0 {
		year, _ := strconv.Atoi(parts[0])
		return year
	}

	return 0
}

func (ls *LifeSegment) TitleHTML(person *Person) template.HTML {
	birthYear := extractYear(person.BirthDate)
	startYear := birthYear + ls.LifePhase.StartYear
	endYear := birthYear + ls.LifePhase.EndYear
	return template.HTML(fmt.Sprintf("<b>%s</b>: %d--%d Años (%d--%d)",
		ls.LifePhase.Name, ls.LifePhase.StartYear, ls.LifePhase.EndYear, startYear, endYear))
}

func (ls *LifeSegment) IDInput(message, value string) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (ls *LifeSegment) PersonIDInput(message, value string) template.HTML {
	return BuildHiddenIDInput("inputPersonID", value)
}

func (ls *LifeSegment) LifePhaseIDInput(message, value string) template.HTML {
	return BuildHiddenIDInput("inputLifePhaseID", value)
}

func (ls *LifeSegment) SummaryLabel() template.HTML {
	return BuildLabel("inputSummary", "Resumen")
}

func (ls *LifeSegment) SummaryInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputSummary", value, message)
}

func (ls *LifeSegment) DescriptionLabel() template.HTML {
	return BuildLabel("inputDescription", "Descripción")
}

func (ls *LifeSegment) DescriptionInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputDescription", value, message)
}


type LifeSegmentRequest struct {
	ID          int    `json:"id"`
	PersonID    int    `json:"person_id"`
	LifePhaseID int    `json:"life_phase_id"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}
