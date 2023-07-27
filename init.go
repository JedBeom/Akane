package main

import (
	"time"

	"github.com/BurntSushi/toml"
	madon "github.com/McKael/madon/v3"
)

const CONFIG_FILE = "config.toml"
const REACTION_FILE = "reactions.toml"

var RESTART_MS = 2000

type Config struct {
	Instance     string
	AccessToken  string
	ClientKey    string
	ClientSecret string
	RestartMs    int
}

func loadConfig() (Config, error) {
	var config Config
	_, err := toml.DecodeFile(CONFIG_FILE, &config)
	RESTART_MS = config.RestartMs
	return config, err
}

type Reaction struct {
	Keywords []string
	Answers  []string
}

func loadReactions() ([]Reaction, error) {
	file := struct {
		Reactions []Reaction
	}{
		Reactions: nil,
	}

	_, err := toml.DecodeFile(REACTION_FILE, &file)
	return file.Reactions, err
}

func initApp(config Config) (*madon.Client, error) {
	userToken := madon.UserToken{
		AccessToken: config.AccessToken,
		CreatedAt:   time.Now().UnixNano(),
		Scope:       "read write",
		TokenType:   "urn:ietf:wg:oauth:2.0:oob",
	}

	mc, err := madon.RestoreApp("=AkaneBot=", config.Instance, config.ClientKey, config.ClientSecret, &userToken)
	return mc, err
}
