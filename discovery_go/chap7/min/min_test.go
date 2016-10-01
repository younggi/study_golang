package min

import (
	"fmt"
	"sync"
)

func Min(a []int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, e := range a[1:] {
		if min > e {
			min = e
		}
	}
	return min
}

func ExampleMin() {
	fmt.Println(Min([]int{
		83, 46, 49, 23, 92,
		48, 39, 91, 44, 99,
		25, 42, 74, 56, 23,
	}))
	// Output: 23
}

func ParallelMin(a []int, n int) int {
	if len(a) < n {
		return Min(a)
	}
	mins := make([]int, n)
	// size가 마지막 부분까지 나눌 수 있도록 n - 1을 더하여 나눔
	size := (len(a) + n - 1) / n
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			begin, end := i*size, (i+1)*size
			if end > len(a) {
				end = len(a)
			}
			mins[i] = Min(a[begin:end])
		}(i)
	}
	wg.Wait()
	return Min(mins)
}

func ExampleParallelMin() {
	fmt.Println(ParallelMin([]int{
		83, 46, 49, 23, 92,
		48, 39, 91, 44, 99,
		25, 42, 74, 56, 23,
	}, 4))
	// Output: 23
}
