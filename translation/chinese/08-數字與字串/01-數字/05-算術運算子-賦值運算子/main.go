// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
)

func main() {
	width, height := 5., 12.

	// 計算矩形的面積
	area := width * height
	fmt.Printf("%gx%g=%g\n", width, height, area)

	area = area - 10 // 面積減少 10
	area = area + 10 // 面積增加 10
	area = area * 2  // 面積乘以  2
	area = area / 2  // 面積除以  2
	fmt.Printf("area=%g\n", area)

	// // 賦值運算子
	area -= 10 // 面積減少 10
	area += 10 // 面積增加 10
	area *= 2  // 面積乘以  2
	area /= 2  // 面積除以  2
	fmt.Printf("area=%g\n", area)

	// 計算面積的餘數
	// 注意: 如果 area 是浮點數, 就不能這樣寫:
	// area %= 7

	// 可以這樣寫
	area = float64(int(area) % 7)
	fmt.Printf("area=%g\n", area)
}
