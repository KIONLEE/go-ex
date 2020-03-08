package main

import (
	"fmt"
	"strings"
)

func multiply(a int, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func repeatMe(words ...string) {
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
