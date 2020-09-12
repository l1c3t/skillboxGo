package main

import (
	"fmt"
)

func main() {
	fmt.Println("Эта программа считает рост бамбука за сутки при наличии гуссениц")

	var heigth float32 = 100
	var growthSpeed float32 = 50
	var caterpillarAte float32 = 20

	overallHeight := (growthSpeed-caterpillarAte)*2 + growthSpeed/2 + heigth

	fmt.Printf("/n По неправильной формуле в середине третьего дня бамбук станет:\t%vсм.\n\n", overallHeight)

	//усложнение - считаем от любого дня, гусченицы, скорости роста

	fmt.Println("--------------усложняем-----------")
	fmt.Print("Введите сколько дней мы наблюдаем за бамбуком\t")
	var days float32
	fmt.Scan(&days)

	fmt.Print("Введите начальную высоту бамбука\t")
	fmt.Scan(&heigth)

	fmt.Print("Введите скорость роста бамбука\t")
	fmt.Scan(&growthSpeed)

	fmt.Print("Введите сколько съедают гусеницы\t")
	fmt.Scan(&caterpillarAte)

	night := int(days - 1)

	overallHeight = growthSpeed*days - caterpillarAte*float32(night) + heigth
	fmt.Printf("\nЧерез %v дня бамбук станет:\t%vсм.\n\n", days, overallHeight)

	/*
		рассчет до...
	*/

	fmt.Println("--------------рассчет высоты до...-----------")
	fmt.Print("Введите какой высоты мы хотим вырастить бамбук\t")
	var targetHeight float32
	fmt.Scan(&targetHeight)
	fmt.Println(targetHeight)

	overallHeight = 0
	days = 0
	night = 0
	for (targetHeight - overallHeight) > 0.1 {

		days += 0.1
		if days >= 2 {
			night = int(days - 1)
		}

		overallHeight = growthSpeed*days - caterpillarAte*float32(night) + heigth
	}
	fmt.Printf("Для достижения высоты бамбука в %vсм необходимо %v дней", targetHeight, days)
}
