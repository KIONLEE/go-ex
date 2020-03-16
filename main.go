package main

import (
	"fmt"

	"github.com/KIONLEE/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("keon")
	fmt.Println(account)
}
