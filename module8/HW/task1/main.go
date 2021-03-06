/*
## Задание 1

Пользователь вводит месяц, программа должна вывести, на какое время года (зима, весна, лето, осень) этот месяц выпадает.

Как группировать:
декабрь, январь, февраль - зима
март, апрель, май - весна
июнь, июль, август - лето
сентябрь, октябрь, ноябрь - осень
*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	var month string
	for {

		fmt.Print("Ведите месяц года:")
		fmt.Scan(&month)

		switch strings.ToLower(month) {
		case "декабрь", "январь", "февраль":
			fmt.Println("Зима")
		case "март", "апрель", "май":
			fmt.Println("Весна")
		case "июнь", "июль", "август":
			fmt.Println("Лето")
		case "сентябрь", "октябрь", "ноябрь":
			fmt.Println("Осень")
		case "надоело":
			break
		default:
			fmt.Println("Я не знаю такого месяца")
			break
		}
	}

}
