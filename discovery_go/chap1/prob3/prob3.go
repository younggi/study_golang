package main

import "fmt"

// Pibo is pibonacci function
func Pibo(n int) {
	p, q := 0, 1
	for n >= 0 {
		fmt.Println(p)
		p, q = q, p+q
		n--
	}
}

func main() {
	Pibo(10)
}
