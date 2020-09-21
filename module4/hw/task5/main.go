/*
В ресторане действуют следующие правила:

по понедельникам должна применяться скидка 10% на всё меню, потому что понедельник — день тяжёлый!

по пятницам, если сумма чека превышает 10 000 руб., включается дополнительная скидка в размере 5%;

если число гостей в одной компании превышает 5 человек, автоматически включается надбавка на обслуживание 10%.

Напишите программу, которая запрашивает день недели, число гостей и сумму чека и рассчитывает сумму к оплате.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	monDiscount := 0.1
	fridayCheck := 10000
	friDiscount := 0.05
	servicePrice := 0.1
	servicePriceMens := 5
	weekDayNumber := 0

	fmt.Println("Эта программа рассчитывает сумму к оплате")

	fmt.Print("введите день недели на русском:\t")
	var weekDay string
	fmt.Scan(&weekDay)
	weekDay = strings.ToLower(weekDay)

	switch weekDay {
	case "понедельник":
		weekDayNumber = 1
	case "пятница":
		weekDayNumber = 5

	default:
		weekDayNumber = 2
	}

	fmt.Print("введите число гостей:\t")
	var people int = 0
	fmt.Scan(&people)

	fmt.Print("введите сумму чека:\t")
	var price int = 0
	fmt.Scan(&price)

	priceWithoutMensCount := 0.0
	if weekDayNumber == 1 {
		priceWithoutMensCount = float64(price) * (1 - monDiscount)
	} else if (weekDayNumber == 5) && (price > fridayCheck) {
		priceWithoutMensCount = float64(price) * (1 - friDiscount)
	}

	priceWithMens := priceWithoutMensCount
	if people > servicePriceMens {
		priceWithMens = (1 + servicePrice) * float64(price)
	}

	fmt.Println("Итоговая сумма к оплате:", priceWithMens)
}
