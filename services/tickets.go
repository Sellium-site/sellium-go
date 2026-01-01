package services

import (
	"context"
	"net/url"
	"strconv"

	"github.com/Sellium-site/sellium-go/core"
)

type TicketsService struct{ c *core.Client }

func NewTickets(c *core.Client) *TicketsService { return &TicketsService{c: c} }

type ListTicketsParams struct {
	Page     int
	Limit    int
	Status   string
	Priority string
	Email    string
}

type ListTicketsResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Tickets    []core.Ticket   `json:"tickets"`
		Pagination core.Pagination `json:"pagination"`
	} `json:"data"`
}

func (s *TicketsService) List(ctx context.Context, p *ListTicketsParams) (*ListTicketsResponse, *core.ResponseMeta, error) {
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
		if p.Priority != "" {
			q.Set("priority", p.Priority)
		}
		if p.Email != "" {
			q.Set("email", p.Email)
		}
	}

	var out ListTicketsResponse
	meta, err := s.c.Do(ctx, "GET", "/tickets", q, nil, &out)
	return &out, meta, err
}

type GetTicketResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Ticket   core.Ticket          `json:"ticket"`
		Messages []core.TicketMessage `json:"messages"`
	} `json:"data"`
}

func (s *TicketsService) Get(ctx context.Context, ticketID string) (*GetTicketResponse, *core.ResponseMeta, error) {
	var out GetTicketResponse
	meta, err := s.c.Do(ctx, "GET", "/tickets/"+ticketID, nil, nil, &out)
	return &out, meta, err
}

type ReplyTicketRequest struct {
	Message string `json:"message"`
	Status  string `json:"status,omitempty"` // optional
}

type ReplyTicketResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Message      core.TicketMessage `json:"message"`
		TicketStatus string             `json:"ticket_status"`
	} `json:"data"`
}

func (s *TicketsService) Reply(ctx context.Context, ticketID string, req ReplyTicketRequest) (*ReplyTicketResponse, *core.ResponseMeta, error) {
	var out ReplyTicketResponse
	meta, err := s.c.Do(ctx, "POST", "/tickets/"+ticketID+"/reply", nil, req, &out)
	return &out, meta, err
}

type UpdateTicketRequest struct {
	Status   *string `json:"status,omitempty"`
	Priority *string `json:"priority,omitempty"`
}

type UpdateTicketResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Ticket core.Ticket `json:"ticket"`
	} `json:"data"`
}

func (s *TicketsService) Update(ctx context.Context, ticketID string, req UpdateTicketRequest) (*UpdateTicketResponse, *core.ResponseMeta, error) {
	var out UpdateTicketResponse
	meta, err := s.c.Do(ctx, "PATCH", "/tickets/"+ticketID, nil, req, &out)
	return &out, meta, err
}
