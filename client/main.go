package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
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

	log.Printf("Bid is %s", exchangeBid)

	createFile(exchangeBid)

	defer res.Body.Close()
	io.Copy(os.Stdout, res.Body)
}

func createFile(bid string) {
	content := fmt.Sprintf("Dólar: %s", bid)
	err := os.WriteFile("cotacao.txt", []byte(content), 0666)
	if err != nil {
		log.Print("Faild to create a file with the bid value")
	}
}
