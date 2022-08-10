package bot

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"log"
	"os"
	"text/template"

	_ "embed"

	"github.com/BertBR/RetroGamesBot-Go/cmd/service"
	"github.com/jackc/pgx/v4/pgxpool"
	tb "gopkg.in/telebot.v3"
)

var (
	//go:embed templates/*.html
	files embed.FS
	t     = map[string]string{
		"/consoles": "templates/top10Consoles.html",
		"/genres":   "templates/top10Genres.html",
		"/games":    "templates/top10Games.html",
	}
)

func New(pool *pgxpool.Pool) {
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
		svc := service.New(pool)
		ctx := context.Background()
		top10Games, err := svc.GetTop10Games(ctx)
		if err != nil {
			return err
		}
		file, err := files.ReadFile(t[c.Message().Text])
		if err != nil {
			log.Fatalln("error reading file", err)
			return err
		}
		s, err := parseTemplate(file, top10Games)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		return c.Reply(s, tb.ModeMarkdown, tb.NoPreview)
	})

	b.Handle("/consoles", func(c tb.Context) error {
		svc := service.New(pool)
		ctx := context.Background()
		top10Consoles, err := svc.GetTop10Console(ctx)
		if err != nil {
			return err
		}
		file, err := files.ReadFile(t[c.Message().Text])
		if err != nil {
			log.Fatalln("error reading file", err)
			return err
		}
		s, err := parseTemplate(file, top10Consoles)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		return c.Reply(s, tb.ModeMarkdown, tb.NoPreview)
	})

	b.Handle("/genres", func(c tb.Context) error {
		svc := service.New(pool)
		ctx := context.Background()
		top10Genres, err := svc.GetTop10Genre(ctx)
		if err != nil {
			return err
		}
		file, err := files.ReadFile(t[c.Message().Text])
		if err != nil {
			log.Fatalln("error reading file", err)
			return err
		}
		s, err := parseTemplate(file, top10Genres)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		return c.Reply(s, tb.ModeMarkdown, tb.NoPreview)
	})
	b.Start()
}

func parseTemplate(file []byte, data interface{}) (string, error) {
	t, err := template.New("index").Parse(string(file))
	if err != nil {
		log.Fatalln("error parsing template", err)
		return "", err
	}
	var buf bytes.Buffer
	numbers := []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣", "🔟"}
	err = t.Execute(&buf, struct {
		Data    interface{}
		Numbers []string
	}{data, numbers})
	if err != nil {
		log.Fatalln("error executing template", err)
		return "", err
	}
	return buf.String(), nil
}
