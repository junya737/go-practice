// hello world

package main

import "fmt"

func main() {
	fmt.Println("hello world")

	// 変数の宣言
	var x int = 1
	var y int = 2
	fmt.Println(add(x, y))

}

// 足し算をする関数
func add(x int, y int) int {
	return x + y
}
