package core

import "encoding/json"

type Socials struct {
	Twitter   string `json:"twitter,omitempty"`
	Instagram string `json:"instagram,omitempty"`
	Youtube   string `json:"youtube,omitempty"`
	Discord   string `json:"discord,omitempty"`
	Telegram  string `json:"telegram,omitempty"`
	Tiktok    string `json:"tiktok,omitempty"`
}

type Store struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Slug         string  `json:"slug"`
	Description  string  `json:"description,omitempty"`
	LogoURL      string  `json:"logo_url,omitempty"`
	CustomDomain string  `json:"custom_domain,omitempty"`
	ThemeColor   string  `json:"theme_color,omitempty"`
	SupportEmail string  `json:"support_email,omitempty"`
	IsActive     bool    `json:"is_active"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	URL          string  `json:"url,omitempty"`
	Socials      Socials `json:"socials,omitempty"`
}

type StoreStats struct {
	TotalSales        int     `json:"total_sales"`
	TotalRevenueCents int     `json:"total_revenue_cents"`
	TotalReviews      int     `json:"total_reviews"`
	AverageRating     float64 `json:"average_rating"`
	ProductCount      int     `json:"product_count"`
	CompletedOrders   int     `json:"completed_orders"`
}

type Product struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description,omitempty"`
	ImageURL      string `json:"image_url,omitempty"`
	PriceInCents  int    `json:"price_in_cents"`
	IsActive      bool   `json:"is_active"`
	StockQuantity int    `json:"stock_quantity"`
	DeliveryType  string `json:"delivery_type"`

	MinimumQuantity int    `json:"minimum_quantity,omitempty"`
	MaximumQuantity int    `json:"maximum_quantity,omitempty"`
	Unlisted        bool   `json:"unlisted,omitempty"`
	IsPrivate       bool   `json:"is_private,omitempty"`
	OnHold          bool   `json:"on_hold,omitempty"`
	Warranty        string `json:"warranty,omitempty"`
	ProductTerms    string `json:"product_terms,omitempty"`

	GroupID string `json:"group_id,omitempty"`

	FileURL           string   `json:"file_url,omitempty"`
	Serials           []string `json:"serials,omitempty"`
	ServiceMessage    string   `json:"service_message,omitempty"`
	DeliveryText      string   `json:"delivery_text,omitempty"`
	DynamicWebhookURL string   `json:"dynamic_webhook_url,omitempty"`
	RedirectURL       string   `json:"redirect_url,omitempty"`
	YoutubeURL        string   `json:"youtube_url,omitempty"`

	DiscordEnabled  bool `json:"discord_enabled,omitempty"`
	DiscordOptional bool `json:"discord_optional,omitempty"`

	EnableLicenseSystem bool `json:"enable_license_system,omitempty"`
	LicenseMaxDevices   int  `json:"license_max_devices,omitempty"`
	LicenseExpiresDays  *int `json:"license_expires_days,omitempty"`

	CustomFields    json.RawMessage `json:"custom_fields,omitempty"`
	VolumeDiscounts json.RawMessage `json:"volume_discounts,omitempty"`
	PaymentMethods  json.RawMessage `json:"payment_methods,omitempty"`
	WebhookURLs     json.RawMessage `json:"webhook_urls,omitempty"`

	AvailableStock int `json:"available_stock,omitempty"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CouponAnalytics struct {
	TotalUses         int  `json:"total_uses"`
	TotalRevenueCents int  `json:"total_revenue_cents"`
	RemainingUses     int  `json:"remaining_uses"`
	IsExpired         bool `json:"is_expired"`
}

type Coupon struct {
	ID              string  `json:"id"`
	Code            string  `json:"code"`
	Type            string  `json:"type"`
	Value           int     `json:"value"`
	MinimumPurchase *int    `json:"minimum_purchase,omitempty"`
	MaximumUses     *int    `json:"maximum_uses,omitempty"`
	UsesCount       int     `json:"uses_count"`
	IsActive        bool    `json:"is_active"`
	ExpiresAt       *string `json:"expires_at,omitempty"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`

	Analytics *CouponAnalytics `json:"analytics,omitempty"`
}

type OrderProductMini struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	PriceInCents int    `json:"price_in_cents"`
	DeliveryType string `json:"delivery_type,omitempty"`
	ImageURL     string `json:"image_url,omitempty"`
	ProductName  string `json:"product_name,omitempty"`
}

type Order struct {
	ID            string `json:"id"`
	CustomerEmail string `json:"customer_email"`
	CustomerName  string `json:"customer_name,omitempty"`

	Status        string `json:"status"`
	AmountInCents int    `json:"amount_in_cents"`
	Quantity      int    `json:"quantity"`

	PaymentMethod   string `json:"payment_method,omitempty"`
	PaymentVerified bool   `json:"payment_verified,omitempty"`
	Delivered       bool   `json:"delivered,omitempty"`

	CheckoutURL     string `json:"checkout_url,omitempty"`
	TransactionID   string `json:"transaction_id,omitempty"`
	DeliveryContent string `json:"delivery_content,omitempty"`

	PaymentCurrency string `json:"payment_currency,omitempty"`
	PaymentAmount   string `json:"payment_amount,omitempty"`
	Country         string `json:"country,omitempty"`
	City            string `json:"city,omitempty"`
	Region          string `json:"region,omitempty"`
	DeviceType      string `json:"device_type,omitempty"`
	Browser         string `json:"browser,omitempty"`
	OS              string `json:"os,omitempty"`
	Referrer        string `json:"referrer,omitempty"`
	AffiliateCode   string `json:"affiliate_code,omitempty"`

	CustomFields json.RawMessage `json:"custom_fields,omitempty"`

	CreatedAt string `json:"created_at"`

	Product OrderProductMini `json:"product"`
}

type Pagination struct {
	Page       int  `json:"page"`
	Limit      int  `json:"limit"`
	Total      int  `json:"total"`
	TotalPages int  `json:"total_pages"`
	HasMore    bool `json:"has_more,omitempty"`
}

type CustomerRow struct {
	Email               string `json:"email"`
	Name                string `json:"name,omitempty"`
	TotalOrders         int    `json:"total_orders"`
	CompletedOrders     int    `json:"completed_orders"`
	TotalSpentCents     int    `json:"total_spent_cents"`
	TotalSpentFormatted string `json:"total_spent_formatted"`
	FirstOrderAt        string `json:"first_order_at,omitempty"`
	LastOrderAt         string `json:"last_order_at,omitempty"`
}

type CustomerStats struct {
	TotalOrders                int    `json:"total_orders"`
	CompletedOrders            int    `json:"completed_orders"`
	PendingOrders              int    `json:"pending_orders"`
	CanceledOrders             int    `json:"canceled_orders"`
	RefundedOrders             int    `json:"refunded_orders"`
	TotalSpentCents            int    `json:"total_spent_cents"`
	TotalSpentFormatted        string `json:"total_spent_formatted"`
	AverageOrderValueCents     int    `json:"average_order_value_cents"`
	AverageOrderValueFormatted string `json:"average_order_value_formatted"`
}

type CustomerTopProduct struct {
	ProductID           string `json:"product_id"`
	ProductName         string `json:"product_name"`
	QuantityPurchased   int    `json:"quantity_purchased"`
	TotalSpentCents     int    `json:"total_spent_cents"`
	TotalSpentFormatted string `json:"total_spent_formatted"`
}

type CustomerDetail struct {
	Email              string               `json:"email"`
	Name               string               `json:"name,omitempty"`
	Stats              CustomerStats        `json:"stats"`
	FirstOrderAt       string               `json:"first_order_at,omitempty"`
	LastOrderAt        string               `json:"last_order_at,omitempty"`
	PaymentMethodsUsed []string             `json:"payment_methods_used,omitempty"`
	TopProducts        []CustomerTopProduct `json:"top_products,omitempty"`
}

type CustomerRecentOrder struct {
	ID              string `json:"id"`
	Status          string `json:"status"`
	AmountInCents   int    `json:"amount_in_cents"`
	AmountFormatted string `json:"amount_formatted,omitempty"`
	Quantity        int    `json:"quantity"`
	PaymentMethod   string `json:"payment_method,omitempty"`
	Product         struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"product"`
	CreatedAt string `json:"created_at"`
}

type Feedback struct {
	ID            string  `json:"id"`
	CustomerEmail string  `json:"customer_email"`
	CustomerName  string  `json:"customer_name,omitempty"`
	Message       string  `json:"message"`
	Response      *string `json:"response,omitempty"`
	Rating        int     `json:"rating"`
	IsVisible     bool    `json:"is_visible"`
	RespondedAt   *string `json:"responded_at,omitempty"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     *string `json:"updated_at,omitempty"`
	OrderID       string  `json:"order_id,omitempty"`

	// single feedback includes order/product info
	Order *struct {
		ID            string `json:"id"`
		AmountInCents int    `json:"amount_in_cents"`
		Status        string `json:"status"`
		CreatedAt     string `json:"created_at"`
		Product       struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"product"`
	} `json:"order,omitempty"`
}

type TicketOrderSummary struct {
	ID            string `json:"id"`
	AmountInCents int    `json:"amount_in_cents"`
	Status        string `json:"status"`
	ProductName   string `json:"product_name,omitempty"`
}

type Ticket struct {
	ID            string              `json:"id"`
	Subject       string              `json:"subject"`
	Status        string              `json:"status"`
	Priority      string              `json:"priority"`
	CustomerEmail string              `json:"customer_email"`
	CustomerName  string              `json:"customer_name,omitempty"`
	MessageCount  int                 `json:"message_count,omitempty"`
	Order         *TicketOrderSummary `json:"order,omitempty"`
	CreatedAt     string              `json:"created_at"`
	UpdatedAt     string              `json:"updated_at"`
	ClosedAt      *string             `json:"closed_at,omitempty"`
}

type TicketMessage struct {
	ID          string `json:"id"`
	TicketID    string `json:"ticket_id,omitempty"`
	Message     string `json:"message"`
	SenderType  string `json:"sender_type"`
	SenderEmail string `json:"sender_email"`
	CreatedAt   string `json:"created_at"`
}

type BlacklistEntry struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Value     string `json:"value"`
	Reason    string `json:"reason,omitempty"`
	CreatedAt string `json:"created_at"`
}

type Group struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description,omitempty"`
	ImageURL     *string `json:"image_url,omitempty"`
	DisplayOrder int     `json:"display_order"`
	IsActive     bool    `json:"is_active"`
	ProductCount int     `json:"product_count"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}

type GroupProductMini struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	PriceInCents  int    `json:"price_in_cents"`
	IsActive      bool   `json:"is_active"`
	StockQuantity int    `json:"stock_quantity"`
	CreatedAt     string `json:"created_at"`
}
