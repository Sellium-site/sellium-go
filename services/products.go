package services

import (
	"context"
	"net/url"
	"strconv"

	"github.com/Sellium-site/sellium-go/core"
)

type ProductsService struct{ c *core.Client }

func NewProducts(c *core.Client) *ProductsService { return &ProductsService{c: c} }

type ListProductsParams struct {
	Page    int
	Limit   int
	Active  *bool
	GroupID string
}

type ListProductsResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Products   []core.Product  `json:"products"`
		Pagination core.Pagination `json:"pagination"`
	} `json:"data"`
}

func (s *ProductsService) List(ctx context.Context, p *ListProductsParams) (*ListProductsResponse, *core.ResponseMeta, error) {
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
		if p.GroupID != "" {
			q.Set("group_id", p.GroupID)
		}
	}
	var out ListProductsResponse
	meta, err := s.c.Do(ctx, "GET", "/products", q, nil, &out)
	return &out, meta, err
}

type CreateProductRequest struct {
	Name         string `json:"name"`
	PriceInCents int    `json:"price_in_cents"`
	DeliveryType string `json:"delivery_type"` // file|serials|service|dynamic

	Description     string `json:"description,omitempty"`
	ImageURL        string `json:"image_url,omitempty"`
	IsActive        *bool  `json:"is_active,omitempty"`
	StockQuantity   *int   `json:"stock_quantity,omitempty"`
	MinimumQuantity *int   `json:"minimum_quantity,omitempty"`
	MaximumQuantity *int   `json:"maximum_quantity,omitempty"`
	Unlisted        *bool  `json:"unlisted,omitempty"`
	IsPrivate       *bool  `json:"is_private,omitempty"`
	OnHold          *bool  `json:"on_hold,omitempty"`
	Warranty        string `json:"warranty,omitempty"`
	ProductTerms    string `json:"product_terms,omitempty"`
	GroupID         string `json:"group_id,omitempty"`

	Serials           []string `json:"serials,omitempty"`
	FileURL           string   `json:"file_url,omitempty"`
	ServiceMessage    string   `json:"service_message,omitempty"`
	DeliveryText      string   `json:"delivery_text,omitempty"`
	DynamicWebhookURL string   `json:"dynamic_webhook_url,omitempty"`
	RedirectURL       string   `json:"redirect_url,omitempty"`
	YoutubeURL        string   `json:"youtube_url,omitempty"`

	DiscordEnabled  *bool `json:"discord_enabled,omitempty"`
	DiscordOptional *bool `json:"discord_optional,omitempty"`

	EnableLicenseSystem *bool `json:"enable_license_system,omitempty"`
	LicenseMaxDevices   *int  `json:"license_max_devices,omitempty"`
	LicenseExpiresDays  *int  `json:"license_expires_days,omitempty"`

	CustomFields    any `json:"custom_fields,omitempty"`
	VolumeDiscounts any `json:"volume_discounts,omitempty"`
	PaymentMethods  any `json:"payment_methods,omitempty"`
	WebhookURLs     any `json:"webhook_urls,omitempty"`
}

type GetProductResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Product core.Product `json:"product"`
	} `json:"data"`
}

func (s *ProductsService) Create(ctx context.Context, req CreateProductRequest) (*GetProductResponse, *core.ResponseMeta, error) {
	var out GetProductResponse
	meta, err := s.c.Do(ctx, "POST", "/products", nil, req, &out)
	return &out, meta, err
}

func (s *ProductsService) Get(ctx context.Context, productID string) (*GetProductResponse, *core.ResponseMeta, error) {
	var out GetProductResponse
	meta, err := s.c.Do(ctx, "GET", "/products/"+productID, nil, nil, &out)
	return &out, meta, err
}

type UpdateProductRequest struct {
	Name            *string `json:"name,omitempty"`
	Description     *string `json:"description,omitempty"`
	ImageURL        *string `json:"image_url,omitempty"`
	PriceInCents    *int    `json:"price_in_cents,omitempty"`
	DeliveryType    *string `json:"delivery_type,omitempty"`
	IsActive        *bool   `json:"is_active,omitempty"`
	StockQuantity   *int    `json:"stock_quantity,omitempty"`
	MinimumQuantity *int    `json:"minimum_quantity,omitempty"`
	MaximumQuantity *int    `json:"maximum_quantity,omitempty"`
	Unlisted        *bool   `json:"unlisted,omitempty"`
	IsPrivate       *bool   `json:"is_private,omitempty"`
	OnHold          *bool   `json:"on_hold,omitempty"`
	Warranty        *string `json:"warranty,omitempty"`
	ProductTerms    *string `json:"product_terms,omitempty"`
	GroupID         *string `json:"group_id,omitempty"`

	Serials           *[]string `json:"serials,omitempty"`
	FileURL           *string   `json:"file_url,omitempty"`
	ServiceMessage    *string   `json:"service_message,omitempty"`
	DeliveryText      *string   `json:"delivery_text,omitempty"`
	DynamicWebhookURL *string   `json:"dynamic_webhook_url,omitempty"`
	RedirectURL       *string   `json:"redirect_url,omitempty"`
	YoutubeURL        *string   `json:"youtube_url,omitempty"`

	DiscordEnabled      *bool `json:"discord_enabled,omitempty"`
	DiscordOptional     *bool `json:"discord_optional,omitempty"`
	EnableLicenseSystem *bool `json:"enable_license_system,omitempty"`
	LicenseMaxDevices   *int  `json:"license_max_devices,omitempty"`
	LicenseExpiresDays  *int  `json:"license_expires_days,omitempty"`

	CustomFields    any `json:"custom_fields,omitempty"`
	VolumeDiscounts any `json:"volume_discounts,omitempty"`
	PaymentMethods  any `json:"payment_methods,omitempty"`
	WebhookURLs     any `json:"webhook_urls,omitempty"`
}

func (s *ProductsService) Update(ctx context.Context, productID string, req UpdateProductRequest) (*GetProductResponse, *core.ResponseMeta, error) {
	var out GetProductResponse
	meta, err := s.c.Do(ctx, "PATCH", "/products/"+productID, nil, req, &out)
	return &out, meta, err
}

type DeleteProductResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Deleted     bool   `json:"deleted"`
		ProductID   string `json:"product_id"`
		ProductName string `json:"product_name,omitempty"`
	} `json:"data"`
}

func (s *ProductsService) Delete(ctx context.Context, productID string) (*DeleteProductResponse, *core.ResponseMeta, error) {
	var out DeleteProductResponse
	meta, err := s.c.Do(ctx, "DELETE", "/products/"+productID, nil, nil, &out)
	return &out, meta, err
}
