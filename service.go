package main

import (
	"context"
	"fmt"
)

// An interface for fetching a price
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// priceFetcher is a concrete implementation of the PriceFetcher interface
type priceFetcher struct {
}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var priceMocks = map[string]float64{
	"BTC": 10000,
	"ETH": 500,
	"SOL": 200,
}

func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMocks[ticker]

	if !ok {
		return price, fmt.Errorf("price not found with ticker %s", ticker)
	}

	return price, nil
}
