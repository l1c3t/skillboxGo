/*
Задача 5. Кратность числа

Напишите программу, которая проверяет, делится ли одно число на другое без остатка.
*/

package main

import "fmt"

func main() {
	fmt.Println("Введите число")
	var a int
	fmt.Scan(&a)
	fmt.Println("Введите на которое будем делить число")
	var b int
	fmt.Scan(&b)

	if a == 0 {
		fmt.Println("Вы ввели 0")
	} else if (b) == 0 {
		fmt.Println("На ноль делить нельзя")
	} else if (a % b) != 0 {
		fmt.Println("Первое чсисло не кратно второму числу")
	} else if (a % b) == 0 {
		fmt.Println("Первое чсисло кратно второму числу")
	}

}
