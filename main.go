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

func canIDrink(age int) bool {
	// if koAge := age + 2; koAge < 18 {
	// 	return false
	// }
	// return true
	switch koAge := age + 2; koAge {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func main() {
	keon := map[string]int{"name": 1, "age": 2}
	for key, value := range keon {
		fmt.Println(key, ":", value)
	}
}
