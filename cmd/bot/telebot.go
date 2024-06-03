package bot

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/BertBR/RetroGamesBot-Go/cmd/service"
	"github.com/BertBR/RetroGamesBot-Go/pkg/storage/postgres"
	"github.com/jackc/pgx/v4/pgxpool"
	tb "gopkg.in/telebot.v3"
)

var BOT_NAME = os.Getenv("BOT_NAME")

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

	file, err := files.ReadFile(templates[strings.TrimSuffix(c.Message().Text, BOT_NAME)])
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
	file, err := files.ReadFile(templates[strings.TrimSuffix(c.Message().Text, BOT_NAME)])
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
	file, err := files.ReadFile(templates[strings.TrimSuffix(c.Message().Text, BOT_NAME)])
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
	file, err := files.ReadFile(templates[strings.TrimSuffix(c.Message().Text, BOT_NAME)])
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

// Function to check if a value is present in a list
func contains(list []int64, value int64) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func (t *Telebot) Sort(c tb.Context) error {
	list := strings.Split(os.Getenv("ALLOWED_LIST"), ",")
	//convert list to int64[]
	listInt64 := make([]int64, len(list))
	for i, v := range list {
		num, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			fmt.Println("Error converting string to int64:", err)
			continue
		}
		listInt64[i] = num
	}

	senderID := c.Message().Sender.ID
	// Check if the sender ID is in the list
	if !contains(listInt64, senderID) {
		return c.Reply("You are not allowed to use this command")
	}

	// Printing each int64 value
	for _, v := range listInt64 {
		fmt.Println(v)
	}

	t.SortThreeRandomGames(t.svc, c.Bot())
	return c.Reply("Sorting three random games")
}

func (t *Telebot) SortThreeRandomGames(svc *service.Service, b *tb.Bot) {
	chatId, err := strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

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
	chat, err := b.ChatByID(chatId)
	if err != nil {
		log.Fatalln("error getting chat ID: ", err)
	}

	msgs, err := b.SendAlbum(chat, a, tb.ModeMarkdown, tb.NoPreview)
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
}
