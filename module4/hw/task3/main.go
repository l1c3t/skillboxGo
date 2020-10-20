/*
Пользователи банкомата хотят снимать деньги. Но банкомат умеет выдавать только купюры по 100 рублей,
а максимальная сумма снятия — 100 000 рублей.

Напишите программу, которая проверяет допустимость суммы средств, которую ввёл пользователь.
Сумма не должна быть меньше 100 и больше 100 000 руб.
*/

package main

import "fmt"

func main() {

	fmt.Println("Эта программа проверяет сможете ли Вы снять интересующую сумму в банкомате")

	moneyNominal := 100
	moneyMaximumPerOne := 100000

	fmt.Print("Сукмму снятия:\t")
	var moneySum int
	fmt.Scan(&moneySum)

	if (moneySum >= moneyNominal) && (moneySum <= moneyMaximumPerOne) {
		fmt.Println("Вы можете снять сумму в", moneySum)
	} else {
		fmt.Println("Вы не можете снять сумму в", moneySum, "попробуйте от 100 и до 1000000")
	}

}