package main

import "fmt"

func main() {
	var command = "go east"

	if command == "go east" {
		fmt.Println("きみはさらに山を登る。")
	} else if command == "go inside" {
		fmt.Println("きみは洞窟に入り、そこで一生を過ごす。")
	} else {
		fmt.Println("なんだかよくわからない")
	}
}