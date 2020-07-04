package main

import "fmt"

func main() {
	fmt.Println("2100年はうるう年ですか？")
	var year = 2100
	var leap = year % 400 == 0 || (year % 4 == 0 && year % 100 != 0)
	if leap {
		fmt.Println("はい。うるう年です")
	} else {
		fmt.Println("いいえ。ふつうの年です。")
	}
}