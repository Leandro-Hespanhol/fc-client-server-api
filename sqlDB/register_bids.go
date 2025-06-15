package sqlDB

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Exchange struct {
	Bid string `json:"bid"`
}

func NewExchange(exchange string) Exchange {
	return Exchange{
		Bid: exchange,
	}
}

func RegisterExchange(exchangeBid string) error {
	db := StartMySQL()
	defer db.Close()

	exchange := NewExchange(exchangeBid)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	id, err := execute(ctx, db, exchange.Bid)
	log.Printf("Exchange registered successfully, ID: %d", id)

	return err
}

func execute(ctx context.Context, db *sql.DB, bid string) (int64, error) {
	result, err := db.Exec("INSERT INTO exchanges (bid) VALUES (?)", bid)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	select {
	case <-ctx.Done():
		return id, fmt.Errorf("")
	case <-time.After(1 * time.Second):
		return id, err
	}
}
