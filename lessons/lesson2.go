package lessons

import (
	"errors"
	"fmt"
)

func Lesson2() {
	// Операторы

	// арифметические
	fmt.Println("Арифметические операторы")
	// деление - возвращает целую часть
	var a = 5 / 2
	fmt.Println(a)
	// % - остаток от деления
	var b = (5 % 2)
	fmt.Println(b)
	// сложение, вычитание, умножение

	// операторы сравниения
	fmt.Println("Операторы сравнения")
	res := (5 == 6) // bool
	fmt.Println(res)

	res1 := (5 != 6) // bool
	fmt.Println(res1)

	res2 := (5 < 6) // bool
	fmt.Println(res2)

	res3 := (5 >= 6) // bool
	fmt.Println(res3)

	// логические операторы
	// && и
	// || или
	// ! не
	fmt.Println("Логические операторы")
	c := true
	d := false
	fmt.Println(c)
	fmt.Println(c && d)
	fmt.Println(c || d)
	fmt.Println(!c)

	// Управляющие конструкции, циклы, сложные типы данных, функции

	// Условные операторы - помогают управлять потоком программы
	// if - если. Ожидает true/false
	fmt.Println("Условные операторы")
	fmt.Println("if")

	f, g := 2, 3
	if !(f >= g) && g == 3 {
		fmt.Println("g = ", g)
	}

	newMap := map[string]string{}
	val, ok := newMap["a"]
	if ok {
		fmt.Println("val =", val)
	} else if g == 3 {
		fmt.Println("OK not found")
	} else {
		fmt.Println("OK")
	}

	fmt.Println("switch")
	key := "five"
	switch key {
	case "zero":
		fmt.Println("Not here!")
	case "five":
		fmt.Println("Here!")
	}

	// циклы
	fmt.Println("Циклы")
	animalSlice := []string{"cat", "dog", "catodog"}
	for i := 0; i < len(animalSlice); i++ { // index
		fmt.Println(animalSlice[i])
	}

	//for {
	//	fmt.Println("stop please")
	//	time.Sleep(1 * time.Second)
	//}

	for index, value := range animalSlice { // for each
		fmt.Println(index, value)
	}

	for _, value := range animalSlice { // for each, если индекс не используем
		fmt.Println(value)
	}

	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(numbers); i++ {
		if (numbers[i] % 2) == 0 {
			fmt.Println(numbers[i])
		}
	}

	// Функции и возврат значений
	fmt.Println("Функции")
	odd1(numbers)
	fmt.Println(summarize(5, 5))
	resik, _ := summarize(-1, 5)
	fmt.Println(resik)

	// Анонимная функция
	fmt.Println(func() int { return 1 }())

	perem := func() int { return 3 }()
	fmt.Println(perem)

	// Функция как аргумент другой функции
	tem(func() { fmt.Println("I'm inside") })

	// Именование как для переменных, так и для функций - camelCase
	// Название должно отражать суть
	userName := "ss"
	fmt.Println(userName)

	// Функция, которая реализует цикл с условиями и возвращает счетчик итераций
	numbers = []int{1, 2, 3, 100, 5}
	fmt.Println(iterationCounter(numbers))
	// Функция, возвращающая факториал числа
	fmt.Println(factorial(5))

	// Обработка ошибок и паник
	// Ошибка при обработке - все ок (мб обработаны и исправлены)
	// Паника - программа останавливается и падает (можем поймать и более мягко завершить программу)
	// Ошибка - тип данных
	fmt.Println("Ошибки и паники")
	i, err := divide(5, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
	//p, err := divide(5, 0)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(p)
	// Переменная с ошибкой - с большой буквы
	ty, err := divideWithNewError(5, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ty)

	// defer - вызов самым последним внутри, идут снизу - вверх, если их несколько
	fmt.Println("Defer")
	defer func() {
		fmt.Println("Hello from 2-nd defer")
	}()
	fmt.Println("Hello from before-defer")

	defer func() {
		fmt.Println("Hello from 1-st defer")
	}()
	fmt.Println("Hello from before-before-defer")

	// обработка паники с помощью defer
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Hello")
	panic("Word")
	fmt.Println("end")
}

// Объявление функции четности
func odd1(arr []int) {
	for i := 0; i < len(arr); i++ {
		if (arr[i] % 2) == 0 {
			fmt.Println(arr[i])
		}
	}
}

// Объявление функции суммирования
func summarize(x, y int) (int, error) {
	if x < 0 || y < 0 {
		return 0, fmt.Errorf("1 || 2 number < 0")
	}
	return x + y, nil
}

func tem(fufu func()) {
	fufu()
}

func iterationCounter(numbers []int) int {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == 100 {
			return i
		}
	}
	return 0
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Переменная с ошибкой - с большой буквы
var ErrInvalidInput = errors.New("invalid input")

func divideWithNewError(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrInvalidInput
	}
	return a / b, nil
}
