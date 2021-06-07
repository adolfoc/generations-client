package model

type GroupType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GroupTypes struct {
	GroupTypes  []*GroupType `json:"group_types"`
	RecordCount int          `json:"record_count"`
}

type GroupTypeRequest struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
