// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	// 當`整數`與`浮點數`一起做運算時
	// 最後的型態會是`浮點數`
	fmt.Println(8 * -4.0) // -32.0 not -32

	// 兩個`整數`最後的型態會是`整數`
	fmt.Println(-4 / 2)

	// 餘數運算子
	// 只能使用在`整數`中
	fmt.Println(5 % 2)
	// fmt.Println(5.0 % 2) // 錯誤: 不能使用在浮點數中

	// 加法運算子
	fmt.Println(1 + 2.5)
	fmt.Println(2 - 3)

	// 減法運算子
	fmt.Println(-(-2))
	fmt.Println(- -2) // 也可以這樣寫
}
