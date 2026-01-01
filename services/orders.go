package services

import (
	"context"
	"net/url"
	"strconv"

	"github.com/Sellium-site/sellium-go/core"
)

type OrdersService struct{ c *core.Client }

func NewOrders(c *core.Client) *OrdersService { return &OrdersService{c: c} }

type ListOrdersParams struct {
	Page          int
	Limit         int
	Status        string
	ProductID     string
	CustomerEmail string
}

type ListOrdersResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Orders     []core.Order    `json:"orders"`
		Pagination core.Pagination `json:"pagination"`
	} `json:"data"`
}

func (s *OrdersService) List(ctx context.Context, p *ListOrdersParams) (*ListOrdersResponse, *core.ResponseMeta, error) {
	q := url.Values{}
	if p != nil {
		if p.Page > 0 {
			q.Set("page", strconv.Itoa(p.Page))
		}
		if p.Limit > 0 {
			q.Set("limit", strconv.Itoa(p.Limit))
		}
		if p.Status != "" {
			q.Set("status", p.Status)
		}
		if p.ProductID != "" {
			q.Set("product_id", p.ProductID)
		}
		if p.CustomerEmail != "" {
			q.Set("customer_email", p.CustomerEmail)
		}
	}

	var out ListOrdersResponse
	meta, err := s.c.Do(ctx, "GET", "/orders", q, nil, &out)
	return &out, meta, err
}

type CreateOrderRequest struct {
	ProductID     string `json:"product_id"`
	CustomerEmail string `json:"customer_email"`
	Quantity      int    `json:"quantity"`

	CustomerName  string `json:"customer_name,omitempty"`
	PaymentMethod string `json:"payment_method,omitempty"` // stripe|paypal|crypto|cashapp
	CustomFields  any    `json:"custom_fields,omitempty"`  // docs: object
	AffiliateCode string `json:"affiliate_code,omitempty"`
}

type OrderResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Order core.Order `json:"order"`
	} `json:"data"`
}

func (s *OrdersService) Create(ctx context.Context, req CreateOrderRequest) (*OrderResponse, *core.ResponseMeta, error) {
	var out OrderResponse
	meta, err := s.c.Do(ctx, "POST", "/orders", nil, req, &out)
	return &out, meta, err
}

func (s *OrdersService) Get(ctx context.Context, orderID string) (*OrderResponse, *core.ResponseMeta, error) {
	var out OrderResponse
	meta, err := s.c.Do(ctx, "GET", "/orders/"+orderID, nil, nil, &out)
	return &out, meta, err
}

type UpdateOrderRequest struct {
	Status        *string `json:"status,omitempty"` // pending|completed|canceled|refunded
	CustomerName  *string `json:"customer_name,omitempty"`
	TransactionID *string `json:"transaction_id,omitempty"`
	CustomFields  any     `json:"custom_fields,omitempty"`
}

type UpdateOrderResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Order         core.Order `json:"order"`
		Delivery      any        `json:"delivery,omitempty"`
		Warning       string     `json:"warning,omitempty"`
		DeliveryError string     `json:"delivery_error,omitempty"`
	} `json:"data"`
}

func (s *OrdersService) Update(ctx context.Context, orderID string, req UpdateOrderRequest) (*UpdateOrderResponse, *core.ResponseMeta, error) {
	var out UpdateOrderResponse
	meta, err := s.c.Do(ctx, "PATCH", "/orders/"+orderID, nil, req, &out)
	return &out, meta, err
}

type DeleteOrderResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Deleted bool   `json:"deleted"`
		ID      string `json:"id,omitempty"`
	} `json:"data"`
}

func (s *OrdersService) Delete(ctx context.Context, orderID string) (*DeleteOrderResponse, *core.ResponseMeta, error) {
	var out DeleteOrderResponse
	meta, err := s.c.Do(ctx, "DELETE", "/orders/"+orderID, nil, nil, &out)
	return &out, meta, err
}
