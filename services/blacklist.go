package services

import (
	"context"
	"net/url"
	"strconv"

	"github.com/Sellium-site/sellium-go/core"
)

type BlacklistService struct{ c *core.Client }

func NewBlacklist(c *core.Client) *BlacklistService { return &BlacklistService{c: c} }

type ListBlacklistParams struct {
	Page   int
	Limit  int
	Type   string
	Search string
}

type ListBlacklistResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Entries    []core.BlacklistEntry `json:"entries"`
		Pagination core.Pagination       `json:"pagination"`
	} `json:"data"`
}

func (s *BlacklistService) List(ctx context.Context, p *ListBlacklistParams) (*ListBlacklistResponse, *core.ResponseMeta, error) {
	q := url.Values{}
	if p != nil {
		if p.Page > 0 {
			q.Set("page", strconv.Itoa(p.Page))
		}
		if p.Limit > 0 {
			q.Set("limit", strconv.Itoa(p.Limit))
		}
		if p.Type != "" {
			q.Set("type", p.Type)
		}
		if p.Search != "" {
			q.Set("search", p.Search)
		}
	}

	var out ListBlacklistResponse
	meta, err := s.c.Do(ctx, "GET", "/blacklist", q, nil, &out)
	return &out, meta, err
}

type GetBlacklistEntryResponse struct {
	Success bool                `json:"success"`
	Data    core.BlacklistEntry `json:"data"`
}

func (s *BlacklistService) Get(ctx context.Context, entryID string) (*GetBlacklistEntryResponse, *core.ResponseMeta, error) {
	var out GetBlacklistEntryResponse
	meta, err := s.c.Do(ctx, "GET", "/blacklist/"+entryID, nil, nil, &out)
	return &out, meta, err
}

type CreateBlacklistEntryRequest struct {
	Type   string `json:"type"` // email|ip|country
	Value  string `json:"value"`
	Reason string `json:"reason,omitempty"`
}

type CreateBlacklistEntryResponse struct {
	Success bool                `json:"success"`
	Data    core.BlacklistEntry `json:"data"`
}

func (s *BlacklistService) Create(ctx context.Context, req CreateBlacklistEntryRequest) (*CreateBlacklistEntryResponse, *core.ResponseMeta, error) {
	var out CreateBlacklistEntryResponse
	meta, err := s.c.Do(ctx, "POST", "/blacklist", nil, req, &out)
	return &out, meta, err
}

type DeleteBlacklistEntryResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Deleted bool `json:"deleted"`
	} `json:"data"`
}

func (s *BlacklistService) Delete(ctx context.Context, entryID string) (*DeleteBlacklistEntryResponse, *core.ResponseMeta, error) {
	var out DeleteBlacklistEntryResponse
	meta, err := s.c.Do(ctx, "DELETE", "/blacklist/"+entryID, nil, nil, &out)
	return &out, meta, err
}
