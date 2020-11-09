/*
## 1. Зеркальные билеты

Что нужно сделать

Вывести сколько билетов находится среди всех шестизначных номеров от 100000 до 999999.
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var mirrorTicketCount int

	for i := 100000; i <= 999999; i++ {
		ticket := i

		stringTicket := strconv.Itoa(ticket)
		var ticketsNumbers [6]int
		for j := 0; j <= 5; j++ {

			var err error
			ticketsNumbers[j], err = strconv.Atoi(string(stringTicket[j]))
			if err != nil {
				break
			}

		}
		if (ticketsNumbers[0] == ticketsNumbers[5]) && (ticketsNumbers[1] == ticketsNumbers[4]) && (ticketsNumbers[2] == ticketsNumbers[3]) {
			mirrorTicketCount++
		}

	}
	fmt.Println("всего зеркальных билетов", mirrorTicketCount)

}
