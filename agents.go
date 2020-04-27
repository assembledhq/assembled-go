// Code generated by riza; DO NOT EDIT.

package assembled

import (
	"context"
	"fmt"
)

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

type ListAgentsRequest struct {
	Channels []string `json:"channels,omitempty"` // One of: 'phone', 'email', or 'chat'.
	Queue    string   `json:"queue,omitempty"`    // Name of the queue to filter on.
	Site     string   `json:"site,omitempty"`     // Name of the site to filter on.
	Team     string   `json:"team,omitempty"`     // Name of the team to filter on.
}

// Returns ListAgentsRequest with `url` tags instead of `json` tags for URL encoding.
func (r *ListAgentsRequest) params() interface{} {
	if r == nil {
		return r
	}
	type params struct {
		Channels []string `url:"channels,omitempty"`
		Queue    string   `url:"queue,omitempty"`
		Site     string   `url:"site,omitempty"`
		Team     string   `url:"team,omitempty"`
	}
	p := params(*r)
	return &p
}

type ListAgentsResponse struct {
	Agents map[string]Agent `json:"agents,omitempty"`
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

// Returns UpdateAgentRequest with ID set to the empty string so that it's
// not included in the JSON request body.
func (r *UpdateAgentRequest) body() interface{} {
	if r == nil {
		return r
	}
	req := *r
	req.ID = ""
	return &req
}

// Creates an agent profile with the specified parameters. Valid IDs for site,
// teams, and queues can be retrieved from endpoints for filters. This
// endpoint will return 400 if invalid IDs are provided.
func (c *Client) CreateAgent(ctx context.Context, r *CreateAgentRequest) (*Agent, error) {
	var resp Agent
	if err := c.request(ctx, "POST", "/v0/agents", nil, r, &resp); err != nil {
		return nil, fmt.Errorf("CreateAgent: %w", err)
	}
	return &resp, nil
}

// Returns a list of agent objects that match the provided query.
func (c *Client) ListAgents(ctx context.Context, r *ListAgentsRequest) (*ListAgentsResponse, error) {
	var resp ListAgentsResponse
	if err := c.request(ctx, "GET", "/v0/agents", r.params(), nil, &resp); err != nil {
		return nil, fmt.Errorf("ListAgents: %w", err)
	}
	return &resp, nil
}

// Partial update of an agent with the specified fields. Fields that are not
// included in the request are not updated, while fields that are explicitly
// set to null or an appropriate empty value (for example, "[]" for lists) are
// set to empty. Valid IDs for site, teams, and queues can be retrieved from
// endpoints for filters. This endpoint will return 400 if invalid filter IDs
// are provided.
func (c *Client) UpdateAgent(ctx context.Context, r *UpdateAgentRequest) (*Agent, error) {
	var resp Agent
	if err := c.request(ctx, "PATCH", fmt.Sprintf("/v0/agents/%s", r.ID), nil, r.body(), &resp); err != nil {
		return nil, fmt.Errorf("UpdateAgent: %w", err)
	}
	return &resp, nil
}