package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculate(input string) (int, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return 0, fmt.Errorf("Неверный формат ввода")
	}

	operand1 := parts[0]
	operator := parts[1]
	operand2 := parts[2]

	isArabic := isArabicNumber(operand1) && isArabicNumber(operand2)
	isRoman := isRomanNumber(operand1) && isRomanNumber(operand2)

	if isArabic && isRoman {
		return 0, fmt.Errorf("Используются разные системы исчисления")
	}

	var result int

	if isArabic {
		num1, num2 := toArabic(operand1), toArabic(operand2)
		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				return 0, fmt.Errorf("Деление на ноль")
			}
			result = num1 / num2
		default:
			return 0, fmt.Errorf("Неподдерживаемая операция" + operator)
		}
	} else if isRoman {
		num1, num2 := romanToArabic(operand1), romanToArabic(operand2)
		if num1 <= 0 || num2 <= 0 {
			return 0, fmt.Errorf("Римские числа могут быть только положительными")
		}
		if num1 < num2 {
			return 0, fmt.Errorf("Римские числа не поддерживают отрицательные значения")
		}
		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				return 0, fmt.Errorf("Деление на ноль")
			}
			result = num1 / num2
		default:
			return 0, fmt.Errorf("Неподдерживаемая операция" + operator)
		}
	} else {
		return 0, fmt.Errorf("Операнды не являются числами")
	}
	return result, nil
}

func isArabicNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func toArabic(s string) int {
	num, _ := strconv.Atoi(s)
	return num
}
func isRomanNumber(roman string) bool {
	return romanToArabic(roman) > 0
}

func romanToArabic(roman string) int {
	romanValues := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	result := 0

	if value, exists := romanValues[roman]; exists {
		result = value
	} else {
		return 0
	}
	return result
}
