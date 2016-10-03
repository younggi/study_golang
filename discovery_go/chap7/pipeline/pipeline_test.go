package pipeline

import "fmt"

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

func ExamplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
	}()
	for num := range PlusOne(PlusOne(c)) {
		fmt.Println(num)
	}
	// Output:
	// 7
	// 5
	// 10
	PlusTwo := Chain(PlusOne, PlusOne)
	for num := range PlusTwo(c) {
		fmt.Println(num)
	}
	// Output:
	// 7
	// 5
	// 10
}

type IntPipe func(<-chan int) <-chan int

func Chain(ps ...IntPipe) IntPipe {
	return func(in <-chan int) <-chan int {
		c := in
		for _, p := range ps {
			c = p(c)
		}
		return c
	}
}
