package main

//first

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите значение")
		text, _ := reader.ReadString('\n') // Ждет ввода данных в формате строки
		text = strings.TrimSpace(text)     // Очищает все пустоты (пробелы, табуляцию)
		toNumber, _ := strconv.Atoi(text)  // Преобразование строки в число
		fmt.Println(toNumber + 8)          // Вывод результата

		arrStrig := strings.Split(text, " ")
		chekNumber(arrStrig)
	}
}

func chekNumber(arr []string) {
	if len(arr) != 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	} else {
		if contains(arr[0]) {
			if contains(arr[2]) {
				calcRim(arr)
				return
			}
			panic("Выдача паники, так как используются одновременно разные системы счисления.")
		} else {
			calcNumb(arr)
		}

		fmt.Println(arr[0])
		fmt.Println(arr[1])
		fmt.Println(arr[2])
	}
}

func calcRim(arr []string) {
	var num1 = rimToInt(arr[0])
	var num2 = rimToInt(arr[2])
	var res int = 0
	if num1 == -1 {
		panic("Выдача паники, так как строка не является математической операцией.")
	}
	if num2 == -1 {
		panic("Выдача паники, так как строка не является математической операцией.")
	}
	switch arr[1] {
	case "+":
		res = num1 + num2
		fmt.Println(intToRim(res))
	case "-":
		res = num1 - num2
		if res < 1 {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		} else {
			fmt.Println(intToRim(res))
		}
	case "*":
		res = num1 * num2
		fmt.Println(intToRim(res))
	case "/":
		res = num1 / num2
		fmt.Println(intToRim(res))
	default:
		panic("Выдача паники, так как строка не является математической операцией.")
	}

}

func rimToInt(item string) int {
	a := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, n := range a {
		if item == n {
			return num[i]
		}
	}
	return -1
}

func intToRim(item int) string {
	if item > 10 {
		panic("Выдача паники, так как число больше 10.")

	} else {
		a := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
		num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for i, n := range num {
			if item == n {
				return a[i]
			}
		}
	}
	return ""
}

func calcNumb(arr []string) {
	toNumber1, err := strconv.Atoi(arr[0])
	if err != nil {
		panic("Выдача паники, так как строка не является математической операцией.")
	}
	toNumber2, err := strconv.Atoi(arr[2])
	if err != nil {
		panic("Выдача паники, так как строка не является математической операцией.")
	}
	switch arr[1] {
	case "+":
		fmt.Println(toNumber1 + toNumber2)
	case "-":
		fmt.Println(toNumber1 - toNumber2)
	case "*":
		fmt.Println(toNumber1 * toNumber2)
	case "/":
		fmt.Println(toNumber1 / toNumber2)
	default:
		panic("Выдача паники, так как строка не является математической операцией.")
	}
}

func contains(item string) bool {
	a := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, n := range a {
		if item == n {
			return true
		}
	}
	return false
}
