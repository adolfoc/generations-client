package model

type GroupMembershipMap struct {
	ID        int     `json:"id"`
	GroupID   int     `json:"group_id"`
	PersonID  int     `json:"person_id"`
	Role      string  `json:"role"`
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date"`
	Person    *Person `json:"person"`
}

type GroupMembershipRequest struct {
	ID        int    `json:"id"`
	GroupID   int    `json:"group_id"`
	PersonID  int    `json:"person_id"`
	Role      string `json:"role"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

