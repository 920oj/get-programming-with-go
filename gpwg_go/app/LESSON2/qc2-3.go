package main

import "fmt"

func main() {
	const ( 
		speed = 100800
		distance = 96300000
	)
	fmt.Println(distance / speed, "秒")
}