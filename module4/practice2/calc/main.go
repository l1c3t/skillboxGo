package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("введите стоимость трех товаров через запятую")
	var costs string
	fmt.Scanln(&costs)

	items := strings.Split(costs, ",")
	fmt.Println(items)

	i := 0
	sum := 0
	//itemInt:=0
	for _, item := range items {
		i++
		fmt.Printf("вы ввели стоимость товара №%v:\t%v\n", i, item)
		itemInt, err := strconv.Atoi(item)
		if err != nil {
			break
		}
		sum += itemInt
	}

	fmt.Printf("С вас к оплате\t%v\n", sum)
	overall := float32(sum)
	if sum > 10000 {
		fmt.Println("Сумма Вашего заказа превышает 10000, поэтому Вам полагается скидка 10%")
		overall = 0.9 * float32(sum)
	}

	fmt.Printf("Итого к оплате:\t%v,\n", overall)

}
