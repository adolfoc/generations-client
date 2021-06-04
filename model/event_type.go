package model

type EventType struct {
	ID          int    `json:"id"`
	IsNatural   bool   `json:"is_natural"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

