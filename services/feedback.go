package services

import (
	"context"
	"net/url"
	"strconv"

	"github.com/Sellium-site/sellium-go/core"
)

type FeedbackService struct{ c *core.Client }

func NewFeedback(c *core.Client) *FeedbackService { return &FeedbackService{c: c} }

type ListFeedbackParams struct {
	Page        int
	Limit       int
	Rating      *int
	HasResponse *bool
	IsVisible   *bool
	Email       string
}

type ListFeedbackResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Feedback   []core.Feedback `json:"feedback"`
		Pagination core.Pagination `json:"pagination"`
	} `json:"data"`
}

func (s *FeedbackService) List(ctx context.Context, p *ListFeedbackParams) (*ListFeedbackResponse, *core.ResponseMeta, error) {
	q := url.Values{}
	if p != nil {
		if p.Page > 0 {
			q.Set("page", strconv.Itoa(p.Page))
		}
		if p.Limit > 0 {
			q.Set("limit", strconv.Itoa(p.Limit))
		}
		if p.Rating != nil {
			q.Set("rating", strconv.Itoa(*p.Rating))
		}
		if p.HasResponse != nil {
			q.Set("has_response", strconv.FormatBool(*p.HasResponse))
		}
		if p.IsVisible != nil {
			q.Set("is_visible", strconv.FormatBool(*p.IsVisible))
		}
		if p.Email != "" {
			q.Set("email", p.Email)
		}
	}

	var out ListFeedbackResponse
	meta, err := s.c.Do(ctx, "GET", "/feedback", q, nil, &out)
	return &out, meta, err
}

type GetFeedbackResponse struct {
	Success bool          `json:"success"`
	Data    core.Feedback `json:"data"`
}

func (s *FeedbackService) Get(ctx context.Context, feedbackID string) (*GetFeedbackResponse, *core.ResponseMeta, error) {
	var out GetFeedbackResponse
	meta, err := s.c.Do(ctx, "GET", "/feedback/"+feedbackID, nil, nil, &out)
	return &out, meta, err
}

type UpdateFeedbackRequest struct {
	Response  *string `json:"response,omitempty"` // docs allow null/empty to remove
	IsVisible *bool   `json:"is_visible,omitempty"`
}

type UpdateFeedbackResponse struct {
	Success bool          `json:"success"`
	Data    core.Feedback `json:"data"`
}

func (s *FeedbackService) Update(ctx context.Context, feedbackID string, req UpdateFeedbackRequest) (*UpdateFeedbackResponse, *core.ResponseMeta, error) {
	var out UpdateFeedbackResponse
	meta, err := s.c.Do(ctx, "PATCH", "/feedback/"+feedbackID, nil, req, &out)
	return &out, meta, err
}
