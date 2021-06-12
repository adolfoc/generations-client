package model

import (
	"fmt"
	"html/template"
	"strings"
)

type GenerationSchema struct {
	ID                    int               `json:"id"`
	Name                  string            `json:"name"`
	Description           string            `json:"description"`
	StartYear             int               `json:"start_year"`
	EndYear               int               `json:"end_year"`
	MinimumGenerationSpan int               `json:"minimum_generation_span"`
	MaximumGenerationSpan int               `json:"maximum_generation_span"`
	GenerationalTypes     []*GenerationType `json:"generation_types"`
	MomentTypes           []*MomentType     `json:"moment_types"`
	Place                 *Place            `json:"place"`
}

func (gs *GenerationSchema) MakeStudyTitle() string {
	return fmt.Sprintf("%s (%d-%d)", gs.Name, gs.StartYear, gs.EndYear)
}

func (gs *GenerationSchema) GenerationalSpan() template.HTML {
	return template.HTML(fmt.Sprintf("%d&mdash;%d", gs.MinimumGenerationSpan, gs.MaximumGenerationSpan))
}

func (gs *GenerationSchema) DurationSpan() template.HTML {
	extension := gs.EndYear - gs.StartYear
	return template.HTML(fmt.Sprintf("%d&mdash;%d (%d años)", gs.StartYear, gs.EndYear, extension))
}

func (gs *GenerationSchema) IDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (gs *GenerationSchema) NameLabel() template.HTML {
	return BuildLabel("inputName", "Nombre")
}

func (gs *GenerationSchema) NameInput(message, value string) template.HTML {
	return BuildTextInput("inputName", value, message)
}

func (gs *GenerationSchema) DescriptionLabel() template.HTML {
	return BuildLabel("inputDescription", "Descripción")
}

func (gs *GenerationSchema) DescriptionInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputDescription", value, message)
}

func (gs *GenerationSchema) StartYearLabel() template.HTML {
	return BuildLabel("inputStartYear", "Año de inicio")
}

func (gs *GenerationSchema) StartYearInput(message string, value int) template.HTML {
	return BuildIntegerInput("inputStartYear", value, message)
}

func (gs *GenerationSchema) EndYearLabel() template.HTML {
	return BuildLabel("inputEndYear", "Año de término")
}

func (gs *GenerationSchema) EndYearInput(message string, value int) template.HTML {
	return BuildIntegerInput("inputEndYear", value, message)
}

func (gs *GenerationSchema) MinimumGenerationSpanLabel() template.HTML {
	return BuildLabel("inputMinimumGenerationSpan", "Espacio mínimo entre generaciones")
}

func (gs *GenerationSchema) MinimumGenerationSpanInput(message string, value int) template.HTML {
	return BuildIntegerInput("inputMinimumGenerationSpan", value, message)
}

func (gs *GenerationSchema) MaximumGenerationSpanLabel() template.HTML {
	return BuildLabel("inputMaximumGenerationSpan", "Espacio máximo entre generaciones")
}

func (gs *GenerationSchema) MaximumGenerationSpanInput(message string, value int) template.HTML {
	return BuildIntegerInput("inputMaximumGenerationSpan", value, message)
}

func (gs *GenerationSchema) PlaceLabel() template.HTML {
	return BuildLabel("inputPlaceID", "Lugar")
}

func (gs *GenerationSchema) PlaceSelectBox(places []*Place, selectedID int) template.HTML {
	var selectBox []string
	startSelect := fmt.Sprintf("<select class='form-select library-control' id='inputPlaceID' name='inputPlaceID'>")
	selectBox = append(selectBox, startSelect)

	blankOption := "<option value='' selected disabled hidden>Elija un lugar...</option>"
	selectBox = append(selectBox, blankOption)

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

type GenerationSchemas struct {
	GenerationSchemas []*GenerationSchema `json:"generation_schemas"`
	RecordCount       int                 `json:"record_count"`
}

type GenerationSchemaRequest struct {
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	StartYear             int    `json:"start_year"`
	EndYear               int    `json:"end_year"`
	MinimumGenerationSpan int    `json:"minimum_generation_span"`
	MaximumGenerationSpan int    `json:"maximum_generation_span"`
	PlaceID               int    `json:"place_id"`
}

type SchemaComparativePosition struct {
	Role       string            `json:"role"`
	LifePhase  *LifePhase        `json:"life_phase"`
	Generation *Generation       `json:"generation"`
	Landscape  *HistoricalMoment `json:"landscape"`
}

type SchemaComparativeItem struct {
	Moment    *HistoricalMoment            `json:"historical_moment"`
	Positions []*SchemaComparativePosition `json:"schema_comparative_position"`
}

type SchemaComparative struct {
	Schema *GenerationSchema        `json:"generation_schema"`
	Items  []*SchemaComparativeItem `json:"comparative_items"`
}

