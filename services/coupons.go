package services

import (
	"context"
	"net/url"
	"strconv"

	"github.com/Sellium-site/sellium-go/core"
)

type CouponsService struct{ c *core.Client }

func NewCoupons(c *core.Client) *CouponsService { return &CouponsService{c: c} }

type ListCouponsParams struct {
	Page   int
	Limit  int
	Active *bool
	Code   string
}

type ListCouponsResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Coupons    []core.Coupon   `json:"coupons"`
		Pagination core.Pagination `json:"pagination"`
	} `json:"data"`
}

func (s *CouponsService) List(ctx context.Context, p *ListCouponsParams) (*ListCouponsResponse, *core.ResponseMeta, error) {
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
		if p.Code != "" {
			q.Set("code", p.Code)
		}
	}
	var out ListCouponsResponse
	meta, err := s.c.Do(ctx, "GET", "/coupons", q, nil, &out)
	return &out, meta, err
}

type CreateCouponRequest struct {
	Code  string `json:"code"`
	Type  string `json:"type"`  // percentage|fixed
	Value int    `json:"value"` // percentage 1-100, fixed cents

	MinimumPurchase *int    `json:"minimum_purchase,omitempty"`
	MaximumUses     *int    `json:"maximum_uses,omitempty"`
	IsActive        *bool   `json:"is_active,omitempty"`
	ExpiresAt       *string `json:"expires_at,omitempty"`
}

type CouponResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

func (s *CouponsService) Create(ctx context.Context, req CreateCouponRequest) (*CouponResponse, *core.ResponseMeta, error) {
	var out CouponResponse
	meta, err := s.c.Do(ctx, "POST", "/coupons", nil, req, &out)
	return &out, meta, err
}

type GetCouponResponse struct {
	Success bool        `json:"success"`
	Data    core.Coupon `json:"data"`
}

func (s *CouponsService) Get(ctx context.Context, couponID string) (*GetCouponResponse, *core.ResponseMeta, error) {
	var out GetCouponResponse
	meta, err := s.c.Do(ctx, "GET", "/coupons/"+couponID, nil, nil, &out)
	return &out, meta, err
}

type UpdateCouponRequest struct {
	Code            *string `json:"code,omitempty"`
	Type            *string `json:"type,omitempty"`
	Value           *int    `json:"value,omitempty"`
	MinimumPurchase *int    `json:"minimum_purchase,omitempty"`
	MaximumUses     *int    `json:"maximum_uses,omitempty"`
	IsActive        *bool   `json:"is_active,omitempty"`
	ExpiresAt       *string `json:"expires_at,omitempty"`
}

func (s *CouponsService) Update(ctx context.Context, couponID string, req UpdateCouponRequest) (*CouponResponse, *core.ResponseMeta, error) {
	var out CouponResponse
	meta, err := s.c.Do(ctx, "PATCH", "/coupons/"+couponID, nil, req, &out)
	return &out, meta, err
}

type DeleteCouponResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Deleted bool   `json:"deleted"`
		ID      string `json:"id,omitempty"`
	} `json:"data"`
}

func (s *CouponsService) Delete(ctx context.Context, couponID string) (*DeleteCouponResponse, *core.ResponseMeta, error) {
	var out DeleteCouponResponse
	meta, err := s.c.Do(ctx, "DELETE", "/coupons/"+couponID, nil, nil, &out)
	return &out, meta, err
}
