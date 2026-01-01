package services

import (
	"context"

	"github.com/Sellium-site/sellium-go/core"
)

type StoreService struct{ c *core.Client }

func NewStore(c *core.Client) *StoreService { return &StoreService{c: c} }

type GetStoreResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Store core.Store      `json:"store"`
		Stats core.StoreStats `json:"stats"`
	} `json:"data"`
}

func (s *StoreService) Get(ctx context.Context) (*GetStoreResponse, *core.ResponseMeta, error) {
	var out GetStoreResponse
	meta, err := s.c.Do(ctx, "GET", "/store", nil, nil, &out)
	return &out, meta, err
}
