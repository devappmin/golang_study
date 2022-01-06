package main

import (
	"fmt"

	"main.com/accounts"
)

func main() {
	account := accounts.NewAccount("Kim Seung Hwan")
	account.Deposit(30)

	fmt.Println(account)
}
