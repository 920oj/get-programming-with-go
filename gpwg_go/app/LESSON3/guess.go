package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	const num = 30

	rand.Seed(time.Now().UnixNano())
	var ans = rand.Intn(100)

	fmt.Printf("コンピュータが選んだ数は%vです。\n", ans)
	if ans == num {
		fmt.Println("大正解！\n")
	} else if ans > num {
		fmt.Println("残念！ちょっと大きいかも\n")
	} else if ans < num {
		fmt.Println("残念！ちょっと小さいかも\n")
	} else {
		fmt.Println("不明なエラーです。\n")
	}
}