package main

import (
	"fmt"
)

const BoilingF = 212.0
const FreezingF = 32.0

func main() {
	var boiling = BoilingF
	var c = (boiling - 32) / 1.8
	fmt.Printf("The temperature of boiling in celcius is %g°C\n", c)
}
