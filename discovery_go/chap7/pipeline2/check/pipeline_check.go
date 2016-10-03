package main

import (
	"fmt"
	"runtime"
	"time"
)

// PlusOne returns a channel of num + 1 for nums received from in
func PlusOne(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num + 1
		}
	}()
	return out
}

func main() {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 3; i < 103; i += 10 {
			c <- i
		}
	}()
	nums := PlusOne(PlusOne(PlusOne(PlusOne(PlusOne(c)))))
	for num := range nums {
		fmt.Println(num)
		if num == 18 {
			break
		}
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
	for _ = range nums {
		// Consume all nums
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println("NumGoroutine: ", runtime.NumGoroutine())
}
