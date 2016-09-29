package prob5

import "fmt"

func IteratorGenerator(a []string) func() string {
	var index int
	return func() string {
		var r string
		if index < len(a) {
			r = a[index]
			index++
			return r
		}
		return ""
	}
}

func ExampleInteratorGenerator() {
	ss := []string{"a", "b", "c", "d", "e"}
	iter := IteratorGenerator(ss)

	for {
		s := iter()
		if s == "" {
			break
		}
		fmt.Println(s)
	}
	// Output:
	// a
	// b
	// c
	// d
	// e
}
