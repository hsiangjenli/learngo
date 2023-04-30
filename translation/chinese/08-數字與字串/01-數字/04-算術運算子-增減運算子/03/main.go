// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// incdec 代表的是`增減運算子`
// 遞增運算子 inc = increase
// 遞減運算子 dec = decrease

func main() {
	var counter int

	// 正確寫法:

	counter++ // 1
	counter++ // 2
	counter++ // 3
	counter-- // 2
	
	fmt.Printf(
		"There are %d line(s) in the file\n",
		counter
	)

	// 錯誤寫法:

	// counter = 5+counter--
	// counter = ++counter + counter--
}
