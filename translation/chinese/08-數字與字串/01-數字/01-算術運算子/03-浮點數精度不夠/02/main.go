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
	// Go 的編譯器會將這些數字視為`整數`
	// 因為`整數`值中沒有小數部分
	// 所以結果變成了1, 而不是1.5

	// 因此, 這裡的ratio變量是一個`整數`變量
	// 這是因為3除以2的結果是一個`整數`, 而不是一個`浮點數`

	ratio := 3 / 2

	fmt.Printf("%d", ratio)
}
