// Code generated by riza; DO NOT EDIT.

package assembled

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type CreateRequirementRequest struct {
	EndTime           time.Time `json:"end_time,omitempty"`
	Required          int       `json:"required,omitempty"`
	RequirementTypeID string    `json:"requirement_type_id,omitempty"`
	StartTime         time.Time `json:"start_time,omitempty"`
}

func (r CreateRequirementRequest) MarshalJSON() ([]byte, error) {
	type Alias CreateRequirementRequest
	return json.Marshal(&struct {
		Alias
		EndTime   int64 `json:"end_time,omitempty"`
		StartTime int64 `json:"start_time,omitempty"`
	}{
		Alias:     (Alias)(r),
		EndTime:   timestamp(r.EndTime),
		StartTime: timestamp(r.StartTime),
	})
}

func (r *CreateRequirementRequest) UnmarshalJSON(b []byte) error {
	type Alias CreateRequirementRequest
	var a struct {
		Alias
		EndTime   int64 `json:"end_time,omitempty"`
		StartTime int64 `json:"start_time,omitempty"`
	}
	err := json.Unmarshal(b, &a)
	if err != nil {
		return err
	}
	*r = CreateRequirementRequest(a.Alias)
	r.EndTime = time.Unix(a.EndTime, 0)
	r.StartTime = time.Unix(a.StartTime, 0)
	return nil
}

type ListRequirementsRequest struct {
	// Filter results to a specific set of requirement types.
	RequirementTypes []string `json:"requirement_types,omitempty"`

	EndTime   time.Time `json:"end_time,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
}

func (r ListRequirementsRequest) MarshalJSON() ([]byte, error) {
	type Alias ListRequirementsRequest
	return json.Marshal(&struct {
		Alias
		EndTime   int64 `json:"end_time,omitempty"`
		StartTime int64 `json:"start_time,omitempty"`
	}{
		Alias:     (Alias)(r),
		EndTime:   timestamp(r.EndTime),
		StartTime: timestamp(r.StartTime),
	})
}

func (r *ListRequirementsRequest) UnmarshalJSON(b []byte) error {
	type Alias ListRequirementsRequest
	var a struct {
		Alias
		EndTime   int64 `json:"end_time,omitempty"`
		StartTime int64 `json:"start_time,omitempty"`
	}
	err := json.Unmarshal(b, &a)
	if err != nil {
		return err
	}
	*r = ListRequirementsRequest(a.Alias)
	r.EndTime = time.Unix(a.EndTime, 0)
	r.StartTime = time.Unix(a.StartTime, 0)
	return nil
}

// Returns ListRequirementsRequest with `url` tags instead of `json` tags for URL encoding.
func (r *ListRequirementsRequest) params() interface{} {
	if r == nil {
		return r
	}
	type params struct {
		RequirementTypes []string  `url:"requirement_types,omitempty"`
		EndTime          time.Time `url:"end_time,omitempty,unix"`
		StartTime        time.Time `url:"start_time,omitempty,unix"`
	}
	p := params(*r)
	return &p
}

type ListRequirementsResponse struct {
	Requirements []Requirement `json:"requirements,omitempty"`
}

// Requirements represent a set of staffing requirements and grouped staffing
// against the requirement within a given time window. They correspond to
// metrics about required staffing and scheduled staffing that appear above
// the team calendar.
type Requirement struct {
	// Count of required staffing in the interval, can be partial.
	Required int `json:"required,omitempty"`

	// Unique identifier for the corresponding requirement type.
	RequirementTypeID string `json:"requirement_type_id,omitempty"`

	// Count of scheduled staffing in the interval, can be partial.
	Scheduled int `json:"scheduled,omitempty"`

	EndTime   time.Time `json:"end_time,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
}

func (r Requirement) MarshalJSON() ([]byte, error) {
	type Alias Requirement
	return json.Marshal(&struct {
		Alias
		EndTime   int64 `json:"end_time,omitempty"`
		StartTime int64 `json:"start_time,omitempty"`
	}{
		Alias:     (Alias)(r),
		EndTime:   timestamp(r.EndTime),
		StartTime: timestamp(r.StartTime),
	})
}

func (r *Requirement) UnmarshalJSON(b []byte) error {
	type Alias Requirement
	var a struct {
		Alias
		EndTime   int64 `json:"end_time,omitempty"`
		StartTime int64 `json:"start_time,omitempty"`
	}
	err := json.Unmarshal(b, &a)
	if err != nil {
		return err
	}
	*r = Requirement(a.Alias)
	r.EndTime = time.Unix(a.EndTime, 0)
	r.StartTime = time.Unix(a.StartTime, 0)
	return nil
}

// Creates or overwrites a requirement with the specified parameters.
func (c *Client) CreateRequirement(ctx context.Context, r *CreateRequirementRequest) (*Requirement, error) {
	var resp Requirement
	if err := c.request(ctx, "POST", "/v0/requirements", nil, r, &resp); err != nil {
		return nil, fmt.Errorf("CreateRequirement: %w", err)
	}
	return &resp, nil
}

// Returns a list of requirement objects that match the provided query.
func (c *Client) ListRequirements(ctx context.Context, r *ListRequirementsRequest) (*ListRequirementsResponse, error) {
	var resp ListRequirementsResponse
	if err := c.request(ctx, "GET", "/v0/requirements", r.params(), nil, &resp); err != nil {
		return nil, fmt.Errorf("ListRequirements: %w", err)
	}
	return &resp, nil
}