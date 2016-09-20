package examples

import (
	"fmt"

	"github.com/younggi/study_golang/discovery_go/chap3/hangul"
)

func Example_array() {
	fruits := [...]string{"사과", "바나나", "토마토", "수박"}
	for _, fruit := range fruits {
		if hangul.HasConsonantSuffix(fruit) {
		} else {
			fmt.Printf("%s는 맛있다.\n", fruit)
		}
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
	// 수박은 맛있다.
}
