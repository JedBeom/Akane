package main

import (
	"fmt"
	"log"
	"time"

	"github.com/JedBeom/madon/v3"
)

var (
	events   = make(chan madon.StreamEvent)
	stopChan = make(chan bool)
	doneChan = make(chan bool)
)

func Run() {

	err := mc.StreamListener("user", "", events, stopChan, doneChan)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		event := <-events

		if event.Error != nil {
			log.Println("Event error:", err)
		}

		// Catch notifications
		if event.Event == "notification" {
			noti := event.Data.(madon.Notification)

			// Avoid bot loop
			if noti.Account.Bot == true {
				continue
			}

			// React only mention
			if noti.Type == "mention" {
				content := contentExtraction(noti.Status.Content)
				go foolAkane(&noti, content)
			}
		}

	}

}

// Restart streamer if session closed
func keepAlive() {

	for {
		fmt.Println("KeepAlive Loop started")
		_, ok := <-doneChan
		if !ok {
			fmt.Println("Restarting...")

			stopChan = make(chan bool)
			doneChan = make(chan bool)

			err := mc.StreamListener("user", "", events, stopChan, doneChan)
			if err != nil {
				panic(err)
			}
		}

		time.Sleep(time.Second)

	}
}
