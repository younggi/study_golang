package prob1

import (
	"regexp"
	"strconv"
	"strings"
)

// BinOp is operator type
type BinOp func(int, int) int

// Eval returns the evaluation result of the given expr.
// The expression can have +, -, *, -, (, ) operators and
// decimal integers. Operators and operands should be
// space delimited.
func Eval(opMap map[string]BinOp, prec PrecMap, expr string) (int, error) {
	ops := []string{"("}
	var nums []int
	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}
	reduce := func(nextOp string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if _, higher := prec[nextOp][op]; nextOp != ")" && !higher {
				// 목록에 없는 연산자이므로 종료
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				// 괄호를 제거하였으므로 종료
				return
			}
			b, a := pop(), pop()
			if f := opMap[op]; f != nil {
				nums = append(nums, f(a, b))
			}
		}
	}
	for _, token := range strings.Split(expr, " ") {
		if len(token) == 0 {
			continue
		}
		if token == "(" {
			ops = append(ops, token)
		} else if _, ok := prec[token]; ok {
			reduce(token)
			ops = append(ops, token)
		} else if token == ")" {
			// 닫는 괄호는 여는 괄호까지 계산하고 제거
			reduce(token)
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, err
			}
			nums = append(nums, num)
		}
	}
	reduce(")")
	return nums[0], nil
}

// StrSet is String set type
type StrSet map[string]struct{}

// NewStrSet returns a new StrSet
func NewStrSet(strs ...string) StrSet {
	m := StrSet{}
	for _, str := range strs {
		m[str] = struct{}{}
	}
	return m
}

// PrecMap keyed by operator to set of higher precednece operators
type PrecMap map[string]StrSet

// NewEvaluator is Eval generator
func NewEvaluator(opMap map[string]BinOp, prec PrecMap) func(expr string) (int, error) {
	return func(expr string) (int, error) {
		return Eval(opMap, prec, expr)
	}
}

// NewEvaluatorEx is return error or result
func NewEvaluatorEx(opMap map[string]BinOp, prec PrecMap) func(expr string) string {
	return func(expr string) string {
		v, err := Eval(opMap, prec, expr)
		if err != nil {
			return err.Error()
		}
		return strconv.Itoa(v)
	}
}

var rx = regexp.MustCompile(`{[^}]+}`)

// EvalReplaceAll replaces calc
func EvalReplaceAll(in string) string {
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

	out := rx.ReplaceAllStringFunc(in, func(expr string) string {
		expr = strings.Trim(expr, "{ }")
		return eval(expr)
	})
	return out
}
