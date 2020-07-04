package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("薄暗い洞窟の中にいます")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println("洞窟を出ます: ", exit)
}