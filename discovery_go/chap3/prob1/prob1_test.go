package prob1

import "fmt"

// 함수명
// 받침 판별식 가rune code 값 % 28 == 0

func Example_HasConsonentSuffix() {
	items := []string{"가나", "각힣"}
	for _, item := range items {
		fmt.Printf("%s 는 %t 이다\n", item, HasConsonentSuffix(item))
	}
	// Output:
	// 가나 는 false 이다
	// 각힣 는 true 이다
}

// 전역 변수 상수 처럼 사용함
var (
	start = rune(44032)
	end   = rune(55204)
)

func HasConsonentSuffix(s string) bool {
	result := false
	numEnds := 28
	for _, r := range s {
		// input value validation은 반드시 한다.
		if start <= r && r < end {
			index := int(r - start)
			result = index%numEnds != 0
		}
	}
	return result
}

func Example_array() {
	fruits := []string{"사과", "바나나", "토마토", "수박"}
	for _, fruit := range fruits {
		if HasConsonentSuffix(fruit) {
			fmt.Printf("%s은 맛있다.\n", fruit)
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
