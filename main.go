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

func main() {
	// fmt.Println("hello world!")
	// ex.ExSay()
	// fmt.Println(multiply(2, 3))
	// totalLength, upperName := lenAndUpper("keon")
	// fmt.Println(totalLength, upperName)
	repeatMe("I", "am", "a", "rockstar", "!!!")
}
