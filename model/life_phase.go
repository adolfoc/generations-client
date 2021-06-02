package model

type LifePhase struct {
	ID        int    `json:"id"`
	SchemaID  int    `json:"schema_id"`
	Name      string `json:"name"`
	StartYear int    `json:"start_year"`
	EndYear   int    `json:"end_year"`
	Role      string `json:"role"`
}

