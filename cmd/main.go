package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/BertBR/RetroGamesBot-Go/cmd/bot"
	"github.com/BertBR/RetroGamesBot-Go/cmd/service"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
	bot.New()
}

func run() error {
	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()
	svc := service.New(pool)
	top10Console, err := svc.GetTop10Console(ctx)
	if err != nil {
		return err
	}
	top10Genre, err := svc.GetTop10Genre(ctx)
	if err != nil {
		return err
	}
	fmt.Println(top10Genre)
	fmt.Println(top10Console)
	return nil
}

func loadEnv() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
