/*
## Задание 6*. Движение лифта

В доме 24 этажа. Лифт должен ходить по кругу, пока не доставит всех пассажиров на первый этаж.
Три пассажира ждут на четвёртом, седьмом и десятом этажах. При движении вверх лифт не должен останавливаться, при движении вниз — собирать всех, но не более двух человек в лифте.
При этом лифт каждый раз доезжает до самого верхнего этажа и только после этого начинает движение вниз. Напишите программу, которая доставит всех пассажиров на первый этаж.
*/

package main

import "fmt"

func main() {

	liftCount := 24
	liftCap := 2
	pass10 := 1
	pass7 := 1
	pass4 := 1

	for {
		fmt.Println("этаж", liftCount)

		if (liftCount == 10) && (liftCap != 0) && (pass10 != 0) {
			fmt.Println("Забираем на 10-м")
			liftCap--
			fmt.Println("В лифте осталось мест:", liftCap)
			pass10 = 0
		}
		if (liftCount == 7) && (liftCap != 0) && (pass7 != 0) {
			fmt.Println("Забираем на 7-м")
			liftCap--
			fmt.Println("В лифте осталось мест:", liftCap)
			pass7 = 0
		}
		if (liftCount == 4) && (liftCap != 0) && (pass4 != 0) {
			fmt.Println("Забираем на 4-м")
			liftCap--
			fmt.Println("В лифте осталось мест:", liftCap)
			pass4 = 0
		}

		if liftCount == 1 {
			fmt.Println("едем вверх")
			liftCount = 24
			liftCap = 2
			if pass4 == 0 {
				fmt.Println("Довезли всех пассажиров")
				break
			}
		}
		liftCount--

	}

}
