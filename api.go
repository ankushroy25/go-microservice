package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/ankushroy25/go-microservice/types"
)

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

func newJsonApiServer(listenAddress string, svc PriceFetcher) *JSONApiServer {
	return &JSONApiServer{
		listenAddress: listenAddress,
		svc:           svc,
	}
}

func (s *JSONApiServer) Run() {
	http.HandleFunc("/", makeHttpHandlerFunc(s.handleFetchPrice))

	http.ListenAndServe(s.listenAddress, nil)
}

func makeHttpHandlerFunc(apiFn APIFunc) http.HandlerFunc {

	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(100000))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			writeJson(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func (s *JSONApiServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	priceResponse := types.PriceResponse{
		Price:  price,
		Ticker: ticker,
	}

	return writeJson(w, http.StatusOK, priceResponse)
}

func writeJson(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
