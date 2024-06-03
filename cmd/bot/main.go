package bot

import (
	"bytes"
	"embed"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/BertBR/RetroGamesBot-Go/cmd/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/robfig/cron/v3"
	tb "gopkg.in/telebot.v3"
)

var (
	//go:embed templates
	files     embed.FS
	templates = map[string]string{
		"/count":    "templates/totalGames.html",
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

	fmt.Printf("Bot started at: %s%s\n", webhook.Endpoint.PublicURL, webhook.Listen)

	bot := NewTelebot(pool)
	cr := cron.New()
	// It runs every saturday at 00:00
	cr.AddFunc("0 0 * * 6", func() {
		svc := service.New(pool)
		bot.SortThreeRandomGames(svc, b)
	})

	cr.Start()

	b.Handle("/start", bot.Start)
	b.Handle("/count", bot.Count)
	b.Handle("/games", bot.Games)
	b.Handle("/consoles", bot.Consoles)
	b.Handle("/genres", bot.Genres)
	b.Handle("/sort", bot.Sort)
	b.Start()
}

func parseTemplate(file []byte, name string, data interface{}) (string, error) {
	t, err := template.New("index").Parse(string(file))
	if err != nil {
		log.Fatalln("error parsing template", err)
		return "", err
	}
	var buf bytes.Buffer
	numbers := []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣", "🔟"}
	err = t.Execute(&buf, struct {
		Name    string
		Data    interface{}
		Numbers []string
	}{name, data, numbers})
	if err != nil {
		log.Fatalln("error executing template", err)
		return "", err
	}
	return buf.String(), nil
}
