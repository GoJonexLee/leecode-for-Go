package main

import "fmt"

func findComplement(num int) int {
	tp := ^0
	for num & tp != 0 {
		tp <<= 1
	}
	return ^num & ^tp
}

func main() {
	fmt.Println(findComplement(5))
}