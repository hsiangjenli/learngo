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
	// 當您在計算中同時使用`浮點數`值和`整數`時
	// 最後結果的型態會變成`浮點數`

	ratio := 3.0 / 2

	// 其他範例:
	// ratio = 3 / 2.0

	// 其他範例:
	// 宣告一個名稱為 n 的變數, 並將其值設為 3
	// 其型別為`整數`, 再將其轉換成`浮點數`
	// n := 3
	// ratio = float64(n) / 2

	fmt.Printf("%f", ratio)
}
