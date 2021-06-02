package model

type MomentType struct {
	ID          int    `json:"id"`
	SchemaID    int    `json:"schema_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

