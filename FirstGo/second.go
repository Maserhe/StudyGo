package main

import "fmt"

func test() (int, int) {
	return 100, 200
}

func main() {
	a, b := test()
	_, c := test()
	d, _ := test()

	fmt.Println(a, b, c, d)

	// 局部变量 和去哪句变量 同名，局部变量 作用域更强

	// 定义 常量
	const URL string = "123"
	const URL2 = "www.baidu.com"
	fmt.Println(URL, URL2)

	// 定义多个常量
	const A, B, C = 1, "333", false
	fmt.Println(A, B, C)

	// iota
	const (
		aa = iota   // 0
		bb          // 1
		cc          // 2
		dd = "haha" // 3
		ee          // 4
		ff = 100    // 5
		gg          // 6
		hh = iota   // 7
		ii          //8
	)

	fmt.Println(aa, bb, cc, ee, gg, ii)
}
