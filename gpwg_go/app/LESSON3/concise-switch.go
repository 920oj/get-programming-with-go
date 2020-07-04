package main

import "fmt"

func main() {
	fmt.Println("洞窟の入口だ。東へ進む道もある。")
	var command = "go inside"

	switch command {
	case "go east":
		fmt.Println("きみはさらに山を登る。")
	case "enter cave", "go inside":
		fmt.Println("きみは薄暗い洞窟の中にいる")
	case "read sign":
		fmt.Println("「未成年立入禁止」と書いてある")
	default:
		fmt.Println("よくわからん")
	}
}