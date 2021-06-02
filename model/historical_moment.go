package model

import "fmt"

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

func (hm *HistoricalMoment) Span() string {
	return fmt.Sprintf("%s--%s", hm.Start, hm.End)
}

type HistoricalMoments struct {
	HistoricalMoments []*HistoricalMoment `json:"historical_moments"`
	RecordCount       int                 `json:"record_count"`
}

