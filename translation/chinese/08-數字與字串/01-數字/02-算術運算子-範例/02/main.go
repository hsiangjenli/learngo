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
	// 最後的值是多少？
	// 3 / 2 = 1.5?
	var ratio float64 = 3 / 2
	fmt.Println(ratio)

	// 解釋:
	// 上面的算式跟下方的算式相同
	ratio = float64(int(3) / int(2))
	fmt.Println(ratio)

	// 如何解決:
	// 當其中任一數值的型別為`浮點數`時
	// 最終的結果的型別也將是`浮點數`
	ratio = float64(3) / 2
	fmt.Println(ratio)

	// 範例:
	ratio = 3.0 / 2
	fmt.Println(ratio)
}
