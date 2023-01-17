package main

import (
	"github.com/ErikPelli/ItalyPassportAlert"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	location := os.Getenv("LOCATION")
	if location == "" {
		log.Fatal("Invalid location")
	}
	enableSound := strings.ToUpper(os.Getenv("ENABLE_SOUND")) == "TRUE"
	enableNotify := strings.ToUpper(os.Getenv("ENABLE_NOTIFY")) == "TRUE"

	req, err := ItalyPassportAlert.CreateRequest(location)
	if err != nil {
		log.Fatal("Unable to initialize request")
	}

	waitTime, err := strconv.Atoi(os.Getenv("WAIT_TIME"))
	if err != nil || waitTime < 0 {
		log.Fatal("Invalid wait time")
	}

	ItalyPassportAlert.RequestLoop(req, enableSound, enableNotify, waitTime)
}
