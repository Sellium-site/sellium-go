package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Sellium-site/sellium-go"
)

func main() {
	c := sellium.NewClient(
		"YOUR_API_KEY",
		"YOUR_STORE_ID",
		sellium.WithBaseURL("https://sellium.site/api/v1"),
	)
	ctx := context.Background()

	store, meta, err := c.Store.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Store:", store.Data.Store.Name)

	if meta != nil && meta.RateLimit != nil {
		fmt.Println("Rate remaining:", meta.RateLimit.Remaining)
	}

	products, _, err := c.Products.List(ctx, &sellium.ListProductsParams{
		Page:  1,
		Limit: 20,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Products:", len(products.Data.Products))

	orders, _, err := c.Orders.List(ctx, &sellium.ListOrdersParams{
		Page:  1,
		Limit: 20,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Orders:", len(orders.Data.Orders))

	coupons, _, err := c.Coupons.List(ctx, &sellium.ListCouponsParams{
		Page:  1,
		Limit: 20,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Coupons:", len(coupons.Data.Coupons))

	customers, _, err := c.Customers.List(ctx, &sellium.ListCustomersParams{
		Page:  1,
		Limit: 20,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Customers:", len(customers.Data.Customers))
}
