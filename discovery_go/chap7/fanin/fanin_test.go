package fanin

import (
	"fmt"
	"sync"
)

func FanIn(ins ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {
		go func(in <-chan int) {
			defer wg.Done()
			for num := range in {
				out <- num
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func ExampleFanIn() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	go func() {
		defer close(c1)
		defer close(c2)
		defer close(c3)
		for i := 0; i < 10; i++ {
			c1 <- i
			c2 <- i
			c3 <- i
		}
	}()
	c := FanIn(c1, c2, c3)
	for in := range c {
		fmt.Println(in)
	}
}
