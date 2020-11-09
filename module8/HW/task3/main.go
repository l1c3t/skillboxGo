/*
## Задание 3

В киоске с лимонадом, каждый лимонад стоит 5 долларов.
Клиенты стоят в очереди, чтобы купить у вас, и заказывают по одному (в порядке, указанном в счетах).
Каждый покупатель может купить только один лимонад и заплатить купюрой на 5, 10 или 20 долларов. Вы должны дать каждому покупателю сдачу с его купюры.

Обратите внимание, что сначала у вас нет сдачи.

Подсказка:
Верните true, в том случае, если вы можете предоставить каждому покупателю правильную сдачу.
Сигнатура функции lemonadeChange(bills []int) bool
где bills - это купюры, которые мы получаем от покупателей, по 1й купюре от каждого
*/

package main

//из задания непонятно что сделать-то нужно - только из подсказки
//функции не проходили
//массивы/срезы не проходили

import (
	"fmt"
)

func main() {

	var money int
	var change bool
	var bills []int
outLoop:
	for {
		fmt.Print("Введите номинал купюры\t")
		fmt.Scan(&money)

		bills = append(bills, money)
		if len(bills) > 1 {
			switch money {
			case 5, 10, 20:
				bills, change = lemonadeChange(bills)
				fmt.Println("В кассе:")
				fmt.Println(bills)
				fmt.Println("Получится сдать сдачу:")
				fmt.Println(change)
			case 0:
				break outLoop
			default:
				fmt.Println("Неверно введенное значение, попробуйте еще раз")
			}
		}
	}
}

func lemonadeChange(bills []int) ([]int, bool) {
	lastCash := bills[len(bills)-1] - 5
	sum := 0
	change := false

	switch lastCash {
	case 0:
		change = true
	case 5:
		for i, cash := range bills[:len(bills)-1] {
			if cash == 5 {
				bills[i] = 0
				change = true
			}
			sum = i
		}
	case 10:
		for i, cash := range bills[:len(bills)-1] {
			if cash == 10 {
				bills[i] = 0
				change = true
			} else if cash == 5 {
				bills[i] = 0
				sum++
			}
			if sum == 2 {
				change = true
			}
		}

	case 15:
		for i, cash := range bills[:len(bills)-1] {
			if cash == 10 {
				bills[i] = 0
				sum = +2
			} else if cash == 5 {
				bills[i] = 0
				sum++
			}
			if sum == 3 {
				change = true
			}
		}

	}

	var newBills []int
	for _, cash := range bills {
		if cash != 0 {
			newBills = append(newBills, cash)
		}

	}

	return newBills, change
}
