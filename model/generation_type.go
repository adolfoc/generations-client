package model

type GenerationType struct {
	ID          int    `json:"id"`
	SchemaID    int    `json:"schema_id"`
	Archetype   string `json:"archetype"`
	Description string `json:"description"`
}

