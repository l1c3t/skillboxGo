/*
## Задание 2

Пользователь вводит будний день недели в сокращенной форме(пн, вт, ср, чт, пт) и получает развернутый список всех
последующих рабочих дней, включая пятницу.
Пример:
пользователь вводит:
вт
программа дает ответ:
вторник
среда
четверг
пятница
*/

package main

import (
	"fmt"
	"strings"
)

func main() {

	var month string
OutLoop:
	for {
		fmt.Print("чтобы получить список будних дней, следующих за указанным\n")
		fmt.Print("Ведите сокращенно день недели:")
		fmt.Scan(&month)

		switch strings.ToLower(month) {
		case "сб", "вс", "пн", "cуб", "воскр":
			fmt.Println("понедельник")
			fallthrough
		case "вт":
			fmt.Println("вторник")
			fallthrough
		case "ср":
			fmt.Println("среда")
			fallthrough
		case "чт":
			fmt.Println("четверг")
			fallthrough
		case "пт":
			fmt.Println("пятница")
		case "надоело":
			break OutLoop
		default:
			fmt.Println("Я не знаю такого дня недели")
			break
		}
	}
}
