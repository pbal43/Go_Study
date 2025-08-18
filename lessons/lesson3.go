package lessons

import (
	"fmt"
	"reflect"
)

func Lesson3() {
	// ООП в Go - на страктурах и интерфейсах. Методы структур. Вложенные структуры
	fmt.Println("Создание структуры")
	// Инициализация экземпляра
	myCar := Car{
		brand:  "Honda",
		color:  "Red",
		engine: 1.6,
		age:    22,
	}

	var myCar2 Car
	myCar2.brand = "Toyota"
	myCar2.color = "Black"
	myCar2.engine = 2.6
	myCar2.age = 2

	fmt.Println(myCar)
	fmt.Println(myCar2)
	fmt.Printf("First car color: %s\n", myCar.color)

	// Использование методов объекта
	fmt.Println("Методы структуры")
	myCar.startEngine()
	myCar2.startEngine()

	// Структурка с вложенной структуркой
	fmt.Println("Вложенныые структуры")
	newCar := Machine{"Civic", Engine{3.2, 202}}
	fmt.Printf("New car: %v", newCar)

	// Интерфейсы - описание метолов, который должна реализовывать структура
	fmt.Println("Интерфейсы")

	Kawa := Motorcycle{"Kawasaki", "Green"}

	Scoopeedon := Scooter{"Scoopy", "Blue"}

	// Функция для интерфейса, где в качестве аргумента - структура, реализующая интерфейс
	washVehicle(Kawa)

	washVehicle(Scoopeedon)

	// Пакеты и модули
	fmt.Println("Пакеты и модули")
	// cmd -> stend - main.go - точка входа в программу stend
	// internal - код скрыт для получения из любой другой папки, не стянуть. Методы отсюда в другом проекте не получить
	// в internal - основная бизнес-логика, разбитая на модули (сервер, сторадж, domain для описания моделек и проч)
	// domain в Go — это про организацию кода: в ней хранят сущности, бизнес-правила и интерфейсы, то есть то, что описывает предметную область, без привязки к инфраструктуре.
	// pkg - хранение реализации для того, чтобы делиться (типа либы)
	// докер-файлы - обычно в корень
	// циклическая зависимость - при импорте чего-либо из разных пакетов друг в друга одновременно. Так нельзя

	// Указатели
	fmt.Println("Указатели")
	str := "string"
	fmt.Println(str)
	fixString(str)   // передали копию аргумента (передача по значению, by value)
	fmt.Println(str) // строка не изменилась
	// чтобы так сделать - надо передавать по ссылке, by reference
	linkStr := &str      // через амперсанд - адрес в памяти
	fmt.Println(linkStr) // выведется адрес переменной в памяти
	newFixedString(&str) // передаем указатель или адрес
	fmt.Println(str)     // значение str переменной изменилось
	// & - взять адрес (типа метод, при котором мы получаем указатель, это еще не указатель, указатель - результат доставания
	// * - разыменование, мы получаем значение по адресу
	fmt.Println("указатель имеет тип:", reflect.TypeOf(&str)) // *string
	// применение в структурах
	fmt.Println("Указатели в структурах")
	bAcc := BankAccount{"Sergey", 100}
	bAcc.withdrawal(50)
	bAcc.deposit(10)
	fmt.Println(bAcc.getBalance()) // баланс не изменился

	bAcc.newWithdrawal(101)
	bAcc.newDeposit(10)
	fmt.Println(bAcc.getBalance()) // баланс изменился

}

// Объявление структуры
type Car struct {
	brand  string
	color  string
	engine float32
	age    int
}

// Метод структуры (привязка через аргументы)
func (c Car) startEngine() {
	fmt.Printf("%s is making patata patata\n", c.brand)
}

type Machine struct {
	name   string
	engine Engine
}

type Engine struct {
	capacity float32
	power    int
}

type EngineStart interface {
	StartEng(int)
	StopEng()
}

func (e Engine) StartEng(chanceCount int) {
	fmt.Printf("Завелась с %s раз\n", chanceCount)
}

func (e Engine) StopEng() {
	fmt.Printf("Двигатель выключен")
}

// Если структура реализует все методы интерфейса (назввание, аргументы), то она реализует интерфейс
type Vehicle interface {
	StartVehicle(string)
	StopVehicle()
}

type Motorcycle struct {
	brand string
	color string
}

type Scooter struct {
	brand string
	color string
}

// Реализация метода интерфейса
func (m Motorcycle) StartVehicle(sound string) {
	fmt.Printf("%s is making \n", sound)
}

// Реализация метода интерфейса
func (m Motorcycle) StopVehicle() {}

func (s Scooter) StartVehicle(sound string) {
	fmt.Printf("%s is making \n", sound)
}
func (s Scooter) StopVehicle() {}

// Функция для интерфейса, где в качестве аргумента - структура, реализующая интерфейс
func washVehicle(v Vehicle) {
	v.StopVehicle()
	v.StartVehicle("Wrooom Wrooootatam")
}

// указатели
// без него
func fixString(s string) {
	s = "fixed string"
}

// с ним
func newFixedString(s *string) { // ожидаем указатель (АДРЕС)
	fmt.Println(s)      // адрес
	*s = "fixed string" // звездочка - разыменование указателя, по адресу идем к значению (УЖЕ НЕ АДРЕС, А ЗНАЧЕНИЕ ПО НЕМУ)
}

type BankAccount struct {
	name    string
	balance int
}

func (b BankAccount) withdrawal(amount int) {
	b.balance -= amount
}

func (b BankAccount) deposit(amount int) {
	b.balance += amount
}

func (b BankAccount) getBalance() int {
	return b.balance
}

func (b *BankAccount) newWithdrawal(amount int) {
	if b.balance < amount {
		fmt.Println("Недостаточно средств")
		return
	}
	b.balance -= amount // не передаем зведочку, так как go использует сахар и сам под капотом со структурами это делает
	// ТОЛЬКО ДЛЯ СТРУКТУР (ТАК КАК МБ МНОГО ПОЛЕЙ)
}

func (b *BankAccount) newDeposit(amount int) {
	b.balance += amount
}
