/*
Проверить, что одно из чисел положительное
Что нужно сделать
Проверка пользовательского ввода на различные ограничения является частой задачей.
Давайте попросим пользователя ввести 3 числа и проверим, что хотя бы одно является положительным.
Результат проверки необходимо сообщить пользователю.
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

	if x > 0 || y > 0 || z > 0 {
		fmt.Println("Одно из чисел положительное")
	} else {
		fmt.Println("все числа отрицательные")
	}

}
