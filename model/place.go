package model

type Place struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	PlaceTypeID   int        `json:"place_type_id"`
	ParentPlaceID int        `json:"parent_place_id"`
	Start         string     `json:"start_date"`
	End           string     `json:"end_date"`
	Summary       string     `json:"summary"`
	Description   string     `json:"description"`
	PlaceType     *PlaceType `json:"place_type"`
	PlaceName     string     `json:"place_name"`
}

type Places struct {
	Places      []*Place `json:"places"`
	RecordCount int      `json:"record_count"`
}

