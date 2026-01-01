# sellium-go

**sellium-go** is the official, production-ready Go (Golang) SDK for the Sellium API.  
It provides fully typed access to all Sellium endpoints, including stores, products, orders, coupons, customers, tickets, feedback, groups, and blacklist management.

The SDK is designed to be clean, safe, and easy to use, with a single client instance and logically grouped services.

---

## Features

- Full coverage of **all Sellium API v1 endpoints**
- Strongly typed request and response models
- Clean service-based API (`Products`, `Orders`, `Coupons`, etc.)
- Built-in authentication via API key and Store ID
- Automatic error handling with typed API errors
- Pagination support on all list endpoints
- Rate limit metadata access
- No external dependencies

---

## Requirements

- Go 1.22 or newer

---

## Installation

```bash
go get github.com/Sellium-site/sellium-go
```

---

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Sellium-site/sellium-go"
)

func main() {
	client := sellium.NewClient(
		"YOUR_API_KEY",
		"YOUR_STORE_ID",
		sellium.WithBaseURL("https://yourdomain.com/api/v1"),
	)

	ctx := context.Background()

	products, _, err := client.Products.List(ctx, &sellium.ListProductsParams{
		Page:  1,
		Limit: 20,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Products:", len(products.Data.Products))
}
```

---

## Authentication

All requests are authenticated using:

- `X-API-Key`
- `X-Store-ID`

These are automatically added by the client when you create it using `NewClient`.

---

## Client Configuration

```go
client := sellium.NewClient(
	"API_KEY",
	"STORE_ID",
	sellium.WithBaseURL("https://yourdomain.com/api/v1"),
	sellium.WithUserAgent("my-app/1.0"),
)
```

Available options:

- `WithBaseURL(string)`
- `WithUserAgent(string)`
- `WithHTTPClient(*http.Client)`

---

## Services Overview

The SDK exposes one service per API group:

- `client.Store`
- `client.Products`
- `client.Orders`
- `client.Coupons`
- `client.Customers`
- `client.Feedback`
- `client.Tickets`
- `client.Blacklist`
- `client.Groups`

Each service contains methods that map directly to the Sellium API endpoints.

---

## Examples

### List Products

```go
products, meta, err := client.Products.List(ctx, &sellium.ListProductsParams{
	Page:  1,
	Limit: 10,
})
```

### Create a Product

```go
product, _, err := client.Products.Create(ctx, sellium.CreateProductRequest{
	Name:         "Example Product",
	PriceInCents: 999,
	DeliveryType: "file",
})
```

### Create an Order

```go
order, _, err := client.Orders.Create(ctx, sellium.CreateOrderRequest{
	ProductID:     "product_id",
	CustomerEmail: "customer@example.com",
	Quantity:      1,
})
```

### Reply to a Ticket

```go
_, _, err := client.Tickets.Reply(ctx, "ticket_id", sellium.ReplyTicketRequest{
	Message: "Thanks for contacting support!",
})
```

---

## Pagination

All list endpoints return pagination data:

```go
pagination := response.Data.Pagination
fmt.Println(pagination.Page, pagination.TotalPages)
```

---

## Error Handling

API errors are returned as `*sellium.APIError`:

```go
if err != nil {
	if apiErr, ok := err.(*sellium.APIError); ok {
		fmt.Println(apiErr.Status, apiErr.Code, apiErr.Message)
	}
}
```

---

## Rate Limiting

Rate limit data (if provided by the API) is available via `ResponseMeta`:

```go
_, meta, _ := client.Products.List(ctx, &sellium.ListProductsParams{})
if meta != nil && meta.RateLimit != nil {
	fmt.Println("Remaining:", meta.RateLimit.Remaining)
}
```

---

## Build & Verify

From the repository root:

```bash
gofmt -w .
go test ./...
```

If you don’t have tests yet, you can still run a compile check:

```bash
go test ./... -run TestDoesNotExist
```

---

## Project Structure

```text
sellium-go/
├── core/        # Low-level HTTP client, models, errors
├── services/    # API endpoint groups
├── examples/    # Usage examples
└── sellium.go   # Public SDK entry point
```

---

## License

MIT License

---

## Contributing

Pull requests are welcome.  
Please ensure all code is formatted with `gofmt` and passes `go test ./...` before submitting.

---

## Support

For questions or issues related to the Sellium API, please refer to the official Sellium documentation or open an issue in this repository.
