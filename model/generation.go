package model

import "fmt"

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

func (g *Generation) Span() string {
	return fmt.Sprintf("%d--%d", g.StartYear, g.EndYear)
}

type Generations struct {
	Generations []*Generation `json:"generations"`
	RecordCount int           `json:"record_count"`
}

