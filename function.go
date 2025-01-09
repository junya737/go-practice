// if

package main

import "fmt"

func main() {
	var x int = 10
	const y string = "hello"

	if x > 1 {
		fmt.Println(y)
	} else {
		fmt.Println("world")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		break
	}
	var z string
	z = switch_practice(33)
	fmt.Println(z)

	var arry = []int{1, 2, 3, 4, 5}
	fmt.Println(arry)
	arry[3] = 100
	fmt.Println(arry)
	arry = append(arry, 200)
	fmt.Println(arry)

	// map[type]typep
	m := make(map[string]int)
	m["key"] = 100
	fmt.Println(m)

	mp := map[string]int{
		"abc": 200,
	}
	fmt.Println(mp)
}

func switch_practice(x int) string {
	switch {
	case x == 1:
		return "1"
	case x == 2:
		return "2"
	case x > 3:
		return "greater than 3"
	default:
		return "default"
	}
}
