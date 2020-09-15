/*
Задача 8. Зарплата* (сложная задача)

В отделе маркетинга работают 3 сотрудника. У каждого — разные зарплаты.
Напишите программу, которая вычисляет разницу между самой высокой и низкой зарплатой сотрудника,
а также среднюю зарплату отдела.
*/
package main

import "fmt"

func main() {

	fmt.Println("Введите зарплату первого сотрудника")
	var a int
	fmt.Scan(&a)

	fmt.Println("Введите зарплату второго сотрудника")
	var b int
	fmt.Scan(&b)

	fmt.Println("Введите зарплату третьего сотрудника")
	var c int
	fmt.Scan(&c)
	min := 0
	max := 0
	half := (a + b + c) / 3
	//минимальное значение
	if (a < b) && (a < c) {
		min = a
	} else if (b < a) && (b < c) {
		min = b
	} else if (c < a) && (c < b) {
		min = b
	}
	//максимальное значение
	if (a > b) && (a > c) {
		max = a
	} else if (b > a) && (b > c) {
		max = b
	} else if (c > a) && (c > b) {
		max = c
	}

	fmt.Println("Минимальная зарплата составляет:", min)
	fmt.Println("Маскимальная зарплата составляет:", max)
	fmt.Println("Средняя зарплата:", half)

}
