package model

type GenerationalLandscape struct {
	ID                int           `json:"id"`
	GenerationID      int           `json:"generation_id"`
	FormationMomentID int           `json:"formation_moment_id"`
	Description       string        `json:"description"`
	Tangibles         []*Tangible   `json:"tangibles"`
	Intangibles       []*Intangible `json:"intangibles"`
}
