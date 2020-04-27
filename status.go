package assembled

import (
	"context"
	"fmt"
)

// Agents can have statuses (sourced from external systems) that indicate what
// they are currently working on. For example, in a phone system an agent may
// be “available for calls”, “busy on a call”, or “wrapping up on a
// call".
type AgentStatus struct {
	// Agent's current status from upstream system e.g. 'ready', 'away', or
	// 'busy'. Values are not validated.
	Status string `json:"status,omitempty"`

	// Identifier for the associated event with the current status.
	EventID string `json:"event_id,omitempty"`

	AgentID   string    `json:"agent_id,omitempty"` // Identifier for the corresponding agent.
	AgentName string    `json:"agent_name,omitempty"`
	Channel   string    `json:"channel,omitempty"`
	EndTime   Timestamp `json:"end_time,omitempty"`
	StartTime Timestamp `json:"start_time,omitempty"`
}

type CreateAgentStatusRequest struct {
	// Agent's current status from upstream system e.g. 'ready', 'away', or
	// 'busy'. Values are not validated.
	Status string `json:"status,omitempty"`

	// Identifier for the associated event with the current status.
	EventID string `json:"event_id,omitempty"`

	AgentID   string    `json:"agent_id,omitempty"` // Identifier for the corresponding agent.
	AgentName string    `json:"agent_name,omitempty"`
	Channel   string    `json:"channel,omitempty"`
	EndTime   Timestamp `json:"end_time,omitempty"`
	StartTime Timestamp `json:"start_time,omitempty"`
}

func (c *Client) CreateAgentStatus(ctx context.Context, r *CreateAgentStatusRequest) (*AgentStatus, error) {
	var resp AgentStatus
	if err := c.request(ctx, "POST", "/v0/agents/status", r, &resp); err != nil {
		return nil, fmt.Errorf("CreateAgentStatus: %w", err)
	}
	return &resp, nil
}

func (c *Client) GetAgentStatus(ctx context.Context, r *Agent) (*AgentStatus, error) {
	var resp AgentStatus
	if err := c.request(ctx, "GET", fmt.Sprintf("/v0/agents/%s/status", r.ID), r, &resp); err != nil {
		return nil, fmt.Errorf("GetAgentStatus: %w", err)
	}
	return &resp, nil
}
