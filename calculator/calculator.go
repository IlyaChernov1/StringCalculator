package calculator

import (
	"strconv"
	"strings"
)

// Сложение строк
func add(x, y string) string {
	return strings.Trim(x, "\"") + strings.Trim(y, "\"")
}

// Вычитание подстроки из строки
func subtract(x, y string) string {
	trimmedX := strings.Trim(x, "\"")
	trimmedY := strings.Trim(y, "\"")

	if !strings.Contains(trimmedX, trimmedY) {
		return trimmedX
	}
	return strings.ReplaceAll(trimmedX, trimmedY, "")
}

// Умножение строки на число
func multiply(x string, y string) string {
	number, err := strconv.Atoi(y)
	if err != nil || number < 1 || number > 10 {
		panic("Число должно быть от 1 до 10")
	}
	return strings.Repeat(strings.Trim(x, "\""), number)
}

// Деление строки на число
func divide(x string, y string) string {
	number, err := strconv.Atoi(y)
	if err != nil || number < 1 || number > 10 {
		panic("Число должно быть от 1 до 10 включительно")
	}

	trimmedX := strings.Trim(x, "\"")
	length := len(trimmedX)
	newLength := length / number

	if newLength == 0 {
		return ""
	}

	return trimmedX[:newLength]
}

// Если строка больше 40 символов, обрезается и добавляется "..."
func truncate(s string) string {
	if len(s) > 40 {
		return s[:40] + "..."
	}
	return s
}

// Функция для проверки наличия строки в кавычках
func isQuoted(s string) bool {
	return strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"")
}

// Функция для проверки, находится ли символ за пределами кавычек
func isOperatorOutsideQuotes(expression string, position int) bool {
	quoteCount := 0
	for i, char := range expression {
		if char == '"' {
			quoteCount++
		}
		if i == position && quoteCount%2 == 0 {
			return true
		}
	}
	return false
}

// Функция для выполнения операций
func Calculate(str string) string {
	var x, y string

	str = strings.TrimSpace(str)

	// Проверяем операцию сложения
	for i, char := range str {
		if char == '+' && isOperatorOutsideQuotes(str, i) {
			parts := strings.SplitN(str, "+", 2)
			x = strings.TrimSpace(parts[0])
			y = strings.TrimSpace(parts[1])
			if !isQuoted(x) {
				panic("Первым аргументом должна быть строка в кавычках")
			}
			if isQuoted(x) && isQuoted(y) {
				return "\"" + truncate(add(x, y)) + "\""
			} else {
				return "Ошибка: обе строки должны быть в кавычках"
			}
		}
	}

	// Проверяем операцию вычитания
	for i, char := range str {
		if char == '-' && isOperatorOutsideQuotes(str, i) {
			parts := strings.SplitN(str, "-", 2)
			x = strings.TrimSpace(parts[0])
			y = strings.TrimSpace(parts[1])
			if !isQuoted(x) {
				panic("Первым аргументом должна быть строка в кавычках")
			}
			if isQuoted(x) && isQuoted(y) {
				return "\"" + truncate(subtract(x, y)) + "\""
			} else {
				return "Ошибка: обе строки должны быть в кавычках"
			}
		}
	}

	// Проверяем операцию умножения
	for i, char := range str {
		if char == '*' && isOperatorOutsideQuotes(str, i) {
			parts := strings.SplitN(str, "*", 2)
			x = strings.TrimSpace(parts[0])
			y = strings.TrimSpace(parts[1])
			if !isQuoted(x) {
				panic("Первым аргументом должна быть строка в кавычках")
			}
			if isQuoted(x) && !isQuoted(y) {
				return "\"" + truncate(multiply(x, y)) + "\""
			} else {
				return "Ошибка: первая строка должна быть в кавычках, второй операнд должен быть числом"
			}
		}
	}

	// Проверяем операцию деления
	for i, char := range str {
		if char == '/' && isOperatorOutsideQuotes(str, i) {
			parts := strings.SplitN(str, "/", 2)
			x = strings.TrimSpace(parts[0])
			y = strings.TrimSpace(parts[1])
			if !isQuoted(x) {
				panic("Первым аргументом должна быть строка в кавычках")
			}
			if isQuoted(x) && !isQuoted(y) {
				return "\"" + truncate(divide(x, y)) + "\""
			} else {
				return "Ошибка: первая строка должна быть в кавычках, второй операнд должен быть числом"
			}
		}
	}

	return "Неизвестная операция"
}
