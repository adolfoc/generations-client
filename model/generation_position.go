package model

type GenerationPosition struct {
	ID         int         `json:"id"`
	MomentID   int         `json:"moment_id"`
	Name       string      `json:"name"`
	Ordinal    int         `json:"ordinal"`
	LifePhase  *LifePhase  `json:"life_phase"`
	Generation *Generation `json:"generation"`
}

