package main

import (
	"fmt"
	"./ui"
)

func main() {
	go func() {
		fmt.Println("Output from Go!")
		fmt.Println("Can you see this message?")
	}()

	ui.MainLoop();
}
