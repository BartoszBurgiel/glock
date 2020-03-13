package main

import (
	"fmt"
	"glock"
	"time"
)

func main() {
	for {
		// clear the console
		fmt.Print("\033[H\033[2J")

		// create time
		t := glock.NewGlockTime()

		// print the display
		fmt.Println(glock.Display(t))

		// sleep
		time.Sleep(time.Second)
	}
}
