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

	fmt.Printf("Bot started at: %s%s", webhook.Endpoint.PublicURL, webhook.Listen)

	b.Handle("/start", func(c tb.Context) error {
		return c.Reply(fmt.Sprintf("Welcome, %s !!!", c.Message().Sender.FirstName))
	})

	b.Handle("/games", func(c tb.Context) error {

		// msg := game.TopGames()
		// b, err := json.Marshal(msg)
		if err != nil {
			log.Println(err)
		}
		return c.Reply("b")
	})

	b.Start()
}
