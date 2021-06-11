package model

import (
	"fmt"
	"html/template"
	"strings"
)

type Generation struct {
	ID                   int             `json:"id"`
	Name                 string          `json:"name"`
	SchemaID             int             `json:"schema_id"`
	Type                 *GenerationType `json:"generation_type"`
	StartYear            int             `json:"start_year"`
	EndYear              int             `json:"end_year"`
	Place                *Place          `json:"place"`
	Description          string          `json:"description"`
	DescriptionHTML      string          `json:"description_html"`
	FormationLandscapeID int             `json:"formation_landscape_id"`
}

func (g *Generation) Span() template.HTML {
	return template.HTML(fmt.Sprintf("%d&mdash;%d", g.StartYear, g.EndYear))
}

func (g *Generation) IDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (g *Generation) SchemaIDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputSchemaID", value)
}

func (g *Generation) NameLabel() template.HTML {
	return BuildLabel("inputName", "Nombre")
}

func (g *Generation) NameInput(message, value string) template.HTML {
	return BuildTextInput("inputName", value, message)
}

func (g *Generation) TypeLabel() template.HTML {
	return BuildLabel("inputTypeID", "Tipo")
}

func (g *Generation) TypeSelectBox(generationTypes []*GenerationType, selectedID int) template.HTML {
	var selectBox []string
	startSelect := fmt.Sprintf("<select class='form-select library-control' id='inputTypeID' name='inputTypeID'>")
	selectBox = append(selectBox, startSelect)

	for _, gt := range generationTypes {
		var option string
		if gt.ID == selectedID {
			option = fmt.Sprintf("<option value='%d' selected>%s</option>", gt.ID, gt.Archetype)
		} else {
			option = fmt.Sprintf("<option value='%d'>%s</option>", gt.ID, gt.Archetype)
		}
		selectBox = append(selectBox, option)
	}

	endSelect := fmt.Sprintf("</select>")
	selectBox = append(selectBox, endSelect)

	return template.HTML(strings.Join(selectBox, "\n"))
}

func (g *Generation) StartYearLabel() template.HTML {
	return BuildLabel("inputStartYear", "Año de comienzo")
}

func (g *Generation) StartYearInput(message string, value int) template.HTML {
	return BuildIntegerInput("inputStartYear", value, message)
}

func (g *Generation) EndYearLabel() template.HTML {
	return BuildLabel("inputEndYear", "Año de término")
}

func (g *Generation) EndYearInput(message string, value int) template.HTML {
	return BuildIntegerInput("inputEndYear", value, message)
}

func (g *Generation) PlaceLabel() template.HTML {
	return BuildLabel("inputPlaceID", "Lugar")
}

func (g *Generation) PlaceSelectBox(places []*Place, selectedID int) template.HTML {
	var selectBox []string
	startSelect := fmt.Sprintf("<select class='form-select library-control' id='inputPlaceID' name='inputPlaceID'>")
	selectBox = append(selectBox, startSelect)

	for _, place := range places {
		var option string
		if place.ID == selectedID {
			option = fmt.Sprintf("<option value='%d' selected>%s</option>", place.ID, place.Name)
		} else {
			option = fmt.Sprintf("<option value='%d'>%s</option>", place.ID, place.Name)
		}
		selectBox = append(selectBox, option)
	}

	endSelect := fmt.Sprintf("</select>")
	selectBox = append(selectBox, endSelect)

	return template.HTML(strings.Join(selectBox, "\n"))
}

func (g *Generation) DescriptionLabel() template.HTML {
	return BuildLabel("inputDescription", "Descripción")
}

func (g *Generation) DescriptionInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputDescription", value, message)
}

func (g *Generation) LandscapeLabel() template.HTML {
	return BuildLabel("inputLandscapeID", "Paisaje de formación")
}

func (g *Generation) LandscapeSelectBox(moments []*HistoricalMoment, selectedID int) template.HTML {
	var selectBox []string
	startSelect := fmt.Sprintf("<select class='form-select library-control' id='inputLandscapeID' name='inputLandscapeID'>")
	selectBox = append(selectBox, startSelect)

	for _, moment := range moments {
		var option string
		if moment.ID == selectedID {
			option = fmt.Sprintf("<option value='%d' selected>%s</option>", moment.ID, moment.Name)
		} else {
			option = fmt.Sprintf("<option value='%d'>%s</option>", moment.ID, moment.Name)
		}
		selectBox = append(selectBox, option)
	}

	endSelect := fmt.Sprintf("</select>")
	selectBox = append(selectBox, endSelect)

	return template.HTML(strings.Join(selectBox, "\n"))
}




type Generations struct {
	Generations []*Generation `json:"generations"`
	RecordCount int           `json:"record_count"`
}

type GenerationRequest struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	SchemaID             int    `json:"schema_id"`
	TypeID               int    `json:"generation_type_id"`
	StartYear            int    `json:"start_year"`
	EndYear              int    `json:"end_year"`
	PlaceID              int    `json:"place_id"`
	FormationLandscapeID int    `json:"formation_landscape_id"`
	Description          string `json:"description"`
}

