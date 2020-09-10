package main

import "fmt"

func main() {
	fmt.Println("Эта программа меняет местами две переменных")
	a := 42
	b := 153

	fmt.Println("До смены мест:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	c := b
	b = a
	a = c

	fmt.Println("После смены:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
