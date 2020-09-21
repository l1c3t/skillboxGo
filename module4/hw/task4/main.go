/*
Напишите программу, которая запрашивает у пользователя три числа и выводит количество чисел, которые больше, либо равны 5.
*/
package main

import "fmt"

func main() {

	fmt.Println("Эта программа запрашивает у Вас три числа и сообщает, сколько из них, больше или равно 5")

	fmt.Print("Число 1:\t")
	var num1 int
	fmt.Scan(&num1)

	fmt.Print("Число 2:\t")
	var num2 int
	fmt.Scan(&num2)

	fmt.Print("Число 3:\t")
	var num3 int
	fmt.Scan(&num3)

	sum := 0
	if num1 >= 5 {
		sum++
	}
	if num2 >= 5 {
		sum++
	}
	if num3 >= 5 {
		sum++
	}

	fmt.Printf("Среди этих трех чисел есть %d чисел, которые больше или равны 5", sum)

}
