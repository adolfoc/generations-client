package model

import (
	"fmt"
	"html/template"
	"strings"
)

type GenerationalLandscape struct {
	ID                int           `json:"id"`
	GenerationID      int           `json:"generation_id"`
	FormationMomentID int           `json:"formation_moment_id"`
	Description       string        `json:"description"`
	Tangibles         []*Tangible   `json:"tangibles"`
	Intangibles       []*Intangible `json:"intangibles"`
}

func (gl *GenerationalLandscape) IDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (gl *GenerationalLandscape) GenerationIDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputGenerationID", value)
}

func (gl *GenerationalLandscape) FormationMomentIDLabel() template.HTML {
	return BuildLabel("inputFormationMomentID", "Momento de formación")
}

func buildMomentName(moment *HistoricalMoment) string {
	return fmt.Sprintf("%s (%d-%d)", moment.Name, moment.StartYear(), moment.EndYear())
}

func (gl *GenerationalLandscape) FormationMomentIDSelectBox(moments []*HistoricalMoment, selectedID int) template.HTML {
	var selectBox []string
	startSelect := fmt.Sprintf("<select class='form-select library-control' id='inputFormationMomentID' name='inputFormationMomentID'>")
	selectBox = append(selectBox, startSelect)

	for _, moment := range moments {
		var option string
		if moment.ID == selectedID {
			option = fmt.Sprintf("<option value='%d' selected>%s</option>", moment.ID, buildMomentName(moment))
		} else {
			option = fmt.Sprintf("<option value='%d'>%s</option>", moment.ID, buildMomentName(moment))
		}
		selectBox = append(selectBox, option)
	}

	endSelect := fmt.Sprintf("</select>")
	selectBox = append(selectBox, endSelect)

	return template.HTML(strings.Join(selectBox, "\n"))
}

func (gl *GenerationalLandscape) DescriptionLabel() template.HTML {
	return BuildLabel("inputDescription", "Descripción")
}

func (gl *GenerationalLandscape) DescriptionInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputDescription", value, message)
}

type GenerationalLandscapeRequest struct {
	ID                int    `json:"id"`
	GenerationID      int    `json:"generation_id"`
	FormationMomentID int    `json:"formation_moment_id"`
	Description       string `json:"description"`
}
