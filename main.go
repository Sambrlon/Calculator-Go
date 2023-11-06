package main

import (
	"bufio"
	"calculator1/calculator"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите выражение: ")
	scanner.Scan()
	input := scanner.Text()

	result, err := calculator.Calculate(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(result)
	}
}
