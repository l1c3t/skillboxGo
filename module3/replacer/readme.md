Для всех задач обязательно:
правильное именование переменных
текстовый интерфейс
описание программы (приветствие или что она делает)
при вводе данных
при выводе данных
вынесение ключевых сущностей в переменные (в первую очередь те, которые могут меняться, и те, которые используются более одного раза в коде)
допускается использование нижних подчеркиваний вместо пробелов в значениях, полученных от пользователя. 

Задача 3. Обмен местами
 
Есть код программы с двумя переменными типа int и выводом этих переменных на экран. Напишите программу, которая меняет значения переменных местами, то есть нужно добиться того, чтобы в переменной "a" лежало значение "b".

package main
 
import (
  "fmt"
)
 
func main() {
 a := 42
 b := 153
 
 fmt.Println("a:", a)
 fmt.Println("b:", b)
}