package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculate(input string) (string, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return "", fmt.Errorf("Неверный формат ввода")
	}

	operand1 := parts[0]
	operator := parts[1]
	operand2 := parts[2]

	isArabic := isArabicNumber(operand1) && isArabicNumber(operand2)
	isRoman := isRomanNumber(operand1) && isRomanNumber(operand2)

	if isArabic && isRoman {
		return "", fmt.Errorf("Используются одновременно разные системы счисления")
	}

	var result int
	var isRomanInput bool
	if isArabic {
		num1, num2 := toArabic(operand1), toArabic(operand2)
		if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
			return "", fmt.Errorf("Числа должны быть от 1 до 10 включительно")
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
				return "", fmt.Errorf("Деление на ноль")
			}
			result = num1 / num2
		default:
			return "", fmt.Errorf("Неподдерживаемая операция: " + operator)
		}
	} else if isRoman {
		num1, num2 := romanToArabic(operand1), romanToArabic(operand2)
		if num1 <= 0 || num1 > 10 || num2 <= 0 || num2 > 10 {
			return "", fmt.Errorf("Римские числа должны быть от I до X включительно")
		}

		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			result = num1 - num2
			if result <= 0 {
				return "", fmt.Errorf("Римские числа не поддерживают отрицательные значения")
			}
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				return "", fmt.Errorf("Деление на ноль")
			}
			result = num1 / num2
		default:
			return "", fmt.Errorf("Неподдерживаемая операция: " + operator)
		}

		isRomanInput = true
	} else {
		return "", fmt.Errorf("Операнды не являются числами")
	}

	if isRomanInput {
		romanResult, err := toRoman(result)
		if err != nil {
			return "", err
		}
		return romanResult, nil
	}

	return strconv.Itoa(result), nil
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

func toRoman(num int) (string, error) {
	if num <= 0 {
		return "", fmt.Errorf("Число должно быть положительным")
	}

	romanNumerals := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}

	if roman, exists := romanNumerals[num]; exists {
		return roman, nil
	}

	standardRomanNumerals := []struct {
		Value int
		Roman string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""
	for _, pair := range standardRomanNumerals {
		for num >= pair.Value {
			result += pair.Roman
			num -= pair.Value
		}
	}

	return result, nil
}
