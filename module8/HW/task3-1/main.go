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

//Второе решение, как мне кажется наиболее близкое к заданию

//из задания непонятно что сделать-то нужно - только из подсказки
//функции не проходили
//массивы/срезы не проходили
//map'ы мы еще не проходили
//все пришлось читать)))

import (
	"fmt"
)

func main() {

	var money int
	var change bool
	var bills []int
outloop:
	for {
		fmt.Print("Введите номинал купюры\t")
		fmt.Scan(&money)

		switch money {
		case 5, 10, 20:
			bills = append(bills, money)
		case 0:
			break outloop
		default:
			fmt.Println("Неверно введенное значение, попробуйте еще раз")
		}
	}

	change = lemonadeChange(bills)
	fmt.Println("В кассе:")
	fmt.Println(bills)
	fmt.Println("Получится сдать сдачу:")
	fmt.Println(change)

}

func lemonadeChange(bills []int) bool {
	lastCash := bills[len(bills)-1] - 5
	change := false

	nominals := map[int]int{
		5:  0,
		10: 0,
	}

	for i := 1; i < (len(bills) - 1); i++ {

		for _, cash := range bills[:i-1] {
			if cash == 5 {
				nominals[5]++
			} else if cash == 10 {
				nominals[10]++
			}
		}

		switch lastCash {
		case 0:
			change = true
		case 5:
			if nominals[5] > 0 {
				change = true
			}
		case 10:
			if nominals[5] > 1 {
				change = true
			} else if nominals[10] > 0 {
				change = true
			}

		case 15:
			if nominals[5] > 2 {
				change = true
			} else if (nominals[10] > 1) && (nominals[5] > 0) {
				change = true
			}
		}
	}

	return change
}
