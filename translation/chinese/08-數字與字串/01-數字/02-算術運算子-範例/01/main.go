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
	fmt.Println("sum :", 3+2)   // 加法 - 整數
	fmt.Println("sum :", 2+3.5) // 加法 - 浮點數
	fmt.Println("dif :", 3-1)   // 減法 - 整數
	fmt.Println("dif :", 3-0.5) // 減法 - 浮點數
	fmt.Println("prod:", 4*5)   // 乘法 - 整數
	fmt.Println("prod:", 5*2.5) // 乘法 - 浮點數
	fmt.Println("quot:", 8/2)   // 除法 - 整數
	fmt.Println("quot:", 8/1.5) // 除法 - 浮點數

	// 餘數僅適用於整數
	fmt.Println("rem :", 8%3)
	// fmt.Println("rem:", 8.0%3) // 錯誤寫法

	// 你可以這樣做
	// 因為浮點數的小數部分為零
	f := 8.0
	fmt.Println("rem :", int(f)%3)
}
