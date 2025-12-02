/*
Задача № 2. Получить реверсную запись трехзначного числа
Пример:
вход: 346, выход: 643
вход: 120, выход: 021
вход: 100, выход: 001
*/
package main

import "fmt"

func main(){
	var value string
	fmt.Println("Введите трехзначное(!) число. Знаки после первых трех будут отброшены!")
	fmt.Scanln(&value)
	fmt.Println(string(value[2])+string(value[1])+string(value[0]))
}