package main

import "fmt"

// Move function
func Move(n, from, to, via int) {
	if n <= 0 {
		return
	}
	Move(n-1, from, via, to)
	fmt.Println(from, "->", to)
	Move(n-1, via, to, from)
}

// Hanoi print seq of change
func Hanoi(n int) {
	fmt.Println("Number of disks:", n)
	Move(n, 1, 2, 3)
}

func main() {
	Hanoi(3)
}
