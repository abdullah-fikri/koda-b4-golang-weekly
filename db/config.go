package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type Product struct {
	Id         int
	Name       string
	Price      int
	Created_at time.Time
	Update_at  time.Time
}

func ConnectDb() []Product {
	err := godotenv.Load()
	if err != nil {
		log.Println("failed load .env")
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("failed connect database")
	}

	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		"SELECT id, name, price, created_at, update_at FROM products",
	)
	if err != nil {
		fmt.Println("failed get database")
	}

	product, err := pgx.CollectRows(rows, pgx.RowToStructByName[Product])
	if err != nil {
		fmt.Println("failed to map", err)
	}

	return product

}
