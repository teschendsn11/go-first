package main

import (
	"log"

	"teschendsn11/go-first/cli"
	"teschendsn11/go-first/h"

	"github.com/joho/godotenv"
)

func main() {
	h := h.LogInit(h.LoggerOptions{
		Channel: nil,
	})

	h.Debug("app start")
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cli.StartCLI()
}
