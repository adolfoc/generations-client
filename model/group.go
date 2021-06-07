package model

type Group struct {
	ID            int                   `json:"id"`
	Name          string                `json:"name"`
	GroupTypeID   int                   `json:"group_type_id"`
	ParentGroupID int                   `json:"parent_group_id"`
	Start         string                `json:"start_date"`
	End           string                `json:"end_date"`
	Summary       string                `json:"summary"`
	Members       []*GroupMembershipMap `json:"members"`
	Description   string                `json:"description"`
	GroupType     *GroupType            `json:"group_type"`
	GroupName     string                `json:"group_name"`
}

type Groups struct {
	Groups      []*Group `json:"groups"`
	RecordCount int      `json:"record_count"`
}

type GroupRequest struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	GroupTypeID   int    `json:"group_type_id"`
	ParentGroupID int    `json:"parent_group_id"`
	Start         string `json:"start_date"`
	End           string `json:"end_date"`
	Summary       string `json:"summary"`
	Description   string `json:"description"`
}
