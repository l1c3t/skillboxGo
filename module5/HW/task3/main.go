/*
Проверить, что есть совпадающие числа
Что нужно сделать
Пользователь при вводе может ошибиться, и многие программы пытаются привлечь внимание к таким ошибкам.
Например, текстовые редакторы подсвечивают ошибки в словах, электронные таблицы — в формулах.
Давайте реализуем программу, запрашивающую у пользователя три числа, и выводящую подсказку,
если два или более числа совпадают.
Советы и рекомендации
Использовать логическое сложение.
*/
package main

import "fmt"

func main() {
	fmt.Println("Введите три числа:")

	var x float64
	var y float64
	var z float64

	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&z)

	if x == z || x == y || y == z {
		fmt.Println("Два или более числа совпадают")
	} else {
		fmt.Println("Совпадающих чисел нет")
	}

}
