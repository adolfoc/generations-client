package model

import "fmt"

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

func (gs *GenerationSchema) Span() string {
	return fmt.Sprintf("%d--%d", gs.MinimumGenerationSpan, gs.MaximumGenerationSpan)
}

type GenerationSchemas struct {
	GenerationSchemas []*GenerationSchema `json:"generation_schemas"`
	RecordCount       int                 `json:"record_count"`
}
