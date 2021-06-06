package model

import (
	"fmt"
	"html/template"
	"strings"
)

type Place struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	PlaceTypeID   int        `json:"place_type_id"`
	ParentPlaceID int        `json:"parent_place_id"`
	Start         string     `json:"start_date"`
	End           string     `json:"end_date"`
	Summary       string     `json:"summary"`
	Description   string     `json:"description"`
	PlaceType     *PlaceType `json:"place_type"`
	ParentPlace   *Place     `json:"parent_place"`
	PlaceName     string     `json:"place_name"`
}

func (p *Place) IDInput(message, value string) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (p *Place) NameLabel() template.HTML {
	return BuildLabel("inputName", "Nombre")
}

func (p *Place) NameInput(message, value string) template.HTML {
	return BuildTextInput("inputName", value, message)
}

func (p *Place) PlaceTypeIDLabel() template.HTML {
	return BuildLabel("inputPlaceTypeID", "Tipo de lugar")
}

func (p *Place) PlaceTypeIDSelectBox(placeTypes []*PlaceType, selectedID int) template.HTML {
	var selectBox []string
	startSelect := fmt.Sprintf("<select class='form-select library-control' id='inputPlaceTypeID' name='inputPlaceTypeID'>")
	selectBox = append(selectBox, startSelect)

	blankOption := "<option value='' selected>Elija el tipo de lugar...</option>"
	selectBox = append(selectBox, blankOption)

	for _, et := range placeTypes {
		var option string
		if et.ID == selectedID {
			option = fmt.Sprintf("<option value='%d' selected>%s</option>", et.ID, et.Name)
		} else {
			option = fmt.Sprintf("<option value='%d'>%s</option>", et.ID, et.Name)
		}
		selectBox = append(selectBox, option)
	}

	endSelect := fmt.Sprintf("</select>")
	selectBox = append(selectBox, endSelect)

	return template.HTML(strings.Join(selectBox, "\n"))
}

func (p *Place) ParentPlaceIDLabel() template.HTML {
	return BuildLabel("inputParentPlaceID", "Pertenece a")
}

func (p *Place) ParentPlaceIDSelectBox(places []*Place, selectedID int) template.HTML {
	var selectBox []string
	startSelect := fmt.Sprintf("<select class='form-select library-control' id='inputParentPlaceID' name='inputParentPlaceID'>")
	selectBox = append(selectBox, startSelect)

	blankOption := "<option value='' selected disabled hidden>Elija el lugar del cual es parte...</option>"
	selectBox = append(selectBox, blankOption)

	emptyOption := "<option value='' selected></option>"
	selectBox = append(selectBox, emptyOption)

	for _, et := range places {
		var option string
		if et.ID == selectedID {
			option = fmt.Sprintf("<option value='%d' selected>%s</option>", et.ID, et.Name)
		} else {
			option = fmt.Sprintf("<option value='%d'>%s</option>", et.ID, et.Name)
		}
		selectBox = append(selectBox, option)
	}

	endSelect := fmt.Sprintf("</select>")
	selectBox = append(selectBox, endSelect)

	return template.HTML(strings.Join(selectBox, "\n"))
}

func (p *Place) StartLabel() template.HTML {
	return BuildLabel("inputStart", "Fecha de comienzo")
}

func (p *Place) StartInput(message, value string) template.HTML {
	return BuildTextInput("inputStart", value, message)
}

func (p *Place) EndLabel() template.HTML {
	return BuildLabel("inputEnd", "Fecha de término")
}

func (p *Place) EndInput(message, value string) template.HTML {
	return BuildTextInput("inputEnd", value, message)
}

func (p *Place) SummaryLabel() template.HTML {
	return BuildLabel("inputSummary", "Resumen")
}

func (p *Place) SummaryInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputSummary", value, message)
}

func (p *Place) DescriptionLabel() template.HTML {
	return BuildLabel("inputDescription", "Descripción")
}

func (p *Place) DescriptionInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputDescription", value, message)
}

type Places struct {
	Places      []*Place `json:"places"`
	RecordCount int      `json:"record_count"`
}

type PlaceRequest struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	PlaceTypeID   int    `json:"place_type_id"`
	ParentPlaceID int    `json:"parent_place_id"`
	Start         string `json:"start_date"`
	End           string `json:"end_date"`
	Summary       string `json:"summary"`
	Description   string `json:"description"`
}
