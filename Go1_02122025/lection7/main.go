/*
Задача №1
Вводим любое натуральное число.
Нужно посчитать сумму цифр числа, с помощью цикла for
Пример:
In: 4521
Out: 12
*Задание 1.1: 4 + 5 + 2 + 1 = 12 - добавить к выводу сумму как выражение
*/
package main

import (
	"fmt"
	// "strings"

)
func main(){
var user_value int

fmt.Println("Enter your digital value: ")
fmt.Scanln(&user_value)
	v := 0
for {
	if user_value == 0 {
		break
	}
	v += user_value % 10
	user_value = user_value / 10
	
}
fmt.Println(v)
}
