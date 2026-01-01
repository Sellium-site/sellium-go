package sellium

import (
	"github.com/Sellium-site/sellium-go/core"
	"github.com/Sellium-site/sellium-go/services"
)

type Client struct {
	core *core.Client

	Store     *services.StoreService
	Products  *services.ProductsService
	Coupons   *services.CouponsService
	Orders    *services.OrdersService
	Customers *services.CustomersService
	Feedback  *services.FeedbackService
	Tickets   *services.TicketsService
	Blacklist *services.BlacklistService
	Groups    *services.GroupsService
}

func (c *Client) Core() *core.Client { return c.core }

// Models
type (
	Store            = core.Store
	StoreStats       = core.StoreStats
	Product          = core.Product
	Coupon           = core.Coupon
	CouponAnalytics  = core.CouponAnalytics
	Order            = core.Order
	OrderProductMini = core.OrderProductMini

	CustomerRow         = core.CustomerRow
	CustomerDetail      = core.CustomerDetail
	CustomerStats       = core.CustomerStats
	CustomerTopProduct  = core.CustomerTopProduct
	CustomerRecentOrder = core.CustomerRecentOrder

	Feedback = core.Feedback

	Ticket             = core.Ticket
	TicketMessage      = core.TicketMessage
	TicketOrderSummary = core.TicketOrderSummary

	BlacklistEntry = core.BlacklistEntry

	Group            = core.Group
	GroupProductMini = core.GroupProductMini

	Pagination = core.Pagination
)

type (
	ResponseMeta = core.ResponseMeta
	APIError     = core.APIError
)

type Option = core.Option

var (
	WithBaseURL    = core.WithBaseURL
	WithHTTPClient = core.WithHTTPClient
	WithUserAgent  = core.WithUserAgent
)

type (
	// Query params
	ListProductsParams  = services.ListProductsParams
	ListCouponsParams   = services.ListCouponsParams
	ListOrdersParams    = services.ListOrdersParams
	ListCustomersParams = services.ListCustomersParams
	ListFeedbackParams  = services.ListFeedbackParams
	ListTicketsParams   = services.ListTicketsParams
	ListBlacklistParams = services.ListBlacklistParams
	ListGroupsParams    = services.ListGroupsParams

	// Products
	CreateProductRequest = services.CreateProductRequest
	UpdateProductRequest = services.UpdateProductRequest

	// Coupons
	CreateCouponRequest = services.CreateCouponRequest
	UpdateCouponRequest = services.UpdateCouponRequest

	// Orders
	CreateOrderRequest = services.CreateOrderRequest
	UpdateOrderRequest = services.UpdateOrderRequest

	// Feedback
	UpdateFeedbackRequest = services.UpdateFeedbackRequest

	// Tickets
	ReplyTicketRequest  = services.ReplyTicketRequest
	UpdateTicketRequest = services.UpdateTicketRequest

	// Blacklist
	CreateBlacklistEntryRequest = services.CreateBlacklistEntryRequest

	// Groups
	CreateGroupRequest = services.CreateGroupRequest
	UpdateGroupRequest = services.UpdateGroupRequest
)

func NewClient(apiKey, storeID string, opts ...Option) *Client {
	cc := core.New(apiKey, storeID, opts...)

	return &Client{
		core:      cc,
		Store:     services.NewStore(cc),
		Products:  services.NewProducts(cc),
		Coupons:   services.NewCoupons(cc),
		Orders:    services.NewOrders(cc),
		Customers: services.NewCustomers(cc),
		Feedback:  services.NewFeedback(cc),
		Tickets:   services.NewTickets(cc),
		Blacklist: services.NewBlacklist(cc),
		Groups:    services.NewGroups(cc),
	}
}
