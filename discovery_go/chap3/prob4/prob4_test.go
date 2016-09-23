package prob4

import "fmt"

func Example_queue() {
	q := []string{"c", "d", "e"}

	printSliceInfo(q)

	// push queue
	q = append(q, "a")

	printSliceInfo(q)
	// pop queue
	e := q[0]
	q = q[1:]

	printSliceInfo(q)

	fmt.Println(e)

	// Output:
}

func printSliceInfo(q []string) {
	fmt.Println("======")
	fmt.Println(q)
	fmt.Println(len(q))
	fmt.Println(cap(q))
}
