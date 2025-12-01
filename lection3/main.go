// Программа запрашивает имя пользователя и возраст и выводит их. Возраст за вычетом одного года.
package main

import "fmt"


func main(){
	var name string
	var age int

	fmt.Print("Введите ваше имя: ")
	fmt.Scan(&name)
	fmt.Print("Введите ваш возраст: ")
	fmt.Scan(&age)
	fmt.Println("Привет, ", name)
	fmt.Println("Ваш возраст: ", age-1)
}