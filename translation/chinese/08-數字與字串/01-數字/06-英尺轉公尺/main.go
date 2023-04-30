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
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]

	// 英尺現在是一個 float64
	feet, _ := strconv.ParseFloat(arg, 64)

	meters := feet * 0.3048

	fmt.Printf("%f feet is %f meters.\n", feet, meters)

	// 印出來的數字太長了，所以我們可以使用 %g 來印出更簡潔的數字
	// fmt.Printf("%g feet is %g meters.\n", feet, meters)
}
