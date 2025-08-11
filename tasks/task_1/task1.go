package task_1

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Task1() {
	names := getAndWriteNames()
	symbol := getAndWriteSymbol()
	count := 0

	for _, name := range names {
		if strings.HasPrefix(name, symbol) {
			count++
			upperFirstName := capitalizeFirst(name)
			fmt.Println(upperFirstName)
		}
	}
	if count == 0 {
		fmt.Println("Ничего не найдено")
	}
}

func getAndWriteNames() []string {
	fmt.Print("Введите имена через пробел: ")

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода имен, попробуйте снова.")
		return getAndWriteNames()
	}

	// Убираем лишние пробелы и перевод строки
	line = strings.TrimSpace(line)

	lineInLowercase := strings.ToLower(line)

	// Разделяем строку по пробелам
	names := strings.Fields(lineInLowercase)

	return names
}

func getAndWriteSymbol() string {
	fmt.Print("Введите первую букву имени: ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода буквы, попробуйте снова. Введи 1 букву")
		return getAndWriteSymbol()
	}

	line = strings.TrimSpace(line)

	lineInLowercase := strings.ToLower(line)

	symbol := strings.Fields(lineInLowercase)[0]

	return symbol
}

func capitalizeFirst(s string) string {
	if s == "" {
		return s
	}

	// Превращаем строку в срез рун, чтобы корректно работать с Unicode
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}
