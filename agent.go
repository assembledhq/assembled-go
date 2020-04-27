package assembled

import (
	"context"
	"fmt"
)

// Agents handle units of work and can be assigned to activities or schedules.
// They can be grouped by site and team and can be assigned to queues or
// channels based on their skills.
type Agent struct {
	ID string `json:"id,omitempty"`

	// Third-party identifier. Supplied to Assembled and is used to uniquely
	// identify agents across different systems.
	ImportID string `json:"import_id,omitempty"`

	Channels []string `json:"channels,omitempty"` // One of: 'phone', 'email', or 'chat'.
	Email    string   `json:"email,omitempty"`
	Name     string   `json:"name,omitempty"`
	Queues   []string `json:"queues,omitempty"` // Unique identifiers for associated queues.
	Site     string   `json:"site,omitempty"`   // Unique identifier for associated site.
	Skills   []string `json:"skills,omitempty"` // Unique identifiers for associated skills.
	Teams    []string `json:"teams,omitempty"`  // Unique identifiers for associated teams.
}

type CreateAgentRequest struct {
	// Third-party identifier. Supplied to Assembled and is used to uniquely
	// identify agents across different systems.
	ImportID string `json:"import_id,omitempty"`

	Channels []string `json:"channels,omitempty"` // One of: 'phone', 'email', or 'chat'.
	Email    string   `json:"email,omitempty"`
	Name     string   `json:"name,omitempty"`
	Queues   []string `json:"queues,omitempty"` // Unique identifiers for associated queues.
	Site     string   `json:"site,omitempty"`   // Unique identifier for associated site.
	Skills   []string `json:"skills,omitempty"` // Unique identifiers for associated skills.
	Teams    []string `json:"teams,omitempty"`  // Unique identifiers for associated teams.
}

// Creates an agent profile with the specified parameters. Valid IDs for site,
// teams, and queues can be retrieved from endpoints for filters. This
// endpoint will return 400 if invalid IDs are provided.
func (c *Client) CreateAgent(ctx context.Context, r *CreateAgentRequest) (*Agent, error) {
	var resp Agent
	if err := c.request(ctx, "POST", "/v0/agents", r, &resp); err != nil {
		return nil, fmt.Errorf("CreateAgent: %w", err)
	}
	return &resp, nil
}

type UpdateAgentRequest struct {
	ID string `json:"id,omitempty"`

	// Third-party identifier. Supplied to Assembled and is used to uniquely
	// identify agents across different systems.
	ImportID string `json:"import_id,omitempty"`

	Channels []string `json:"channels,omitempty"` // One of: 'phone', 'email', or 'chat'.
	Email    string   `json:"email,omitempty"`
	Name     string   `json:"name,omitempty"`
	Queues   []string `json:"queues,omitempty"` // Unique identifiers for associated queues.
	Site     string   `json:"site,omitempty"`   // Unique identifier for associated site.
	Skills   []string `json:"skills,omitempty"` // Unique identifiers for associated skills.
	Teams    []string `json:"teams,omitempty"`  // Unique identifiers for associated teams.
}

// Partial update of an agent with the specified fields. Fields that are not
// included in the request are not updated, while fields that are explicitly
// set to null or an appropriate empty value (for example, "[]" for lists) are
// set to empty. Valid IDs for site, teams, and queues can be retrieved from
// endpoints for filters. This endpoint will return 400 if invalid filter IDs
// are provided.
func (c *Client) UpdateAgent(ctx context.Context, r *UpdateAgentRequest) (*Agent, error) {
	var resp Agent
	if err := c.request(ctx, "PATCH", fmt.Sprintf("/v0/agents/%s", r.ID), r, &resp); err != nil {
		return nil, fmt.Errorf("UpdateAgent: %w", err)
	}
	return &resp, nil
}

type ListAgentsRequest struct {
	Channels []string `url:"channels,omitempty"` // One of: 'phone', 'email', or 'chat'.
	Queue    string   `url:"queue,omitempty"`    // Name of the queue to filter on.
	Site     string   `url:"site,omitempty"`     // Name of the site to filter on.
	Team     string   `url:"team,omitempty"`     // Name of the team to filter on.
}

type ListAgentsResponse struct {
	Agents map[string]Agent `json:"agents,omitempty"`
}

func (c *Client) ListAgents(ctx context.Context, r *ListAgentsRequest) (*ListAgentsResponse, error) {
	var resp ListAgentsResponse
	if err := c.request(ctx, "GET", "/v0/agents", r, &resp); err != nil {
		return nil, fmt.Errorf("ListAgents: %w", err)
	}
	return &resp, nil
}
