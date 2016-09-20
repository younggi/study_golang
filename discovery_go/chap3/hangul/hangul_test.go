package hangul

import (
	"fmt"
	"strconv"
)

func ExampleHasConsonantSuffix() {
	fmt.Println(HasConsonantSuffix("Go 언어"))
	fmt.Println(HasConsonantSuffix("그럼"))
	fmt.Println(HasConsonantSuffix("우리 밥 먹고 합시다."))

	// Output:
	// false
	// true
	// false
}

func Example_printByte() {
	s := "가나다"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x:", s[i])
	}
	fmt.Println()
	// Output:
	// ea:b0:80:eb:82:98:eb:8b:a4:
}

func Example_printByte2() {
	s := "가나다"
	fmt.Printf("%x\n", s)
	fmt.Printf("% x\n", s)
	// Output:
	// eab080eb8298eb8ba4
	// ea b0 80 eb 82 98 eb 8b a4
}

func Example_modifybytes() {
	b := []byte("가나다")
	b[2]++
	fmt.Println(string(b))
	// Output:
	// 각나다
}

func Example_strCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
	// Output:
	// abcdef
	// abcdef
}

func Example_printNum() {
	fmt.Println(int('5'))
	var i int
	var k int64
	var f float64
	var s string
	var err error
	i, err = strconv.Atoi("350")
	fmt.Println(i)
	k, err = strconv.ParseInt("cc7fdd", 16, 32)
	fmt.Println(k)
	k, err = strconv.ParseInt("0xcc7fdd", 0, 32)
	fmt.Println(k)
	f, err = strconv.ParseFloat("3.14", 64)
	fmt.Println(f)
	fmt.Println(err)
	s = strconv.Itoa(340)
	fmt.Println(s)
	s = strconv.FormatInt(13402077, 16)
	fmt.Println(s)

	var num int
	fmt.Sscanf("57", "%d", &num)
	s = fmt.Sprint(3.14)
	fmt.Println(s)
	s = fmt.Sprintf("%x", 13402077)
	fmt.Println(s)

	// Output:
	// 53
	// 350
	// 13402077
	// 13402077
	// 3.14
	// <nil>
	// 340
	// cc7fdd
	// 3.14
	// cc7fdd
}
