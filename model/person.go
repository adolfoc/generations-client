package model

import (
	"fmt"
	"html/template"
	"strings"
)

type Person struct {
	ID           int            `json:"id"`
	Names        string         `json:"names"`
	Aliases      []string       `json:"aliases"`
	KnownAs      []string       `json:"known_as"`
	Sex          int            `json:"sex"`
	BirthDate    string         `json:"birth_date"`
	BirthPlace   *Place         `json:"birth_place"`
	DeathDate    string         `json:"death_date"`
	DeathPlace   *Place         `json:"death_place"`
	Summary      string         `json:"summary"`
	LifeSegments []*LifeSegment `json:"life_segments"`
}

func (p *Person) AliasHTML() string {
	return strings.Join(p.Aliases, ", ")
}

func (p *Person) KnownAsHTML() string {
	return strings.Join(p.KnownAs, ", ")
}

func (p *Person) Lifespan() string {
	from := ""
	if len(p.BirthDate) > 0 {
		from += p.BirthDate
	}

	if p.BirthPlace != nil && p.BirthPlace.ID > 0 {
		from += " (" + p.BirthPlace.Name + ")"
	}

	to := ""
	if len(p.DeathDate) > 0 {
		to += p.DeathDate
	}

	if p.DeathPlace != nil && p.DeathPlace.ID > 0 {
		to += " (" + p.DeathPlace.Name + ")"
	}

	return fmt.Sprintf("%s--%s", from, to)
}

func (p *Person) IDInput(message, value string) template.HTML {
	return BuildHiddenIDInput("inputID", value)
}

func (p *Person) NamesLabel() template.HTML {
	return BuildLabel("inputNames", "Nombres")
}

func (p *Person) NamesInput(message, value string) template.HTML {
	return BuildTextInput("inputNames", value, message)
}

func (p *Person) AliasesLabel() template.HTML {
	return BuildLabel("inputAliases", "Alias")
}

func (p *Person) AliasesInput(message, value string) template.HTML {
	return BuildTextInput("inputAliases", value, message)
}

func (p *Person) KnownAsLabel() template.HTML {
	return BuildLabel("inputKnownAs", "Conocido como")
}

func (p *Person) KnownAsInput(message, value string) template.HTML {
	return BuildTextInput("inputKnownAs", value, message)
}

func (p *Person) BirthDateLabel() template.HTML {
	return BuildLabel("inputBirthDate", "Fecha de nacimiento")
}

func (p *Person) BirthDateInput(message, value string) template.HTML {
	return BuildTextInput("inputBirthDate", value, message)
}

func (p *Person) BirthPlaceLabel() template.HTML {
	return BuildLabel("inputBirthPlaceID", "Lugar de nacimiento")
}

func (p *Person) BirthPlaceSelectBox(places []*Place, selectedID int) template.HTML {
	var selectBox []string
	startSelect := fmt.Sprintf("<select class='form-select library-control' id='inputBirthPlaceID' name='inputBirthPlaceID'>")
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

func (p *Person) DeathDateLabel() template.HTML {
	return BuildLabel("inputDeathDate", "Fecha de fallecimiento")
}

func (p *Person) DeathDateInput(message, value string) template.HTML {
	return BuildTextInput("inputDeathDate", value, message)
}

func (p *Person) DeathPlaceLabel() template.HTML {
	return BuildLabel("inputDeathPlaceID", "Lugar de fallecimiento")
}

func (p *Person) DeathPlaceSelectBox(places []*Place, selectedID int) template.HTML {
	var selectBox []string
	startSelect := fmt.Sprintf("<select class='form-select library-control' id='inputDeathPlaceID' name='inputDeathPlaceID'>")
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

func (p *Person) SummaryLabel() template.HTML {
	return BuildLabel("inputSummary", "Resumen")
}

func (p *Person) SummaryInput(message, value string) template.HTML {
	return BuildTextAreaInput("inputSummary", value, message)
}



type Persons struct {
	Persons     []*Person `json:"persons"`
	RecordCount int       `json:"record_count"`
}

type PersonRequest struct {
	ID           int      `json:"id"`
	Names        string   `json:"names"`
	Aliases      []string `json:"aliases"`
	KnownAs      []string `json:"known_as"`
	Sex          int      `json:"sex"`
	BirthDate    string   `json:"birth_date"`
	BirthPlaceID int      `json:"birth_place_id"`
	DeathDate    string   `json:"death_date"`
	DeathPlaceID int      `json:"death_place_id"`
	Summary      string   `json:"summary"`
}
