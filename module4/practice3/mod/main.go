/*
Задача 3. Проверка на чётное число

Напишите программу, которая проверяет, четное ли число ввел пользователь.

Подсказка: для проверки чётности используйте оператор %. Он вычисляет остаток от деления двух чисел. Пример:

remainder := x % 2 // вычисляет остаток от деления `x` на 2
*/
package main

import "fmt"

func main() {
	fmt.Println("Введите число")
	var a int
	fmt.Scan(&a)
	if a == 0 {
		fmt.Println("Вы ввели 0")
	} else if (a % 2) != 0 {
		fmt.Println("Число нечетное")
	} else if (a % 2) == 0 {
		fmt.Println("Число четное")
	}

}