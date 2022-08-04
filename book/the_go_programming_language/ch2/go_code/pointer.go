package main

import (
	"fmt"
)

func main() {
	x := 1
	p := &x
	fmt.Println(x)
	fmt.Println("location of x:", p)
	fmt.Println("int assigned to the location of x:", *p)
	*p = 2
	fmt.Println("saying basically x is now =", *p)
	fmt.Println("x =", x)
	fmt.Println("address of x is equal to the address of x:", &x == &x, " | the expression of x is equal to the the variable to which p points to:", x == *p, "| var expression of p points to the address of x", p == &x, " |the address of x is equal to nil", &x == nil)
}
