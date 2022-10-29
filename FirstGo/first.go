package main

import "fmt"

func main() {

	// 定义 变量
	var name string = "123"
	var num int = 123

	fmt.Println("helloWorld")
	fmt.Println(name, num)

	// 定义多个变量
	var (
		a string
		b int
		c bool
		d float32
		e float64
	)
	a = "123"
	d = 0.0
	e = 1.000

	fmt.Println(a, b, c, d, e)

	// 自动推导类型
	one := "abc"
	two := 123
	fmt.Println(one, two)

	// 格式化输出
	fmt.Printf("num: %d, 内存地址%p", b, &d)
	fmt.Println()
	fmt.Printf("num: %f, 内存地址%p", d, &d)
	fmt.Println()

	// 数据交换
	var n1, n2 int
	n1 = 100
	n2 = 200

	n1, n2 = n2, n1

	fmt.Println(n1, n2)
}
