package assembled

import (
	"context"
	"fmt"
)

// Agents can have scheduled activities that are either created directly in
// Assembled or sourced from other calendaring applications such as Google
// Calendar.
type Activity struct {
	ID string `json:"id,omitempty"`

	// An arbitrary string to be displayed to end users.
	Description string `json:"description,omitempty"`

	AgentID   string    `json:"agent_id,omitempty"` // Identifier for the corresponding agent.
	EndTime   Timestamp `json:"end_time,omitempty"`
	StartTime Timestamp `json:"start_time,omitempty"`
	TypeID    string    `json:"type_id,omitempty"` // Identifier for the associated activity type.
}

type CreateActivityRequest struct {
	// Whether or not to allow conflicting events. If true, this created
	// activity will be allowed to have conflicts with other activities. If
	// false, any overlapping activities for the specified agent will be
	// deleted. Defaults to false.
	AllowConflicts bool `json:"allow_conflicts,omitempty"`

	// An arbitrary string to be displayed to end users.
	Description string `json:"description,omitempty"`

	// Identifier for the corresponding schedule. Defaults to the master
	// schedule.
	ScheduleID string `json:"schedule_id,omitempty"`

	AgentID   string    `json:"agent_id,omitempty"` // Unique identifier for an agent.
	EndTime   Timestamp `json:"end_time,omitempty"`
	StartTime Timestamp `json:"start_time,omitempty"`
	TypeID    string    `json:"type_id,omitempty"` // Unique identifier for an activity type.
}

// Creates an activity with the specified parameters. Valid IDs for agents can
// be retrieved from the agent endpoints. Valid IDs for activity types can be
// retrieved from the activity types endpoints. This endpoint will return 400
// if invalid IDs are provided.
func (c *Client) CreateActivity(ctx context.Context, r *CreateActivityRequest) (*Activity, error) {
	var resp Activity
	if err := c.request(ctx, "POST", "/v0/activities", r, &resp); err != nil {
		return nil, fmt.Errorf("CreateActivity: %w", err)
	}
	return &resp, nil
}

type ActivityRequest struct {
	// See documentation for the Activity object. Note that the
	// allow_conflicts parameter is not supported in the bulk endpoint.
	Activity Activity `json:"activity,omitempty"`

	// One of create, update, or delete. Note that update and delete are
	// currently in beta. Contact us at support@assembled.com to be notified
	// of updates.
	Action string `json:"action,omitempty"`
}

type CreateBulkActivityRequest struct {
	// An array of requests following the format below.
	Acitivites []ActivityRequest `json:"acitivites,omitempty"`

	// Identifier for the corresponding schedule. Defaults to the master
	// schedule.
	ScheduleID string `json:"schedule_id,omitempty"`
}

type CreateBulkActivityResponse struct {
	Activities map[string]Activity `json:"activities,omitempty"`
}

// Creates, updates, or deletes multiple activities in a single request. A
// given request occurs within a transaction, so in the event of an error
// during the request, it should be assumed that no changes were processed.
func (c *Client) CreateBulkActivity(ctx context.Context, r *CreateBulkActivityRequest) (*CreateBulkActivityResponse, error) {
	var resp CreateBulkActivityResponse
	if err := c.request(ctx, "POST", "/v0/activities/bulk", r, &resp); err != nil {
		return nil, fmt.Errorf("CreateBulkActivity: %w", err)
	}
	return &resp, nil
}

type ListActivitiesRequest struct {
	// If true, include associated agent objects. Defaults to false.
	IncludeAgents bool `url:"include_agents,omitempty"`

	// If true, include activity types active on the account. Defaults to
	// false.
	IncludeActivityTypes bool `url:"include_activity_types,omitempty"`

	// Identifier for the corresponding schedule. Defaults to the master
	// schedule.
	ScheduleID string `url:"schedule_id,omitempty"`

	// Together with start_time, this determines the interval for which shifts
	// will be retrieved.
	EndTime Timestamp `url:"end_time,omitempty"`

	Agents    []string  `url:"agents,omitempty"`  // Filter results to a specific set of agents.
	Channel   string    `url:"channel,omitempty"` // Filter results to a specific channel.
	StartTime Timestamp `url:"start_time,omitempty"`
	Team      string    `url:"team,omitempty"`  // Filter results to a specific team.
	Types     []string  `url:"types,omitempty"` // Filter results to a specific set of types.
}

type Queue struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ListActivitiesResponse struct {
	Activities map[string]Activity `json:"activities,omitempty"`
	Agents     map[string]Agent    `json:"agents,omitempty"`
	Queues     map[string]Queue    `json:"queues,omitempty"`
}

// Returns a list of activity objects that match the provided query.
func (c *Client) ListActivities(ctx context.Context, r *ListActivitiesRequest) (*ListActivitiesResponse, error) {
	var resp ListActivitiesResponse
	if err := c.request(ctx, "GET", "/v0/activities", r, &resp); err != nil {
		return nil, fmt.Errorf("ListActivities: %w", err)
	}
	return &resp, nil
}

type DeleteActivitiesRequest struct {
	// Unique identifiers for agents. All activities in the deletion window
	// will be deleted for each specified agent.
	AgentIDs []string `json:"agent_ids,omitempty"`

	// Identifier for the corresponding schedule. Defaults to the master
	// schedule.
	ScheduleID string `json:"schedule_id,omitempty"`

	EndTime   Timestamp `json:"end_time,omitempty"`   // The end of the deletion window.
	StartTime Timestamp `json:"start_time,omitempty"` // The start of the deletion window.
}

// Deletes all activities that match the specified parameters. Valid IDs for
// agents can be retrieved from the agent endpoints. This endpoint will return
// 400 if invalid IDs are provided.
//
// Activities can be partially deleted. For example, if agent XYZ has an
// activity from 3pm-5pm and the deletion window is from 4pm-5pm, there will
// still exist a 3-4pm activity for agent XYZ after the deletion is completed.
func (c *Client) DeleteActivities(ctx context.Context, r *DeleteActivitiesRequest) error {
	if err := c.request(ctx, "DELETE", "/v0/activities", r, nil); err != nil {
		return fmt.Errorf("DeleteActivities: %w", err)
	}
	return nil
}

// Activity types represent a categorization of different activities.
// Currently they are only editable via the dashboard.
type ActivityType struct {
	ID string `json:"id,omitempty"`

	// Channels associated with the activity. Must be non-empty when the the
	// activity is productive.
	Channels []string `json:"channels,omitempty"`

	// Corresponds to type in the Activity object, will be deprecated in a
	// future API version.
	Value string `json:"value,omitempty"`

	BackgroundColor string `json:"background_color,omitempty"` // Hex string.
	FontColor       string `json:"font_color,omitempty"`       // Hex string.
	Name            string `json:"name,omitempty"`
	Productive      bool   `json:"productive,omitempty"` // If true, timeoff must be false.
	ShortName       string `json:"short_name,omitempty"`
	Timeoff         bool   `json:"timeoff,omitempty"` // If true, productive must be false.
}

type ListActivityTypesResponse struct {
	ActivityTypes map[string]ActivityType `json:"activity_types,omitempty"`
}

// Returns a list of all activity type objects configured on the account.
func (c *Client) ListActivityTypes(ctx context.Context) (*ListActivityTypesResponse, error) {
	var resp ListActivityTypesResponse
	if err := c.request(ctx, "GET", "/v0/activity_types", nil, &resp); err != nil {
		return nil, fmt.Errorf("ListActivityTypes: %w", err)
	}
	return &resp, nil
}

type CreateActivityTypeRequest struct {
	// Channels associated with the activity. Must be non-empty when the the
	// activity is productive.
	Channels []string `json:"channels,omitempty"`

	// Corresponds to type in the Activity object, will be deprecated in a
	// future API version.
	Value string `json:"value,omitempty"`

	BackgroundColor string `json:"background_color,omitempty"` // Hex string.
	FontColor       string `json:"font_color,omitempty"`       // Hex string.
	Name            string `json:"name,omitempty"`
	Productive      bool   `json:"productive,omitempty"` // If true, timeoff must be false.
	ShortName       string `json:"short_name,omitempty"`
	Timeoff         bool   `json:"timeoff,omitempty"` // If true, productive must be false.
}

// Creates an activity type.
func (c *Client) CreateActivityType(ctx context.Context, r *CreateActivityTypeRequest) (*ActivityType, error) {
	var resp ActivityType
	if err := c.request(ctx, "POST", "/v0/activity_types", r, &resp); err != nil {
		return nil, fmt.Errorf("CreateActivityType: %w", err)
	}
	return &resp, nil
}

type DeleteActivityTypeRequest struct {
	ID string `json:"id,omitempty"`
}

// Deletes an activity type.
func (c *Client) DeleteActivityType(ctx context.Context, r *DeleteActivityTypeRequest) (*ActivityType, error) {
	var resp ActivityType
	if err := c.request(ctx, "DELETE", fmt.Sprintf("/v0/activity_types/%s", r.ID), r, &resp); err != nil {
		return nil, fmt.Errorf("DeleteActivityType: %w", err)
	}
	return &resp, nil
}
