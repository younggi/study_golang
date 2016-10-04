package main

import (
	"context"
	"fmt"
	"runtime"
)

// Range returns a channel and sends ints
// (start, start+step, start+2*step, ...).
func Range(ctx context.Context, start, step int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; ; i += step {
			select {
			case out <- i:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// IntPipe returns input channel and output channel
type IntPipe func(context.Context, <-chan int) <-chan int

// FilterMultiple returns a IntPipe that filters multiple of n.
func FilterMultiple(n int) IntPipe {
	return func(ctx context.Context, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for x := range in {
				if x%n == 0 {
					continue
				}
				// n으로 나누어 떨어 지지 않을때 다음에 체인 되어 있는 FilterMultiple 함수로 넘겨진다.
				select {
				case out <- x:
				case <-ctx.Done():
					return
				}
			}
		}()
		return out
	}
}

// Primes returns channel output of primes
func Primes(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		c := Range(ctx, 2, 1)
		for {
			select {
			case i := <-c:
				c = FilterMultiple(i)(ctx, c)
				select {
				case out <- i:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// PrintPrimes prints prime numbers to max
func PrintPrimes(max int) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for prime := range Primes(ctx) {
		if prime > max {
			fmt.Println("\n Goroutines", runtime.NumGoroutine())
			break
		}
		fmt.Print(prime, " ")
	}
	fmt.Println()
}

func main() {
	PrintPrimes(10)
	PrintPrimes(100)
	PrintPrimes(1000)
}
