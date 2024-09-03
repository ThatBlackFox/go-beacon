package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()

	// configure database URL
	conf := &firebase.Config{
		DatabaseURL: "",
	}

	// fetch service account key
	opt := option.WithCredentialsFile("creds.json")

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("error in creating firebase DB client: ", err)
	}

	// create ref at path user_scores/:userId
	ref := client.NewRef("user_scores/" + fmt.Sprint(1))

	if err := ref.Set(context.TODO(), map[string]interface{}{"score": 40}); err != nil {
		log.Fatal(err)
	}

	fmt.Println("score added/updated successfully!")

	type UserScore struct {
		Score int `json:"score"`
	}

	// get database reference to user score
	ref = client.NewRef("user_scores/1")

	// read from user_scores using ref
	var s UserScore
	if err := ref.Get(context.TODO(), &s); err != nil {
		log.Fatalln("error in reading from firebase DB: ", err)
	}
	fmt.Println("retrieved user's score is: ", s.Score)
}
