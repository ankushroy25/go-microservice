package main

import (
	"context"
	"fmt"
)

type metricService struct {
	next PriceFetcher
}

func NewMetricService(next PriceFetcher) PriceFetcher {
	return &metricService{next: next}
}

func (s *metricService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {

	// metrics storage
	fmt.Println("Pushing metrics to monitoring system")

	return s.next.FetchPrice(ctx, ticker)
}
