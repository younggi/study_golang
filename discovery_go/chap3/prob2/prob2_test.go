package prob2

import "fmt"

func Example_sorting() {
	list := []int{5, 4, 3, 2, 1}
	Sorting(list)
	fmt.Println(list)
	// Output:
	// [1 2 3 4 5]
}

func Sorting(list []int) {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
