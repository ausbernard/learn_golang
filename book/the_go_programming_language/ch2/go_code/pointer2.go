package main

import "fmt"

func incr(p *int) int {
	*p++
	fmt.Println(*p)
	return *p
}

func main() {
	v := 1
	fmt.Println("&v =", &v)
	incr(&v)
	fmt.Println("the final increment is:", incr(&v))
}
