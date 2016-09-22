package prob3

import "fmt"

// 이진 검색 알고리즘
func Example_BinarySearch() {
	s := []string{"가", "나", "다", "라", "마", "바", "아", "자", "차", "카"}
	fmt.Printf("%t", BinarySearch(s, "라"))
	// Output:
	// true
}

func BinarySearch(list []string, t string) bool {
	// 중간값을 구한다.
	index := len(list) / 2
	// 비교
	// 작은 쪽 or 큰 쪽으로 슬라이스를 나누어 다시 호출
	if list[index] > t {
		return BinarySearch(list[:index], t)
	} else if list[index] < t {
		return BinarySearch(list[index:], t)
	} else if list[index] == t {
		return true
	}
	return false
}
