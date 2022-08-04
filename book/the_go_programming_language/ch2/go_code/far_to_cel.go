package main

import (
	"fmt"
)

const ff, bpf = 32.0, 212.0

func main() {
	fmt.Printf("The freezing point for °C is: %g °C\n", conversion(ff))
	fmt.Printf("The boiling point for °C is: %g °C\n", conversion(bpf))
}

func conversion(fahrenheit float64) (conv float64) {
	conv = (fahrenheit - 32) / 1.8
	return conv
}
