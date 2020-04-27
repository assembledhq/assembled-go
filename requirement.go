package assembled

import (
	"context"
	"fmt"
)

// Requirements represent a set of staffing requirements and grouped staffing
// against the requirement within a given time window. They correspond to
// metrics about required staffing and scheduled staffing that appear above
// the team calendar.
type Requirement struct {
	// Count of scheduled staffing in the interval, can be partial.
	Scheduled float64 `json:"scheduled,omitempty"`

	// Count of required staffing in the interval, can be partial.
	Required float64 `json:"required,omitempty"`

	// Unique identifier for the corresponding requirement type.
	RequirementTypeID string `json:"requirement_type_id,omitempty"`

	EndTime   Timestamp `json:"end_time,omitempty"`
	StartTime Timestamp `json:"start_time,omitempty"`
}

type CreateRequirementRequest struct {
	// Must be a multiple of the configured interval as above. The difference
	// with start_time must also equal the interval.
	EndTime Timestamp `json:"end_time,omitempty"`

	// Must be a multiple of the configured interval e.g. 3600 for 1 hour or
	// 900 for 15 minutes.
	StartTime Timestamp `json:"start_time,omitempty"`

	Required          float64 `json:"required,omitempty"`
	RequirementTypeID string  `json:"requirement_type_id,omitempty"`
}

// Creates or overwrites a requirement with the specified parameters.
func (c *Client) CreateRequirement(ctx context.Context, r *CreateRequirementRequest) (*Requirement, error) {
	var resp Requirement
	if err := c.request(ctx, "POST", "/v0/requirements", r, &resp); err != nil {
		return nil, fmt.Errorf("CreateRequirement: %w", err)
	}
	return &resp, nil
}

type ListRequirementsRequest struct {
	// Filter results to a specific set of requirement types.
	RequirementTypes []string `url:"requirement_types,omitempty"`

	// Together with start_time, this determines the interval for which
	// requirements will be retrieved.
	EndTime Timestamp `url:"end_time,omitempty"`

	StartTime Timestamp `url:"start_time,omitempty"`
}

type ListRequirementsResponse struct {
	Requirements []Requirement `json:"requirements,omitempty"`
}

// Returns a list of requirement objects that match the provided query.
func (c *Client) ListRequirements(ctx context.Context, r *ListRequirementsRequest) (*ListRequirementsResponse, error) {
	var resp ListRequirementsResponse
	if err := c.request(ctx, "GET", "/v0/requirements", r, &resp); err != nil {
		return nil, fmt.Errorf("ListRequirements: %w", err)
	}
	return &resp, nil
}

// Requirement types contain metadata about named set of staffing
// requirements.
type RequirementType struct {
	ID string `json:"id,omitempty"`

	// List of unique identifiers for activity types that count as scheduled
	// against the requirement.
	ActivityTypeIDs []string `json:"activity_type_ids,omitempty"`

	Name string `json:"name,omitempty"`
}

type ListRequirementTypesResponse struct {
	RequirementTypes map[string]RequirementType `json:"requirement_types,omitempty"`
}

// Returns a list of all requirement type objects configured on the account.
func (c *Client) ListRequirementTypes(ctx context.Context) (*ListRequirementTypesResponse, error) {
	var resp ListRequirementTypesResponse
	if err := c.request(ctx, "GET", "/v0/requirement_types", nil, &resp); err != nil {
		return nil, fmt.Errorf("ListRequirementTypes: %w", err)
	}
	return &resp, nil
}
