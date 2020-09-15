/*
Задача 1. Минимум

Напишите программу, которая ищет минимум из двух чисел.

Подсказка: учтите, что числа могут быть равны!
*/
package main

import "fmt"

func main() {

	fmt.Println("Введите первое число")
	var a int
	fmt.Scan(&a)

	fmt.Println("Введите второе число")
	var b int
	fmt.Scan(&b)

	if a > b {
		fmt.Println("первое число больше второго")
	} else if b > a {
		fmt.Println("Второе число больше второго")
	} else {
		fmt.Println("Числа равны")
	}
}
