package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"modulo-1-go/sqlDB"
	"net/http"
	"time"
)

type Exchange struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func ExchangeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	apiCtx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(apiCtx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if apiCtx.Err() == context.DeadlineExceeded {
			http.Error(w, "API request timeout (over 200ms)", http.StatusGatewayTimeout)
			return
		}
		http.Error(w, "Failed to fetch exchange", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var exchange Exchange
	err = json.NewDecoder(resp.Body).Decode(&exchange)
	if err != nil {
		panic(err)
	}

	log.Printf("bid %s", exchange.USDBRL.Bid)

	log.Println("Request canceled by the client")
	w.WriteHeader(499)
	err = sqlDB.RegisterExchange(exchange.USDBRL.Bid)
	if err != nil {
		fmt.Errorf("Failed to register exchange, %w", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exchange.USDBRL.Bid)
}
