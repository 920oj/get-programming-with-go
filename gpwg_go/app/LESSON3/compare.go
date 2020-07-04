package main

import "fmt"

func main() {
	fmt.Println("入口の近くに「未成年立入禁止」という立て札がある")

	var age = 21
	var minor = age < 18
	fmt.Printf("%v歳のぼくは未成年ですか？ %v\n", age, minor)
}