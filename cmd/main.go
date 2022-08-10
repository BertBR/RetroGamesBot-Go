package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/BertBR/RetroGamesBot-Go/cmd/bot"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
	pool := run()
	if err != nil {
		log.Fatal(err)
	}
	bot.New(pool)
}

func run() *pgxpool.Pool {
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return pool
}
