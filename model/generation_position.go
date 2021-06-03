package model

type GenerationPosition struct {
	ID         int         `json:"id"`
	MomentID   int         `json:"moment_id"`
	Name       string      `json:"name"`
	Ordinal    int         `json:"ordinal"`
	LifePhase  *LifePhase  `json:"life_phase"`
	Generation *Generation `json:"generation"`
}

type GenerationFullPosition struct {
	GenerationID     int                 `json:"generation_id"`
	Position         *GenerationPosition `json:"position"`
	HistoricalMoment *HistoricalMoment   `json:"historical_moment"`
}
