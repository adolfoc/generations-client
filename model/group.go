package model

import (
	"fmt"
	"html/template"
	"strings"
)

type Group struct {
	ID            int                   `json:"id"`
	Name          string                `json:"name"`
	GroupTypeID   int                   `json:"group_type_id"`
	ParentGroupID int                   `json:"parent_group_id"`
	Start         string                `json:"start_date"`
	End           string                `json:"end_date"`
	Summary       string                `json:"summary"`
	Members       []*GroupMembershipMap `json:"members"`
	Description   string                `json:"description"`
	GroupType     *GroupType            `json:"group_type"`
	GroupName     string                `json:"group_name"`
}

func (g *Group) IDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (g *Group) NameLabel() template.HTML {
	return BuildLabel("inputName", "Nombre")
}

func (g *Group) NameInput(message, value string) template.HTML {
	return BuildTextInput("inputName", value, message)
}

func (g *Group) TypeIDLabel() template.HTML {
	return BuildLabel("inputGroupTypeID", "Tipo de grupo")
}

func (g *Group) TypeIDSelectBox(groupTypes []*GroupType, selectedID int) template.HTML {
	var selectBox []string
	startSelect := fmt.Sprintf("<select class='form-select library-control' id='inputGroupTypeID' name='inputGroupTypeID'>")
	selectBox = append(selectBox, startSelect)

	blankOption := "<option value='' selected disabled hidden>Elija un tipo de grupo...</option>"
	selectBox = append(selectBox, blankOption)

	for _, gt := range groupTypes {
		var option string
		if gt.ID == selectedID {
			option = fmt.Sprintf("<option value='%d' selected>%s</option>", gt.ID, gt.Name)
		} else {
			option = fmt.Sprintf("<option value='%d'>%s</option>", gt.ID, gt.Name)
		}
		selectBox = append(selectBox, option)
	}

	endSelect := fmt.Sprintf("</select>")
	selectBox = append(selectBox, endSelect)

	return template.HTML(strings.Join(selectBox, "\n"))
}

func (g *Group) StartLabel() template.HTML {
	return BuildLabel("inputStart", "Fecha de comienzo")
}

func (g *Group) StartInput(message, value string) template.HTML {
	return BuildTextInput("inputStart", value, message)
}

func (g *Group) EndLabel() template.HTML {
	return BuildLabel("inputEnd", "Fecha de término")
}

func (g *Group) EndInput(message, value string) template.HTML {
	return BuildTextInput("inputEnd", value, message)
}

func (g *Group) SummaryLabel() template.HTML {
	return BuildLabel("inputSummary", "Resumen")
}

func (g *Group) SummaryInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputSummary", value, message)
}

func (g *Group) DescriptionLabel() template.HTML {
	return BuildLabel("inputDescription", "Descripción")
}

func (g *Group) DescriptionInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputDescription", value, message)
}


type Groups struct {
	Groups      []*Group `json:"groups"`
	RecordCount int      `json:"record_count"`
}

type GroupRequest struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	GroupTypeID   int    `json:"group_type_id"`
	ParentGroupID int    `json:"parent_group_id"`
	Start         string `json:"start_date"`
	End           string `json:"end_date"`
	Summary       string `json:"summary"`
	Description   string `json:"description"`
}
