/*
Напишите программу, которая запрашивает у пользователя три числа и сообщает, есть ли среди них число, большее, чем 5.
*/
package main

import "fmt"

func main() {

	fmt.Println("Эта программа запрашивает у Вас три числа и сообщает, есть ли среди них число, большее, чем 5")

	fmt.Print("Число 1:\t")
	var num1 int
	fmt.Scan(&num1)

	fmt.Print("Число 2:\t")
	var num2 int
	fmt.Scan(&num2)

	fmt.Print("Число 3:\t")
	var num3 int
	fmt.Scan(&num3)

	if (num1 > 5) || (num2 > 5) || (num3 > 5) {
		fmt.Println("Среди этих трех чисел есть число, большее чем 5")
	} else {
		fmt.Println("Все три числа меньше или равны 5")
	}
}
