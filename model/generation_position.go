package model

import (
	"fmt"
	"html/template"
	"strings"
)

type GenerationPosition struct {
	ID         int         `json:"id"`
	MomentID   int         `json:"moment_id"`
	Name       string      `json:"name"`
	Ordinal    int         `json:"ordinal"`
	LifePhase  *LifePhase  `json:"life_phase"`
	Generation *Generation `json:"generation"`
}

func (gp *GenerationPosition) IDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (gp *GenerationPosition) GenerationIDInput(message string, value int) template.HTML {
	return BuildHiddenIDInput("inputGenerationID", value)
}

func (gp *GenerationPosition) NameLabel() template.HTML {
	return BuildLabel("inputName", "Nombre")
}

func (gp *GenerationPosition) NameInput(message, value string) template.HTML {
	return BuildTextInput("inputName", value, message)
}

func (gp *GenerationPosition) OrdinalLabel() template.HTML {
	return BuildLabel("inputOrdinal", "Ordinal")
}

func (gp *GenerationPosition) OrdinalInput(message string, value int) template.HTML {
	return BuildIntegerInput("inputOrdinal", value, message)
}

func (gp *GenerationPosition) LifePhaseIDLabel() template.HTML {
	return BuildLabel("inputLifePhaseID", "Etapa de vida")
}

func (gp *GenerationPosition) LifePhaseIDSelectBox(lifePhases []*LifePhase, selectedID int) template.HTML {
	var selectBox []string
	startSelect := fmt.Sprintf("<select class='form-select library-control' id='inputLifePhaseID' name='inputLifePhaseID'>")
	selectBox = append(selectBox, startSelect)

	blankOption := "<option value='' selected disabled hidden>Elija una etapa de vida...</option>"
	selectBox = append(selectBox, blankOption)

	for _, lp := range lifePhases {
		var option string
		if lp.ID == selectedID {
			option = fmt.Sprintf("<option value='%d' selected>%s</option>", lp.ID, lp.Name)
		} else {
			option = fmt.Sprintf("<option value='%d'>%s</option>", lp.ID, lp.Name)
		}
		selectBox = append(selectBox, option)
	}

	endSelect := fmt.Sprintf("</select>")
	selectBox = append(selectBox, endSelect)

	return template.HTML(strings.Join(selectBox, "\n"))
}


func (gp *GenerationPosition) MomentIDLabel() template.HTML {
	return BuildLabel("inputMomentID", "Momento")
}

func (gp *GenerationPosition) MomentIDSelectBox(moments []*HistoricalMoment, selectedID int) template.HTML {
	var selectBox []string
	startSelect := fmt.Sprintf("<select class='form-select library-control' id='inputMomentID' name='inputMomentID'>")
	selectBox = append(selectBox, startSelect)

	blankOption := "<option value='' selected disabled hidden>Elija un momento histórico...</option>"
	selectBox = append(selectBox, blankOption)

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

type GenerationPositionRequest struct {
	ID           int    `json:"id"`
	MomentID     int    `json:"moment_id"`
	Name         string `json:"name"`
	Ordinal      int    `json:"ordinal"`
	LifePhaseID  int    `json:"life_phase_id"`
	GenerationID int    `json:"generation_id"`
}

type GenerationFullPosition struct {
	GenerationID     int                 `json:"generation_id"`
	Position         *GenerationPosition `json:"position"`
	HistoricalMoment *HistoricalMoment   `json:"historical_moment"`
}
