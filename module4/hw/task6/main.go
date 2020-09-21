/*
Перед старостой группы стоит задача разделить весь курс, состоящий из N студентов, на K групп.
Напишите программу, которая поможет старосте сделать это: он вводит N, K и порядковый номер студента,
а программа определяет, в какую группу он попадает.

Подсказка: в одну группу могут попадать студенты из разных частей списка.
*/
package main

import "fmt"

func main() {

	fmt.Print("Введите количество студентов:\t")
	var n int
	fmt.Scan(&n)

	fmt.Print("Введите количество групп:\t")
	var k int
	fmt.Scan(&k)

	studentsInGroup := (n - n%k) / k

	fmt.Println("-------------------------------------")
	fmt.Printf("В каждой группе минимум %d студентов\n", studentsInGroup)
	fmt.Println("-------------------------------------")

	fmt.Print("Введите  порядковый номер студента:\t")
	var studentNumber int
	fmt.Scan(&studentNumber)

	groupForStudent := studentNumber % k
	if studentNumber <= k {
		groupForStudent = studentNumber
	}

	fmt.Println(groupForStudent)
}
