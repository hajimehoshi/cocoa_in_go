package main

import (
	"fmt"
	"./ui"
)

func main() {
	go func() {
		receiver := ui.MessageReceiver()
		sender := ui.MessageSender()
		for message := range receiver {
			fmt.Printf("Message from UI: %s\n", message)
			sender<- "OK, received: " + message
		}
	}()

	ui.MainLoop()
}
