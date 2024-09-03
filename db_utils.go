package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

type Config struct {
	DBurl        string `json:DatabaseURL`
	AuthOverride string `json:AuthOveride`
}

type DBConstructor struct {
	DBurl        string
	ConfigPath   string
	AuthOverride string
}

type DBUtils struct {
	opt    option.ClientOption
	app    *firebase.App
	client *db.Client
}

func (dbc *DBConstructor) make() *DBUtils {
	ctx := context.Background()

	ao := map[string]interface{}{"uid": dbc.AuthOverride}

	conf := &firebase.Config{
		DatabaseURL:  dbc.DBurl,
		AuthOverride: &ao,
	}

	// fetch service account key
	opt := option.WithCredentialsFile(dbc.ConfigPath)

	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		log.Fatalln("error in creating firebase DB client: ", err)
	}

	return &DBUtils{opt: opt, app: app, client: client}
}
