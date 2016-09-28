package sorting

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
	"testing"
)

type CaseInsensitive []string

func (c CaseInsensitive) Len() int {
	return len(c)
}

func (c CaseInsensitive) Less(i, j int) bool {
	return strings.ToLower(c[i]) < strings.ToLower(c[j]) ||
		(strings.ToLower(c[i]) == strings.ToLower(c[j]) && c[i] < c[j])
}

func (c CaseInsensitive) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func ExampleCaseInsensitive_sort() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MackBook", "AppStore",
	})
	sort.Sort(apple)
	fmt.Println(apple)
	// Output:
	// [AppStore iPad iPhone MackBook]
}

func TestCaseIntensiveSort(t *testing.T) {
	cases := []struct {
		in, want CaseInsensitive
	}{
		{
			CaseInsensitive{"iPhone", "iPad", "MacBook", "AppStore"},
			CaseInsensitive{"AppStore", "iPad", "iPhone", "MacBook"},
		},
		{
			CaseInsensitive{"1", "7", "3", "5"},
			CaseInsensitive{"1", "3", "5", "7"},
		},
	}
	for _, c := range cases {
		sort.Sort(c.in)
		for i := 0; i < len(c.in); i++ {
			if c.in[i] != c.want[i] {
				t.Errorf("[%d] %s == %s, want %s", i, c.in[i], c.want[i], c.want[i])
			}
		}
	}
}

func (c *CaseInsensitive) Push(x interface{}) {
	*c = append(*c, x.(string))
}

func (c *CaseInsensitive) Pop() interface{} {
	len := c.Len()
	last := (*c)[len-1]
	*c = (*c)[:len-1]
	return last
}

func ExampleCaseInsensitive_heap() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MackBook", "AppStore",
	})
	heap.Init(&apple)
	for apple.Len() > 0 {
		fmt.Println(heap.Pop(&apple))
	}
	// Output:
	// AppStore
	// iPad
	// iPhone
	// MackBook
}

func ExampleCaseInsensitive_heapString() {
	apple := CaseInsensitive([]string{
		"iPhone", "iPad", "MacBook", "AppStore",
	})
	heap.Init(&apple)
	for apple.Len() > 0 {
		popped := heap.Pop(&apple)
		s := popped.(string)
		fmt.Println(s)
	}
	// Output:
	// AppStore
	// iPad
	// iPhone
	// MacBook
}
