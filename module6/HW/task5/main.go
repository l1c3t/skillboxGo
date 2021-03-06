/*
Задание 5*. Задача на постепенное наполнение корзин разной ёмкости (if и continue и break)

Представьте, что у вас есть три корзины разной ёмкости. Пользователю предлагается ввести, какое количество яблок помещается в каждую корзину.
После этого программа должна заполнить все корзины по очереди, учитывая, какие корзины уже заполнены, строго соблюдая очерёдность заполнения и добавляя по одному яблоку в каждой итерации.
Пример: пользователь решил, что у корзин будет ёмкость на шесть, четыре и девять яблок.
Программа должна заполнить корзину 1 и в следующей итерации перейти к корзине 2, далее в следующей итерации перейти к корзине 3.
Если очередная корзина уже заполнена, программа должна переходить к следующей по очереди, и так по кругу, пока не заполнит все.
*/

package main

import "fmt"

func main() {

	fmt.Println("Емкость первой корзинки")
	var n1 int
	fmt.Scan(&n1)

	fmt.Println("Емкость второй корзинки")
	var n2 int
	fmt.Scan(&n2)

	fmt.Println("Емкость третьей корзинки")
	var n3 int
	fmt.Scan(&n3)

	var overall int
	//break не стал ставить. можно было бы перенести условие из цикла в тело цикла и воткнуться с break, но так вроде красиво
	for i := 1; i <= (n1 + n2 + n3); i++ {
		overall = i
		if i <= n1 {
			fmt.Println("Складываем в первую корзину", i, "-е яблоко")
			continue
		}
		if i-n1 <= n2 {
			fmt.Println("Складываем во вторую корзину", i-n1, "-е яблоко")
			continue
		}
		if i-(n1+n2) <= n3 {
			fmt.Println("Складываем в третью корзину", i-(n1+n2), "-е яблоко")
			continue
		}

	}
	fmt.Println("всего яблок", overall)
	fmt.Println("закончились свободные кассы")

}
