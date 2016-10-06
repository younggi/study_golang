package qsort

import "fmt"

func ExampleQSort() {
	a := []int{7, 3, 5, 6, 4, 1, 2}
	a = QSort(a)
	fmt.Println(a)
	// Output:
	// [1 2 3 4 5 6 7]
}

func ExampleEnhancedQSort() {
	a := []int{7, 3, 5, 6, 4, 1, 2}
	EnhancedQSort(a, 0, len(a)-1)
	fmt.Println(a)
	// Output:
	// [1 2 3 4 5 6 7]
}
