package main

import (
	"fmt"
)

func main() {
	summ := 0
	for i := 1; i < 100000000; i++ {
		summ += i // sum = sum + i
	}
	//fmt.Print("Sum : ", summ)
	fmt.Print(summ)
}
