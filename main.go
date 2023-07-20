package main

import (
	"log"
	"os"
)

func main() {
	const LOG_FILE string = "akane.log"
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panicln(err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	config, err := loadConfig()
	if err != nil {
		log.Panicln(err)
	}

	reactions, err := loadReactions()
	if err != nil {
		log.Panicln(err)
	}

	mc, err := initApp(config)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Bot Run starts")
	run(mc, reactions)
}
