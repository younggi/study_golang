package stack

import (
	"fmt"

	"github.com/younggi/study_golang/discovery_go/chap3/stack"
)

func ExampleEval() {
	fmt.Println(stack.Eval("5"))
	fmt.Println(stack.Eval("1 + 2"))
	fmt.Println(stack.Eval("1 + 2 + 3"))
	fmt.Println(stack.Eval("3 * ( 3 + 1 * 3 ) / 2"))
	fmt.Println(stack.Eval("3 * ( ( 3 + 1 ) * 3 ) / 2"))
	// Output:
	// 5
	// 3
	// 6
	// 9
	// 18
}
