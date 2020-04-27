package assembled

// Filters can be used to categorize or group agents as well as activities.
type Filter struct {
	ID string `json:"id,omitempty"`

	// Identifier for the parent filter. May be null.
	ParentID string `json:"parent_id,omitempty"`

	CreatedAt Timestamp `json:"created_at,omitempty"` // Time of creation, in Unix time.
	Name      string    `json:"name,omitempty"`       // Identifier for the filter.
	UpdatedAt Timestamp `json:"updated_at,omitempty"` // Time of last update, in Unix time.
}

type FilterUpdateRequest struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	ParentID string `json:"parent_id,omitempty"`
}
