/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100)
2) Количество (только числа)
3) ФИО покупателя (только буквы)
4) Контактный телефон (10 цифр)
5) Адрес: индекс(ровно 6 цифр), город, улица, дом, квартира
Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.
Реализовать конструктор и несколько методов у типа "Накладная"
Пример:
invoice = NewInvoice()
или
order = NewOrder()
*/
package main
import(
	"fmt"
)

func (s Staff) setName (newName string) {
	s.name = newName
}

func (b Buyer) setName (newName string) {
	b.name = newName
}

func (s Staff) setQuantity (newQuantity int) {
	s.quantity = newQuantity
}

func (b Buyer) setSurname (newSurname string) {
	b.surname = newSurname
}

func (b Buyer) setLastname (newLastname string) {
	b.lastName = newLastname
}

func (a Address) setIdx (newIdx int) {
	a.idx = newIdx
}

func (a Address) setCity (newCity string) {
	a.city = newCity
}

func (a Address) setStreet (newStreet string) {
	a.street = newStreet
}

func (a Address) setHouse (newHouse int) {
	a.house = newHouse
}

func (a Address) setFlat (newFlat int) {
	a.idx = newFlat
}

type Address struct {
	idx int
	city, street string
	house, flat int
}

type Staff struct {
	name string
	quantity int
}

type Buyer struct {
	name, surname, lastName string
	address Address
}

type Invoice struct {
	staff Staff
	buyer Buyer
}






func main() {

invoice := NewInvoice{}











}

