package lessons

import "fmt"

func Lesson1() {
	// Типы данных

	// целые
	fmt.Println("Инты")
	var value int = -34
	fmt.Println(value)

	// целые положительные
	fmt.Println("юИнты")
	var valueUint uint = 4
	fmt.Println(valueUint)

	// строка
	fmt.Println("Строки")
	var stringa string = "sss"
	fmt.Println(stringa)

	// числа с плавающей точкой
	fmt.Println("Флоты")
	var float float32 = 3.1415926
	fmt.Println(float)

	var float1 float64 = 3.1415926
	fmt.Println(float1)

	// булевые
	fmt.Println("Булевы")
	var boolValue bool = true
	fmt.Println(boolValue)

	// массивы
	fmt.Println("Массивы")
	var arr [10]int // изменить размер нельзя
	arr[0] = 10
	arr[9] = 20
	fmt.Println(arr)

	var arr1 [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr1)
	fmt.Println(len(arr1))
	fmt.Println(arr1[6])

	arr2 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(len(arr2))

	// слайсы (динамические массивы)
	fmt.Println("Слайсы")
	var slice []int // не указываем длину
	fmt.Printf("type: %T;\nvalues: %v;\nlen: %d;\ncap: %d;\n", slice, slice, len(slice), cap(slice))
	slice = append(slice, 39)
	fmt.Println(slice)
	slice = append(slice, 41, 29)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Printf("type: %T;\nvalues: %v;\nlen: %d;\ncap: %d;\n", slice, slice, len(slice), cap(slice))
	slice = append(slice, 1)
	fmt.Printf("type: %T;\nvalues: %v;\nlen: %d;\ncap: %d;\n", slice, slice, len(slice), cap(slice)) // длина меньше капасити, если длина превышает капасити за операцию - капасити увеличеивается в 2 раза
	// создание слайма через мейк - указывается размер, капасити
	slice2 := make([]int, 5, 10)
	fmt.Println(len(slice2), cap(slice2))
	fmt.Println(slice2) // 0 0 0 0 0

	// мапы
	fmt.Println("Мапы")
	var myMap map[string]int     // сразу заносить не сможем, надо проинициализировать
	myMap = make(map[string]int) // инициализация мапы
	myMap["key1"] = 123
	fmt.Println(len(myMap))
	myMap["key2"] = 456
	fmt.Println(myMap)
	fmt.Println(len(myMap))
	// ключи - почти любые типы данных, но не булевые
	myMap2 := make(map[int]string)
	myMap2[1] = "Tripod"
	fmt.Println(myMap2[1])
	// проверка ключа в мапке
	fmt.Println(myMap2[2])                 // пустая строка, хотя мы не добавляли ключ
	fmt.Println(myMap["key3"])             // 0, хотя мы не добавляли ключ
	valueFromEmptyKey, ok := myMap["key3"] // в ок - бул, есть ли ключ
	if !ok {
		fmt.Println("OK?", ok)
		fmt.Println("Пустое значение", valueFromEmptyKey)
	}

	// руны
	// Строки в Go представляют собой последовательность байтов,
	// в то время как руны представляют собой последовательность символов Unicode.
	// Каждая руна в Go — это 32 бита, что позволяет ей хранить все символы в кодировке Unicode.
	fmt.Println("Руны")
	var r rune = 'a'
	fmt.Println(r) // 97

	// комплексные числа
	fmt.Println("Комплексные числа")
	var f complex64 = 1 + 2i
	var g complex128 = 4 + 3i
	fmt.Println(f)
	fmt.Println(g)
}
