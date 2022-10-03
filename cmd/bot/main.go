package bot

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"

	"github.com/BertBR/RetroGamesBot-Go/cmd/service"
	"github.com/BertBR/RetroGamesBot-Go/pkg/storage/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/robfig/cron/v3"
	tb "gopkg.in/telebot.v3"
)

var (
	//go:embed templates
	files embed.FS
	t     = map[string]string{
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
	chatId, err := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 32)
	if err != nil {
		log.Fatal(err)
	}

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

	cr := cron.New()
	// It runs every saturday at 00:00
	cr.AddFunc("0 0 * * 6", func() {
		svc := service.New(pool)
		ctx := context.Background()

		threeRandomGames, err := svc.GetThreeRandomGamesRow(ctx)
		if err != nil {
			log.Fatalln("error getting random three games: ", err)
		}

		file, err := files.ReadFile("templates/sortThreeRandomGames.html")
		if err != nil {
			log.Fatalln("error reading file", err)
		}

		s, err := parseTemplate(file, "", threeRandomGames)
		if err != nil {
			log.Fatalln("error parsing template: ", err)
		}

		a := tb.Album{
			&tb.Photo{File: tb.FromURL(threeRandomGames[0].ImageUrl), Caption: s},
			&tb.Photo{File: tb.FromURL(threeRandomGames[1].ImageUrl)},
			&tb.Photo{File: tb.FromURL(threeRandomGames[2].ImageUrl)},
		}
		chatId, err := b.ChatByID(chatId)
		if err != nil {
			log.Fatalln("error getting chat ID: ", err)
		}
		msgs, err := b.SendAlbum(chatId, a, tb.ModeMarkdown, tb.NoPreview)
		if err != nil {
			log.Fatalln("error sending media album: ", err)
		}
		b.Pin(&msgs[0])
		var ids []int32
		for _, v := range threeRandomGames {
			ids = append(ids, v.ID)
		}

		err = svc.UpdateSortedGames(ctx, ids)
		if err != nil {
			log.Fatalln("error updating sorted games: ", err)
		}
	})

	cr.Start()

	b.Handle("/start", func(c tb.Context) error {
		return c.Reply(fmt.Sprintf("Welcome, %s !!!", c.Message().Sender.FirstName))
	})

	b.Handle("/count", func(c tb.Context) error {
		svc := service.New(pool)
		ctx := context.Background()
		totalGames, err := svc.GetTotalGames(ctx)
		if err != nil {
			return err
		}
		totalByConsole, err := svc.GetTotalGamesByConsole(ctx)
		if err != nil {
			return err
		}

		file, err := files.ReadFile(t[c.Message().Text])
		if err != nil {
			log.Fatalln("error reading file", err)
			return err
		}
		data := struct {
			Total int64
			Data  []postgres.GetTotalGamesByConsoleRow
		}{
			Total: totalGames[0],
			Data:  totalByConsole,
		}
		name := c.Message().Sender.FirstName
		s, err := parseTemplate(file, name, data)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		return c.Reply(s, tb.ModeMarkdown, tb.NoPreview)
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
		name := c.Message().Sender.FirstName
		s, err := parseTemplate(file, name, top10Games)
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
		name := c.Message().Sender.FirstName
		s, err := parseTemplate(file, name, top10Consoles)
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
		name := c.Message().Sender.FirstName
		s, err := parseTemplate(file, name, top10Genres)
		if err != nil {
			log.Fatalln(err)
			return err
		}
		return c.Reply(s, tb.ModeMarkdown, tb.NoPreview)
	})
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
