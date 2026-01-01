package services

import (
	"context"
	"net/url"
	"strconv"

	"github.com/Sellium-site/sellium-go/core"
)

type GroupsService struct{ c *core.Client }

func NewGroups(c *core.Client) *GroupsService { return &GroupsService{c: c} }

type ListGroupsParams struct {
	Page   int
	Limit  int
	Active *bool
	Search string
}

type ListGroupsResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Groups     []core.Group    `json:"groups"`
		Pagination core.Pagination `json:"pagination"`
	} `json:"data"`
}

func (s *GroupsService) List(ctx context.Context, p *ListGroupsParams) (*ListGroupsResponse, *core.ResponseMeta, error) {
	q := url.Values{}
	if p != nil {
		if p.Page > 0 {
			q.Set("page", strconv.Itoa(p.Page))
		}
		if p.Limit > 0 {
			q.Set("limit", strconv.Itoa(p.Limit))
		}
		if p.Active != nil {
			q.Set("active", strconv.FormatBool(*p.Active))
		}
		if p.Search != "" {
			q.Set("search", p.Search)
		}
	}
	var out ListGroupsResponse
	meta, err := s.c.Do(ctx, "GET", "/groups", q, nil, &out)
	return &out, meta, err
}

type CreateGroupRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description,omitempty"`
	ImageURL     string `json:"image_url,omitempty"`
	DisplayOrder *int   `json:"display_order,omitempty"`
	IsActive     *bool  `json:"is_active,omitempty"`
}

type GroupResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Group core.Group `json:"group"`
	} `json:"data"`
}

func (s *GroupsService) Create(ctx context.Context, req CreateGroupRequest) (*GroupResponse, *core.ResponseMeta, error) {
	var out GroupResponse
	meta, err := s.c.Do(ctx, "POST", "/groups", nil, req, &out)
	return &out, meta, err
}

type GetGroupResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Group struct {
			core.Group
			Products []core.GroupProductMini `json:"products,omitempty"`
		} `json:"group"`
	} `json:"data"`
}

func (s *GroupsService) Get(ctx context.Context, groupID string) (*GetGroupResponse, *core.ResponseMeta, error) {
	var out GetGroupResponse
	meta, err := s.c.Do(ctx, "GET", "/groups/"+groupID, nil, nil, &out)
	return &out, meta, err
}

type UpdateGroupRequest struct {
	Name         *string `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
	ImageURL     *string `json:"image_url,omitempty"`
	DisplayOrder *int    `json:"display_order,omitempty"`
	IsActive     *bool   `json:"is_active,omitempty"`
}

func (s *GroupsService) Update(ctx context.Context, groupID string, req UpdateGroupRequest) (*GroupResponse, *core.ResponseMeta, error) {
	var out GroupResponse
	meta, err := s.c.Do(ctx, "PATCH", "/groups/"+groupID, nil, req, &out)
	return &out, meta, err
}

type DeleteGroupResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Deleted bool   `json:"deleted"`
		GroupID string `json:"group_id,omitempty"`
	} `json:"data"`
}

func (s *GroupsService) Delete(ctx context.Context, groupID string) (*DeleteGroupResponse, *core.ResponseMeta, error) {
	var out DeleteGroupResponse
	meta, err := s.c.Do(ctx, "DELETE", "/groups/"+groupID, nil, nil, &out)
	return &out, meta, err
}
