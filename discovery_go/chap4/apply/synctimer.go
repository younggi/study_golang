package main

import (
	"fmt"
	"time"
)

// CountDown counts numbers
func CountDown(seconds int) {
	for seconds > 0 {
		fmt.Println(seconds)
		time.Sleep(time.Second)
		seconds--
	}
}

func main() {
	fmt.Println("Ladies and gentlemen!")
	CountDown(5)
}
