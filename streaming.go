package main

import (
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
				close(stopChan) // causes close(doneChan)
				continue
			}

			if event.Event != "notification" {
				continue
			}

			noti, ok := event.Data.(madon.Notification)
			if !ok {
				continue
			}

			if noti.Account.Bot {
				continue
			}

			if noti.Type != "mention" {
				continue
			}

			go reply(mc, noti.Status, reactions)

		case <-doneChan: // if close(doneChan) was executed
			time.Sleep(time.Millisecond * 500)
			openStream(mc, &events, &stopChan, &doneChan)
		}
	}
}
