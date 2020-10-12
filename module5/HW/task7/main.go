/*
## Счастливый билет
Что нужно сделать
Все мы в детстве, да и не только в детстве, получив бумажный билет, проверяли, а не является ли он “счастливым”.
Давайте напишем программу, в которую пользователь будет вводить четырехзначный номер билета, а программа будет выводить,
является ли он счастливым (сумма старших двух разрядов равна сумме двух младших разрядов) или даже зеркальным.
Например:
1234 -> обычный билет
3425 -> счастливый билет
1221 -> зеркальный билет
Советы и рекомендации
При решении задачи необходимо применить целочисленное деление и остаток от деления.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Введите номер билета:")

	var ticket string

	fmt.Scan(&ticket)

	var ticketsNumbers [4]int

	for i := 0; i <= 3; i++ {
		var err error
		ticketsNumbers[i], err = strconv.Atoi(string(ticket[i]))
		if err != nil {
			break
		}

	}

	fmt.Println(ticketsNumbers)
	if len(ticket) == 4 {
		if (ticketsNumbers[0] == ticketsNumbers[3]) && (ticketsNumbers[1] == ticketsNumbers[2]) {
			fmt.Println("Билет зеркальный")
		} else if ticketsNumbers[0]+ticketsNumbers[1] == ticketsNumbers[2]+ticketsNumbers[3] {
			fmt.Println("Счастливый билетик!")
		} else {
			fmt.Println("Обычный билетик")
		}
	} else {
		fmt.Println("введен неверный номер билета")
	}

}
