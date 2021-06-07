package model

import (
	"fmt"
	"html/template"
	"strconv"
	"strings"
)

type HistoricalMoment struct {
	ID              int                   `json:"id"`
	Name            string                `json:"name"`
	SchemaID        int                   `json:"schema_id"`
	Type            *MomentType           `json:"moment_type"`
	Start           string                `json:"start_date"`
	End             string                `json:"end_date"`
	Summary         string                `json:"summary"`
	Description     string                `json:"description"`
	DescriptionHTML string                `json:"description_html"`
	Positions       []*GenerationPosition `json:"positions"`
}

func (hm *HistoricalMoment) extractYear(date string) int {
	parts := strings.Split(date, "-")
	if len(parts) > 0 {
		year, _ := strconv.Atoi(parts[0])
		return year
	}

	return 0
}

func (hm *HistoricalMoment) StartYear() int {
	return extractYear(hm.Start)
}

func (hm *HistoricalMoment) EndYear() int {
	return extractYear(hm.End)
}

func (hm *HistoricalMoment) Span() template.HTML {
	return template.HTML(fmt.Sprintf("%s&mdash;%s", hm.Start, hm.End))
}

func (hm *HistoricalMoment) IDInput(message, value string) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (hm *HistoricalMoment) SchemaIDInput(message, value string) template.HTML {
	return BuildHiddenIDInput("inputSchemaID", value)
}

func (hm *HistoricalMoment) NameLabel() template.HTML {
	return BuildLabel("inputName", "Nombre")
}

func (hm *HistoricalMoment) NameInput(message, value string) template.HTML {
	return BuildTextInput("inputName", value, message)
}

func (hm *HistoricalMoment) TypeLabel() template.HTML {
	return BuildLabel("inputTypeID", "Tipo")
}

func (hm *HistoricalMoment) TypeSelectBox(momentTypes []*MomentType, selectedID int) template.HTML {
	var selectBox []string
	startSelect := fmt.Sprintf("<select class='form-select library-control' id='inputTypeID' name='inputTypeID'>")
	selectBox = append(selectBox, startSelect)

	for _, mt := range momentTypes {
		var option string
		if mt.ID == selectedID {
			option = fmt.Sprintf("<option value='%d' selected>%s</option>", mt.ID, mt.Name)
		} else {
			option = fmt.Sprintf("<option value='%d'>%s</option>", mt.ID, mt.Name)
		}
		selectBox = append(selectBox, option)
	}

	endSelect := fmt.Sprintf("</select>")
	selectBox = append(selectBox, endSelect)

	return template.HTML(strings.Join(selectBox, "\n"))
}

func (hm *HistoricalMoment) StartLabel() template.HTML {
	return BuildLabel("inputStart", "Fecha de comienzo")
}

func (hm *HistoricalMoment) StartInput(message, value string) template.HTML {
	return BuildTextInput("inputStart", value, message)
}

func (hm *HistoricalMoment) EndLabel() template.HTML {
	return BuildLabel("inputEnd", "Fecha de término")
}

func (hm *HistoricalMoment) EndInput(message, value string) template.HTML {
	return BuildTextInput("inputEnd", value, message)
}

func (hm *HistoricalMoment) DescriptionLabel() template.HTML {
	return BuildLabel("inputDescription", "Descripción")
}

func (hm *HistoricalMoment) DescriptionInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputDescription", value, message)
}

func (hm *HistoricalMoment) SummaryLabel() template.HTML {
	return BuildLabel("inputSummary", "Resumen")
}

func (hm *HistoricalMoment) SummaryInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputSummary", value, message)
}



type HistoricalMoments struct {
	HistoricalMoments []*HistoricalMoment `json:"historical_moments"`
	RecordCount       int                 `json:"record_count"`
}

type HistoricalMomentRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	SchemaID    int    `json:"schema_id"`
	TypeID      int    `json:"moment_type_id"`
	Start       string `json:"start_date"`
	End         string `json:"end_date"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

