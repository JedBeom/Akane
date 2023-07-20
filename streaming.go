package main

import (
	"log"
	"time"

	madon "github.com/McKael/madon/v3"
)

func openStream(mc *madon.Client, pEvents *chan madon.StreamEvent, pStopChan *chan bool, pDoneChan *chan bool) error {
	// open again
	*pEvents = make(chan madon.StreamEvent)
	*pStopChan = make(chan bool)
	*pDoneChan = make(chan bool)

	return mc.StreamListener("user", "", *pEvents, *pStopChan, *pDoneChan)
}

func eventHandler(mc *madon.Client, reactions []Reaction, event madon.StreamEvent) {
	if event.Event != "notification" {
		return
	}

	noti, ok := event.Data.(madon.Notification)
	if !ok {
		return
	}

	if noti.Account.Bot {
		return
	}

	if noti.Type != "mention" {
		return
	}

	err := reply(mc, noti.Status, reactions)
	if err != nil {
		log.Println("reply():", err, "|", "noti.status:", noti.Status)
	}
}

func run(mc *madon.Client, reactions []Reaction) {
	var (
		events   chan madon.StreamEvent
		stopChan chan bool
		doneChan chan bool
	)

	openStream(mc, &events, &stopChan, &doneChan)

	for {
		select {
		case event := <-events:
			if event.Error != nil {
				log.Println("event.Error", event.Error)
				continue
			}

			go eventHandler(mc, reactions, event)
		case <-doneChan: // if close(doneChan) was executed

			for {
				time.Sleep(time.Millisecond * 500)

				err := openStream(mc, &events, &stopChan, &doneChan)
				if err != nil {
					close(stopChan)
					log.Println("openStream():", err)
					continue
				}
				log.Println("Restart Success!")
				break
			}

		}
	}
}
