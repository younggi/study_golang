package prob1

import (
	"fmt"
	"strings"
)

func ExampleNewEvaluator() {
	eval := NewEvaluatorEx(map[string]BinOp{
		"**": func(a, b int) int {
			if a == 1 {
				return 1
			}
			if b < 0 {
				return 0
			}
			r := 1
			for i := 0; i < b; i++ {
				r *= a
			}
			return r
		},
		"*":   func(a, b int) int { return a * b },
		"/":   func(a, b int) int { return a / b },
		"mod": func(a, b int) int { return a % b },
		"+":   func(a, b int) int { return a + b },
		"-":   func(a, b int) int { return a - b },
	}, PrecMap{
		"**":  NewStrSet(),
		"*":   NewStrSet("**", "*", "/", "mod"),
		"/":   NewStrSet("**", "*", "/", "mod"),
		"mod": NewStrSet("**", "*", "/", "mod"),
		"+":   NewStrSet("**", "*", "/", "mod", "+", "-"),
		"-":   NewStrSet("**", "*", "/", "mod", "+", "-"),
	})

	fmt.Println(eval("5"))
	fmt.Println(eval("1 + 2"))
	fmt.Println(eval("1  + 2"))
	fmt.Println(eval("1 - 2 - 4"))
	fmt.Println(eval("( 3 - 2 ** 3 ) * ( -2 )"))
	fmt.Println(eval("3 * ( 3 + 1 * 3 ) / ( -2 )"))
	fmt.Println(eval("3 * ( ( 3 + 1 ) * 3 ) / 2"))
	fmt.Println(eval("1 + 2 ** 10 * 2"))
	fmt.Println(eval("2 ** 3 mod 3"))
	fmt.Println(eval("2 ** 2 ** 3"))
	fmt.Println(eval("1 = 2"))
	// Output:
	// 5
	// 3
	// 3
	// -5
	// 10
	// -9
	// 18
	// 2049
	// 2
	// 256
	// strconv.ParseInt: parsing "=": invalid syntax
}

func ExampleEvalReplaceAll() {
	in := strings.Join([]string{
		"다들 그 동안 고생이 많았다.",
		"첫째는 분당에 있는 { 2 ** 4 * 3 }평 아파트를 갖거라.",
		"둘째는 임야 { 10 ** 5 mod 7777 }평을 가져라.",
		"막내는 { 10000 - ( 10 ** 5 mod 7777 ) }평 임야를 갖고",
		"배기량 { 711 * 8 / 9 }cc의 경운기를 갖거라.",
	}, "\n")
	fmt.Println(EvalReplaceAll(in))
	// Output:
	// 다들 그 동안 고생이 많았다.
	// 첫째는 분당에 있는 48평 아파트를 갖거라.
	// 둘째는 임야 6676평을 가져라.
	// 막내는 3324평 임야를 갖고
	// 배기량 632cc의 경운기를 갖거라.
}
