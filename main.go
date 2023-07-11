package main

import "log"

func main() {
	config, err := loadConfig()
	if err != nil {
		panic(err)
	}

	reactions, err := loadReactions()
	if err != nil {
		panic(err)
	}

	log.Println("Loading files completed")

	mc, err := initApp(config)
	if err != nil {
		panic(err)
	}

	run(mc, reactions)
}
