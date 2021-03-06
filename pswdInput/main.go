/*
1. Мы же готовимся к тому, чтобы стать настоящими веб-разработчиками, а что должно быть на каждом сайте? Конечно — регистрация!
 Самое время вам сделать первый шаг в этом направлении!

Напишите программу, которая принимает на вход имя пользователя и пароль как строки, а также его возраст целым числом,
а потом выведите эти данные. Не забывайте делать подсказки пользователю, что ему нужно ввести.

Например, если пользователь введет имя username, пароль password, а возраст 32, то программа должна вывести такой текст:


Поздравляю, username, теперь вы зарегистрированы!
Ваш пароль: password
Ваш возраст: 32


Конечно, выводить пароль небезопасно, вдруг его кто-нибудь подсмотрит. Но мы же с вами только начинаем, и можем позволить
 себе такие допущения.
*/
package main

import "fmt"

func main() {
	fmt.Println("Введите имя пользователя")
	var login string
	fmt.Scan(&login)

	fmt.Println("Введите пароль")
	var password string
	fmt.Scan(&password)

	fmt.Println("Введите Ваш возраст")
	var age int
	fmt.Scan(&age)

	fmt.Println("Поздравляю,", login, ", теперь вы зарегистрированы! \n Ваш пароль:", password, "\nВаш возраст:", age)
}
