package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
)

func main() {
	a := new(accounts.Account)
	fmt.Println(a)
}
