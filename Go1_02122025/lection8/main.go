/*
=======
Задачи:
=======
3.1 Пользователь вводит числа a и b (b больше a).
    Вывести все целые числа, расположенные между ними.
3.2 Доработать предыдущую задачу так, чтобы выводились только числа,
    делящиеся на 5 без остатка.
3.3 Пользователь вводит число. Вывести таблицу умножения на это число (от 1 до 10)
3.4 В цикле получать от пользователя оценки по четырём экзаменам.
    Вывести сумму набранных им баллов.
    Функцию fmt.Scan() в коде используем только один раз.
3.5 В бесконечном цикле приложение запрашивает у пользователя числа.
    Ввод завершается, как только пользователь вводит число "-1".
    После завершения ввода приложение выводит сумму чисел.
    Используем конструкцию:
    for {
      // body
    }
3.6 Получить от пользователя натуральное число, посчитать сумму цифр в нём.
    Решить с помощью индексов, т.е. работаем с числом, как со строкой.
3.7 Вводим строку без знаков препинания(5 слов).
    Найти самое длинное слово в строке и вывести сколько в нем букв.
Пример:
In: Скажи как дела в учебе
Out: учебе, 5
In: Закрепляем циклы в языке Golang
Out: Закрепляем, 10
*/
package main
import (
	"fmt"
    // "unicode/utf8"
    // "strconv"
)

//3.1
// func main(){
// 	var a, b int
//     fmt.Println("Enter first digit\n")
// 	fmt.Scan(&a)
//     fmt.Printf("First digit is %d\n", a)
//     fmt.Println("Enter second digit, greater than first one\n")
// 	fmt.Scan(&b)
//     fmt.Printf("Second digit is %d\n", b)
// 	i := 1
// 	for range(b-a-1){
//         fmt.Println(a+i)
//         i++
// 	}
// }
//3.2
// func main(){
// 	var a, b int
//     fmt.Println("Enter first digit\n")
// 	fmt.Scan(&a)
//     fmt.Printf("First digit is %d\n", a)
//     fmt.Println("Enter second digit, greater than first one\n")
// 	fmt.Scan(&b)
//     fmt.Printf("Second digit is %d\n", b)
// 	i := 1
// 	for range(b-a-1){
//         if (a+i)%5 == 0{
//             fmt.Println(a+i)
//         }
//         i++
// 	}
// }
//3.3
// func main(){
// 	var a int
//     fmt.Println("Enter digit\n")
// 	fmt.Scan(&a)
//     i :=1
// 	for range(10){
//         fmt.Println(a*i)
//         i++
// 	}
// }
//3.4
// func main(){
//     var a string
//     var number, sum int
// fmt.Println("Enter your grades:\n")
//     fmt.Scan(&a)
//     for i:=0; i<utf8.RuneCountInString(a); i++ {
//         number, _ = strconv.Atoi(string(a[i]))
//         sum += number
//     }
//     fmt.Println(sum)
// }
//3.5
// func main(){
//     var sum, value int
//     for {
//         fmt.Println("Enter any digit but -1\n")
//         fmt.Scan(&value)
//         if value == -1{
//             break
//         }
//         sum+=value
//     }
//     fmt.Println(sum)
// }
//3.6
//Решил в 3.4, неверно понял условие
//3.7
