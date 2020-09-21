/*
В последнем уроке мы писали программу, вычисляющую сумму налога по прогрессивной шкале в зависимости от полученного заработка:
13% на доход до 10000 руб., 20% на доход от 10000 до 50000 руб., 30% на доход выше 50000 руб.

Однако, во многих странах, использующих такую шкалу, эта сумма вычисляется более сложным способом.
А именно, налоговая ставка 30% на доход выше 50000 руб. означает, что 30% уплачивается не со всей суммы,
а лишь с той ее части, которая превосходит 50000 руб. Аналогично, ставка 20% на доход от 10000 до 50000 руб.
обязывает уплатить 20% лишь с той части суммы, которая превосходит 10000 руб., но не превосходит 50000 руб.

Так, например, с дохода 100000 руб. придется заплатить такой налог:

30% *(100000-50000) + 20%* (50000-20000) + 13% * 10000 = 15000+6000+1300=22300.

А с дохода 30000 — такой: 20% *(30000-20000) + 13%* 10000 = 2000+1300=3300.

Напишите программу, которая спрашивает у пользователя его доход и высчитывает сумму налога для него
по вышеописанным правилам.
*/
package main

import "fmt"

func main() {

	firstStep := 10000
	percentForFirstStep := 0.13
	secondStep := 50000
	percentForSecondStep := 0.2
	percentForThirdStep := 0.3

	fmt.Println("Программа считает сумму налога")

	fmt.Print("Введите сумму дохода:\t")
	var profit int
	fmt.Scan(&profit)

	nalog := 0.0
	summNalog := 0.0

	if profit >= secondStep {
		summNalog = percentForThirdStep * float64(profit-(secondStep-1))
		nalog = nalog + summNalog
		fmt.Println("Ваш налог на сумму свыше 50000:", summNalog)
		profit = secondStep - 1
	}
	if (profit < secondStep) && (profit >= firstStep) {
		summNalog = percentForSecondStep * float64(profit-(firstStep-1))
		nalog = nalog + summNalog
		fmt.Println("Ваш налог на сумму свыше 10000 и до 50000:", summNalog)
		profit = firstStep - 1
	}
	if profit < firstStep {
		nalog = nalog + percentForFirstStep*float64(profit)
		fmt.Println("Ваш налог на сумму до 10000:", percentForFirstStep*float64(profit))

	}
	fmt.Println("Общая сумма налога для Вас:", nalog)
}
