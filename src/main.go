package main

import (
	"fmt"
	"./ui"
)

func main() {
	go func() {
		receiver := ui.MessageReceiver()
		for message := range receiver {
			fmt.Printf("Message from UI: %s\n", message)
		}
	}()

	ui.MainLoop()
}
