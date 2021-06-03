package model

type Intangible struct {
	ID          int    `json:"id"`
	LandscapeID int    `json:"landscape_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

