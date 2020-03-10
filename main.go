package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (length int, upperCase string) {
	length = len(name)
	upperCase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	defer fmt.Println("repeatMe is done")
	fmt.Println(words)
}

func superAdd(numbers ...int) int {
	var total int = 0 // or simply, total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	// fmt.Println("hello world!")
	// ex.ExSay()
	// fmt.Println(multiply(2, 3))
	// totalLength, upperName := lenAndUpper("keon")
	// fmt.Println(totalLength, upperName)
	// repeatMe("I", "am", "a", "rockstar", "!!!")
	result := superAdd(1, 2, 3, 4, 5)
	fmt.Println(result)
}
