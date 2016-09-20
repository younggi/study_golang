package main

import "fmt"

func main() {
	for i, r := range "가나다" {
		fmt.Println(i, r)
	}
	fmt.Println(len("가나다"))

	for _, r := range "가갛힣" {
		fmt.Println(string(r), r)
	}
}
