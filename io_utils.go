package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func read_config() *Config {
	file, err := os.Open("dbConf.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatal(err)
	}

	return &config
}
