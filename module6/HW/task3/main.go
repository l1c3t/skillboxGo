/*
Задание 3. Расчёт суммы скидки

Напишите программу, которая принимает на вход цену товара и скидку. Посчитайте и верните на экран сумму скидки.
Дополнительное условие: скидка не превышает 30% от цены товара и не больше 2 000 рублей.
*/

//на самом деле странное условие - не совсем понятно как обрабатывать условия - проще же без цикла решить, приравнивая к максимуму, но домашка про циклы, поэтому так.
package main

import "fmt"

func main() {

	fmt.Println("Введите цену товара:")
	var cost int
	fmt.Scan(&cost)

	var discount int

	for {

		fmt.Println("Введите скидку:")
		fmt.Scan(&discount)

		fmt.Println("-------------------")


		//вот тут умножение на 100 - результат танцев с int. можно к float, но "мы же еще не проходили"
		if ((discount*100/cost*100)/100 <= 30) && (discount <= 2000) {
			fmt.Println("Скидка предоставлена успешно")
			break
		}

		fmt.Println("Скидка введена с ошибкой")
		fmt.Println("условия по скидке:")
		fmt.Println("Скидка должна быть меньше 30%")
		fmt.Println("Скидка должна быть меньше 2000")

	}

}
