package main

import "fmt"

func main() {
	fmt.Println("Привет! Расскажи немножко о себе и  может быть я предложу тебе что-нибудь интересное и может выгодное")

	fmt.Print("Для начала скажи как тебя зовут:\t")
	var name string
	fmt.Scanln(&name)

	fmt.Print("Отлично!\nА теперь подскажи с какой ты планеты:\t")
	var planet string
	fmt.Scanln(&planet)

	fmt.Print("Это в звездной системе... запамятовал... /подсказать собеседнику:\t")
	var solarSystem string
	fmt.Scanln(&solarSystem)

	fmt.Print("Сколько денег ты обычно берешь за публтичные выступления и агитацию?:\t")
	var reward int
	fmt.Scanln(&reward)

	fmt.Printf("\nКак вам, %s, известно, мы - раса мирная, поэтому на наших военных кораблях используются наемники с других планет. Система набора отработана давным-давно. Обычно мы приглашаем на наши корабли людей с планеты %s системы %s.\n", name, planet, solarSystem)

	fmt.Printf("\nНо случилась неприятность - в связи с большими потерями, в последнее время, престиж профессии сильно упал, и теперь не так-то просто завербовать призывников. Главный комиссар планеты %s, впрочем, отлично справлялся, но недавно его наградили орденом Сутулого с закруткой на спине, и его на радостях парализовало! Призыв под угрозой срыва!", planet)

	fmt.Printf("\n%s, вы должны прибыть на планету %s системы %s и помочь выполнить план призыва. Мы готовы выплатить вам премию в %d кредитов за эту маленькую услугу.", name, planet, solarSystem, reward)

}
