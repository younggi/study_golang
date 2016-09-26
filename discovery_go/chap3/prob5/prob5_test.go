package prob5

import "fmt"

// 새로운 MultiSet을 생성하여 반환한다.
func NewMultiSet() map[string]int {
	return map[string]int{}
}

// Insert 함수는 집합에 val을 추가한다.
func Insert(m map[string]int, val string) {
	if v, exists := m[val]; exists {
		m[val] = v + 1
	} else {
		m[val] = 1
	}
}

// Erase 함수는 집합에서 val을 제거한다.집합에 val이 없는
// 경우에는 아무 일도 일어나지 않는다.
func Erase(m map[string]int, val string) {
	if v, exists := m[val]; exists {
		m[val] = v - 1
	}
}

// Count 함수는 집합에 val이 들어있는 횟수를 구한다.
func Count(m map[string]int, val string) int {
	return m[val]
}

// String 함수는 집합에 들어 있는 원소들을 { } 안에 빈칸으로
// 구분하여 넣은 문자열을 반환한다.
func String(m map[string]int) string {
	r := "{ "
	for key := range m {
		for i := 0; i < m[key]; i++ {
			r += key + " "
		}
	}
	r += "}"
	return r
}

func ExampleMultiSet() {
	m := NewMultiSet()
	fmt.Println(String(m))
	fmt.Println(Count(m, "3"))
	Insert(m, "3")
	Insert(m, "3")
	Insert(m, "3")
	Insert(m, "3")
	fmt.Println(String(m))
	fmt.Println(Count(m, "3"))
	Insert(m, "1")
	Insert(m, "2")
	Insert(m, "5")
	Insert(m, "7")
	Erase(m, "3")
	Erase(m, "5")
	fmt.Println(Count(m, "3"))
	fmt.Println(Count(m, "1"))
	fmt.Println(Count(m, "2"))
	fmt.Println(Count(m, "5"))
	fmt.Println(Count(m, "10"))
	// Output:
	// { }
	// 0
	// { 3 3 3 3 }
	// 4
	// 3
	// 1
	// 1
	// 0
	// 0
}
