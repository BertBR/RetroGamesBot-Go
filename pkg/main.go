package main

import (
	"log"
	"os"

	"github.com/BertBR/RetroGamesBot-Go/pkg/bot"
	"github.com/BertBR/RetroGamesBot-Go/pkg/database"
	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
	pid := os.Getenv("PROJECT_ID")
	err := database.Connect(pid)
	if err != nil {
		log.Fatalf("error initializing firestore: %v\n", err)
	}
}

func main() {
	bot.New()
}

func loadEnv() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
