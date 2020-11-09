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

	change := false

	nominals := map[int]int{
		5:  0,
		10: 0,
	}

	for i := 0; i <= (len(bills) - 1); i++ {
		if bills[i] == 5 {
			nominals[5]++
		} else if bills[i] == 10 {
			nominals[10]++
		}
		fmt.Println(bills[:i+1])
		fmt.Println(nominals)

		for j := 0; j <= i; j++ {
			lastCash := bills[j] - 5

			switch lastCash {
			case 0:
				change = true
			case 5:
				if nominals[5] > 0 {
					change = true
					nominals[5]--
				} else {
					change = false
				}
			case 15:
				if nominals[10] > 0 && nominals[5] > 0 {
					change = true
					nominals[5]--
					nominals[10]--
				} else if nominals[5] > 2 {
					change = true
					nominals[5] = -2
				} else {
					change = false
				}
			default:
				change = false

			}

		}
		fmt.Println(nominals)
	}

	return change
}
