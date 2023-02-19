package main

import (
	"log"
	"wallet-checker/internal/app"
	"wallet-checker/internal/config"
)

func main() {
	log.Println("config initializing...")
	cfg := config.GetConfig()

	log.Println("app initializing...")
	newApp, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("start application...")
	if err := newApp.Run(); err != nil {
		log.Fatalln(err)
	}
	log.Println("done!")
}
