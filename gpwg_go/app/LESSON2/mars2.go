package main

import "fmt"

func main() {
	fmt.Printf("火星の表面で、私の体重は、%vキログラム。", 62.0 * 0.3783)
	fmt.Printf("年齢は、%v歳になるでしょう。\n", 20 * 365 / 687)
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}