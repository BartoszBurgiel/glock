package main

import (
	"fmt"
	"time"
	"watch"
)

func main() {
	fmt.Println("hello")

	for i := 0; i <= 12; i++ {
		cc := watch.Clock(i, 12)
		fmt.Println(cc)

	}

	for {
		time.Sleep(time.Second)
		fmt.Println("\033[H\033[2J")
		fmt.Println(watch.ClocksWithCurrentTime())
	}
}
