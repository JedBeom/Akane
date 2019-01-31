package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/McKael/madon"
)

type Config struct {
	Instance     string `json:"instance"`
	ClientKey    string `json:"client_key"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
}

var (
	config Config
	mc     *madon.Client
)

func init() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	userToken := madon.UserToken{
		AccessToken: config.AccessToken,
		CreatedAt:   time.Now().UnixNano(),
		Scope:       "read write",
		TokenType:   "urn:ietf:wg:oauth:2.0:oob",
	}

	mc, err = madon.RestoreApp("MillionTC", config.Instance, config.ClientKey, config.ClientSecret, &userToken)

	if err != nil {
		panic(err)
	}
}

func toot(content, cw string) (st *madon.Status, err error) {
	status := madon.PostStatusParams{
		Text:       content,
		Visibility: "unlisted",
	}

	status.SpoilerText = cw

	st, err = mc.PostStatus(status)
	return
}

func reply(noti *madon.Notification, content string) (st *madon.Status, err error) {

	current, err := mc.GetCurrentAccount()
	if err != nil {
		fmt.Println("Err", err)
		return
	}

	s := noti.Status

	var mentions []string
	// Add the sender if she is not the connected user
	if s.Account.Acct != current.Acct {
		mentions = append(mentions, "@"+s.Account.Acct)
	}
	for _, m := range s.Mentions {
		if m.Acct != current.Acct && m.Acct != s.Account.Acct {
			mentions = append(mentions, "@"+m.Acct)
		}
	}
	mentionsStr := strings.Join(mentions, " ")
	content = mentionsStr + " " + content

	var visibility string
	if s.Visibility != "public" {
		visibility = s.Visibility
	} else {
		visibility = "unlisted"
	}

	status := madon.PostStatusParams{
		Text:       content,
		InReplyTo:  noti.Status.ID,
		Visibility: visibility,
	}

	st, err = mc.PostStatus(status)

	return
}
