package services

import (
	"context"
	"net/url"
	"strconv"

	"github.com/Sellium-site/sellium-go/core"
)

type CustomersService struct{ c *core.Client }

func NewCustomers(c *core.Client) *CustomersService { return &CustomersService{c: c} }

type ListCustomersParams struct {
	Page  int
	Limit int
	Email string
}

type ListCustomersResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Customers  []core.CustomerRow `json:"customers"`
		Pagination core.Pagination    `json:"pagination"`
	} `json:"data"`
}

func (s *CustomersService) List(ctx context.Context, p *ListCustomersParams) (*ListCustomersResponse, *core.ResponseMeta, error) {
	q := url.Values{}
	if p != nil {
		if p.Page > 0 {
			q.Set("page", strconv.Itoa(p.Page))
		}
		if p.Limit > 0 {
			q.Set("limit", strconv.Itoa(p.Limit))
		}
		if p.Email != "" {
			q.Set("email", p.Email)
		}
	}

	var out ListCustomersResponse
	meta, err := s.c.Do(ctx, "GET", "/customers", q, nil, &out)
	return &out, meta, err
}

type GetCustomerResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Customer     core.CustomerDetail        `json:"customer"`
		RecentOrders []core.CustomerRecentOrder `json:"recent_orders"`
	} `json:"data"`
}

func (s *CustomersService) Get(ctx context.Context, email string) (*GetCustomerResponse, *core.ResponseMeta, error) {
	encoded := url.PathEscape(email)
	var out GetCustomerResponse
	meta, err := s.c.Do(ctx, "GET", "/customers/"+encoded, nil, nil, &out)
	return &out, meta, err
}
