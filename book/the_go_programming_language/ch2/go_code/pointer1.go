package main

import "fmt"

func f() *int {
	v := 1
	return &v
}

func main() {
	var p = f()
	fmt.Println("p points to the address of v:", p)
	fmt.Println("f()1:", f(), "| f()2:", f(), "| f()3:", f())
	fmt.Println("Therefore, the returned address of v will be different each time it is called (f() != f()):", f() != f())
}
