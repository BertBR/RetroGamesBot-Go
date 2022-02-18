package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type Game struct {
	Active   bool   `firestore:"active"`
	Console  string `firestore:"console"`
	FileURL  string `firestore:"file_url"`
	Genre    string `firestore:"genre"`
	ImageURL string `firestore:"image_url"`
	Sorted   int64  `firestore:"sorted"`
	Title    string `firestore:"title"`
}

var f *firestore.Client

func Connect(pid string) error {
	ctx := context.Background()
	app := startApp(pid)
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	f = client
	return nil
}

func startApp(pid string) *firebase.App {
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: pid}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return app

}

func LoadGames(ctx context.Context) []*firestore.DocumentSnapshot {
	snap, err := f.Collection("games").Where("sorted", ">", 0).OrderBy("sorted", firestore.Desc).Limit(10).Documents(ctx).GetAll()
	if err != nil {
		log.Fatalf("error while getting documents: %v\n", err)
	}

	return snap
}
