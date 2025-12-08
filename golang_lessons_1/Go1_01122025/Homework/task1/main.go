/*
Задача №1
Вход:
    расстояние(50 - 10000 км),
    расход в литрах (5-25 литров) на 100 км и
    стоимость бензина(константа) = 48 руб

Выход: стоимость поездки в рублях
*Проверка условий по желанию
*/
package main

import (
    "fmt"
)

func main(){
    var distance, lpk float32 
    const cost = 48
    fmt.Println("Введите дистанцию (50-10000 км): ")
    fmt.Scanln(&distance)
    fmt.Println("Введите расход (5-25 литров на сотню): ")
    fmt.Scanln(&lpk)
    if distance > 10000 {
        fmt.Println("Слишком большая дистанция, ошибка ввода?")
        return
    } 
    if distance < 50 {
        fmt.Println("Слишком маленькая дистанция, ошибка ввода?")
        return
    } 
    if lpk > 25 {
        fmt.Println("Слишком большой расход, ошибка ввода?")
        return
    } 
    if lpk < 5 {
        fmt.Println("Слишком маленький расход, ошибка ввода?")
        return
    } 
    result := distance*lpk*cost/100
    fmt.Println("Стоимость поездки: ", result)
}