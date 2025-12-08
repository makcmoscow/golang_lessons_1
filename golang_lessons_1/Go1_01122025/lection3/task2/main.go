// Программа запрашивает имя пользователя и возраст и выводит их. Возраст за вычетом одного года.
package main

import "fmt"


func main(){
	var value string

	fmt.Print("Введите значение, три символа: ")
	fmt.Scan(&value)
	fmt.Println(string(rune(value[0])))
	fmt.Println(string(rune(value[1])))
	fmt.Println(string(rune(value[2])))
// подразумевалось value / 100, value / 10 % 10, value % 10
}