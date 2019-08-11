package main

import "fmt"

type A []int

func (as A) teste() {
	as[1] = 7
	fmt.Println(as)
}

func main() {
	a := A{1, 2, 3}
	fmt.Println(a)
	a.teste()
	fmt.Println(a)
}
