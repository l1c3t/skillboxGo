//Напишите программу, которая вычисляет модуль числа.

package main

import "fmt"

func main() {

	fmt.Println("--------Программа вычисляет модуль введенного числа--------")

	fmt.Print("введите число, для которого необходимо определить модуль:\t")
	var x int
	fmt.Scan(&x)
	y := x

	if x < 0 {
		y = -x
	}

	fmt.Printf("Для введенного числа %v модулем %v\n", x, y)
}
