/*
Задача 4. Калькулятор опыта

Напишите программу, которая определяет уровень персонажа в компьютерной игре.
Пользователь вводит число "очков опыта", а программа вычисляет уровень.
Новый уровень дается при достижении 1000, 2500 и 5000 "очков опыта".
Начальный уровень равен 1, а максимальный 4 (при достижении 5000 очков и больше).
*/

package main

import "fmt"

func main() {
	fmt.Println("Введите количество опыта")
	var exp int
	fmt.Scan(&exp)

	if exp >= 5000 {
		fmt.Println("вы получили уровень: 4. Максимальный")
	} else if exp >= 2500 {
		fmt.Println("вы получили уровень: 3. есть куда расти")
	} else if exp >= 1000 {
		fmt.Println("вы получили уровень: 2. начало пути положено")
	} else {
		fmt.Println("Ваш уровень: 1. Вы только начали")
	}

}