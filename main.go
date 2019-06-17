package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	go keepAlive()
	go Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	func() {
		for sig := range c {
			fmt.Println(sig)
			close(stopChan)
			os.Exit(1)
		}
	}()

}
