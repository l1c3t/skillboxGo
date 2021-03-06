/*
Определение координатной плоскости точки
Что нужно сделать
В различных программах часто приходится работать с координатами, это могут быть системы проектирования или игры.
Ключевой момент в таких программах, это работа с системой координат.
Давайте поможем пользователю определить, к какой координатной четверти принадлежит точка.
Пользователь вводит числа x, y, а программе необходимо вывести, какой координатной четверти принадлежит точка с
 координатами (x, y), при условии:
если обе координаты положительны, то точка находится в первой четверти координатной плоскости;
если координата х отрицательна, а координата у положительна, то точка находится во второй четверти;
если обе координаты отрицательны, то число находится в третьей четверти;
если координата х положительна, а координата у отрицательна, то точка лежит в четвертой четверти.
Советы и рекомендации - Использовать логическое умножение.
*/

package main

import "fmt"

func main() {
	fmt.Println("Введите координаты x, а затем y:")
	var x float64
	var y float64

	fmt.Scan(&x)
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		fmt.Println("Точка находится в первой четверти координатной плоскости")
	} else if x < 0 && y > 0 {
		fmt.Println("Точка находится во второй четверти координатной плоскости")
	} else if x > 0 && y < 0 {
		fmt.Println("Точка находится в четвертой четверти координатной плоскости")
	} else if x < 0 && y < 0 {
		fmt.Println("Точка находится в третьей четверти координатной плоскости")
	} else {
		fmt.Println("Вы указали центр координат")
	}

}
