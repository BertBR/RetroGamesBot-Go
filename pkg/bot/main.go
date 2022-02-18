package bot

import (
	"fmt"
	"log"
	"os"

	tb "gopkg.in/telebot.v3"
)

func New() {
	webhookUrl := os.Getenv("WEBHOOK_URL")
	botToken := os.Getenv("BOT_TOKEN")
	port := os.Getenv("PORT")

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: webhookUrl},
	}

	pref := tb.Settings{
		Token:  botToken,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatalln(err)
	}

	b.Handle("/start", func(c tb.Context) error {
		return c.Reply(fmt.Sprintf("Welcome, %s !!!", c.Message().Sender.FirstName))
	})

	b.Handle("/games", func(c tb.Context) error {
		msg := handleTopGames(c.Message().Sender.FirstName)
		return c.Reply(msg)
	})

	b.Start()
}
