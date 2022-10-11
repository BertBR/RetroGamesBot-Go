package bot

import (
	"context"
	"fmt"
	"log"

	"github.com/BertBR/RetroGamesBot-Go/cmd/service"
	"github.com/BertBR/RetroGamesBot-Go/pkg/storage/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
	tb "gopkg.in/telebot.v3"
)

type Telebot struct {
	svc *service.Service
}

func NewTelebot(pool *pgxpool.Pool) *Telebot {
	return &Telebot{
		svc: service.New(pool),
	}
}

func (t *Telebot) Start(c tb.Context) error {
	return c.Reply(fmt.Sprintf("Welcome, %s !!!", c.Message().Sender.FirstName))
}

func (t *Telebot) Count(c tb.Context) error {
	ctx := context.Background()
	totalGames, err := t.svc.GetTotalGames(ctx)
	if err != nil {
		return err
	}
	totalByConsole, err := t.svc.GetTotalGamesByConsole(ctx)
	if err != nil {
		return err
	}

	file, err := files.ReadFile(templates[c.Message().Text])
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
}

func (t *Telebot) Games(c tb.Context) error {
	ctx := context.Background()
	top10Games, err := t.svc.GetTop10Games(ctx)
	if err != nil {
		return err
	}
	file, err := files.ReadFile(templates[c.Message().Text])
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
}

func (t *Telebot) Consoles(c tb.Context) error {
	ctx := context.Background()
	top10Consoles, err := t.svc.GetTop10Console(ctx)
	if err != nil {
		return err
	}
	file, err := files.ReadFile(templates[c.Message().Text])
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
}

func (t *Telebot) Genres(c tb.Context) error {
	ctx := context.Background()
	top10Genres, err := t.svc.GetTop10Genre(ctx)
	if err != nil {
		return err
	}
	file, err := files.ReadFile(templates[c.Message().Text])
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
}
