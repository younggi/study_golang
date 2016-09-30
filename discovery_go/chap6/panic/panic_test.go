package panic

import "fmt"

func f() (i int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			i = -1
		}
	}()
	g() // This function panics.
	return 100
}

func g() {
	panic("I panic!")
}

func Example_f() {
	fmt.Println("f() =", f())
	// Output:
	// Recovered in f I panic!
	// f() = -1
}
