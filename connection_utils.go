package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

type DBConstructor struct {
	DBurl  string
	Config string
}

type DBUtils struct {
	opt    option.ClientOption
	app    *firebase.App
	client *db.Client
}

func (dbc DBConstructor) make() DBUtils {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: dbc.DBurl,
	}

	// fetch service account key
	opt := option.WithCredentialsFile(dbc.Config)

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("error in creating firebase DB client: ", err)
	}

	return DBUtils{opt: opt, app: app, client: client}
}
