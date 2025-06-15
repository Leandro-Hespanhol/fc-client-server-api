package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
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
	createFile(exchangeBid)

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
	result, err := db.Exec("INSERT INTO exchanges (bid) VALUES (?)", exchange.Bid)
	if err != nil {
		panic(err)
	}

	id, _ := result.LastInsertId()

	log.Printf("Exchange registered successfully, ID: %d", id)
}

func createFile(bid string) {
	content := fmt.Sprintf("Dólar: %s", bid)
	err := os.WriteFile("cotacao.txt" , []byte(content), 0666)
	if err != nil {
		log.Print("Faild to create a file with the bid value")
	}
}
