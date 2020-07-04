package main

import "fmt"

func main() {
	const distance = 56000000 // km
	const days = 28 // 日
	fmt.Println(distance / (days * 24) , "km/時")
}