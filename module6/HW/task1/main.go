/*
Задание 1. Написание простого цикла

Напишите программу, которая будет принимать от пользователя произвольное число и в цикле выводить на экран все значения от нуля до указанного числа.
*/
package main

import "fmt"

func main() {

	fmt.Println("Введите любое число:")

	var number int

	fmt.Scan(&number)

	for i := 0; i <= number; i++ {
		fmt.Println(i)
	}

}
