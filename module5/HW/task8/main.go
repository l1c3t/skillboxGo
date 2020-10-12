/*
## Игра “Угадывание числа” (дополнительно)
Что нужно сделать
Ну, и какой же компьютер без игр? Давайте научим его играть в “угадывание чисел”.
Пользователь загадывает число от 1 до 10 (включительно).
Программа пытается это число угадать, для этого она выводит число, а пользователь должен ответить угадала программа,
загаданное число больше или меньше.
Советы и рекомендации
Программа не должна делать больше 4 попыток в процессе угадывания.

*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Загадайте число от 0 до 10:")

	var number int

	fmt.Scan(&number)

	guessed := true
	guessTry := 0
	try := 0
	moreOr := ""

	for guessed {
		try = guessNumber(guessTry, moreOr)
		fmt.Println("Компьютер предполагает что это число:", try)
		if try == number {
			fmt.Println("Ура! Компьютер угадал!")
			guessed = false
		} else {
			fmt.Println("Больше/меньше?:")
			fmt.Scan(&moreOr)
			guessTry = try
		}
	}

}

func guessNumber(guess int, moreOr string) int {
	guessTry := 0
	if moreOr == "" {
		guessTry = rand.Intn(11)
	} else if moreOr == "+" {
		guessTry = guess + rand.Intn(10-guess+1)
	} else if moreOr == "-" {
		guessTry = rand.Intn(guess)
	}

	return guessTry
}
