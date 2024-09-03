package main

import (
	"context"
	"log"
	"strconv"
	"time"
)

func main() {
	//Intialise the rtdb client
	config := read_config()
	dbc := DBConstructor{DBurl: config.DBurl, ConfigPath: "creds.json", AuthOverride: config.AuthOverride}
	utils := dbc.make()

	today := time.Now().Weekday().String()

	//get last update
	var Data lastUpdate
	ref := utils.client.NewRef("/last_update/")
	err := ref.Get(context.Background(), &Data)
	if err != nil {
		log.Fatalf("error getting data: %v", err)
	}
	lastUpdateDay := Data.Weekday
	lastIndex := Data.Index
	log.Println(lastUpdateDay)
	log.Println(today)

	// reset today
	if lastUpdateDay != today {
		delRef := utils.client.NewRef("/updates/" + today)
		delRef.Delete(context.Background())
		lastIndex = 0
		ref.Child("Weekday").Set(context.TODO(), today)
	}

	//send a heartbeat
	newRef := utils.client.NewRef("/updates/" + today)
	newRef.Child(strconv.Itoa(lastIndex)).Set(context.Background(), time.Now().Unix())

	//update index
	ref.Child("Index").Set(context.TODO(), (lastIndex + 1))
}
