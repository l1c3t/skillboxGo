/*

Формат вывода (после прибытия на последнюю остановку):
Всего заработали: 100 руб.
Зарплата водителя: 25 руб.
Расходы на топливо: 20 руб.
Налоги: 20 руб.
Расходы на ремонт машины: 20 руб.
Итого доход: 15 руб.
*/

package main

import "fmt"

func main() {

	fmt.Println("Эта программа рассчитывает прибыль маршрутного такси в зависимости от количества остановок и входящих и выходящих людей")

	var passengers int
	allPassengers := 0
	profit := 0
	overallProfit := 0
	salary := 0
	fuelCost := 0
	tax := 0
	repairsCost := 0

	//1
	fmt.Println("Прибываем на остановку Улица Программистов. В салоне пассажиров:", allPassengers)
	fmt.Print("Сколько пассажиров зашло на остановке?\t")
	fmt.Scan(&passengers)
	allPassengers += passengers
	profit += 20 * passengers
	fmt.Println("Отправляемся с остановки Улица Программистов. В салоне пассажиров:", allPassengers)
	fmt.Println("-----------Едем---------")
	//2
	fmt.Println("Прибываем на остановку Улица Алгоритмов. В салоне пассажиров:", allPassengers)
	fmt.Print("Сколько пассажиров вышло на остановке?\t")
	fmt.Scan(&passengers)
	allPassengers -= passengers
	fmt.Print("Сколько пассажиров зашло на остановке?\t")
	fmt.Scan(&passengers)
	allPassengers += passengers
	profit += 20 * passengers
	fmt.Println("Отправляемся с остановки Улица Алгоритмов. В салоне пассажиров:", allPassengers)
	fmt.Println("-----------Едем---------")
	//3
	fmt.Println("Прибываем на остановку Улица s'pdfpasdk. В салоне пассажиров:", allPassengers)
	fmt.Print("Сколько пассажиров вышло на остановке?\t")
	fmt.Scan(&passengers)
	allPassengers -= passengers
	fmt.Print("Сколько пассажиров зашло на остановке?\t")
	fmt.Scan(&passengers)
	allPassengers += passengers
	profit += 20 * passengers
	fmt.Println("Отправляемся с остановки s'pdfpasdk. В салоне пассажиров:", allPassengers)
	fmt.Println("-----------Едем---------")
	//4
	fmt.Println("Прибываем на остановку бульвар Кошки Наступившей на клаву. Конечная. В салоне пассажиров:", allPassengers)
	fmt.Print("На остановке вышли все " + string(allPassengers) + " пассажира(ов)")
	allPassengers = 0

	fmt.Println("-----------Считаем денежки---------")

	salary = profit / 4
	fuelCost = profit / 5
	tax = profit / 5
	repairsCost = profit / 5
	overallProfit = profit - salary - fuelCost - tax - repairsCost
	fmt.Printf("Всего заработали:\t%d руб. \n", profit)
	fmt.Printf("Зарплата водителя:\t%d руб. \n", salary)
	fmt.Printf("Расходы на топливо:\t%d руб. \n", fuelCost)
	fmt.Printf("Налоги:\t%d руб. \n", tax)
	fmt.Printf("Расходы на ремонт машины:\t%d руб. \n", repairsCost)
	fmt.Printf("Итого доход:\t%d руб. \n", overallProfit)

}
