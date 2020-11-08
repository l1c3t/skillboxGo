/*
Задание 2. Сумма двух чисел по единице

Напишите программу, которая запрашивает у пользователя два числа и складывает их в цикле следующим образом: берёт первое число и прибавляет к нему по единице, пока не достигнет суммы двух чисел.
*/
package main

import "fmt"

func main() {

	fmt.Println("Введите первое число:")
	var number1 int
	fmt.Scan(&number1)

	fmt.Println("Введите второе число:")
	var number2 int
	fmt.Scan(&number2)

	currentNumber := number1

	for {
		fmt.Println(currentNumber)

		if currentNumber == number1+number2 {
			fmt.Println("Условие выполнено")
			break
		}
		currentNumber++
	}

}
