package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"modulo-1-go/db"
	"net/http"
	"os"
	"time"
)

type Exchange struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/exchange", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	var exchangeBid string
	json.NewDecoder(res.Body).Decode(&exchangeBid)
	registerExchange(exchangeBid)

	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}

func NewExchange(exchange string) Exchange {
	return Exchange{
		Bid: exchange,
	}
}

func registerExchange(exchangeBid string) {
	db := db.StartMySQL()
	defer db.Close()

	exchange := NewExchange(exchangeBid)
	result, err := db.Exec("INSERT INTO exchange (bid) VALUES (?)", exchange.Bid)
	if err != nil {
		panic(err)
	}

	log.Println(result)
	log.Println("Exchange registered successfully")
}
