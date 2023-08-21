package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	oneNum, twoNum, operator := getInput()
	if getType(oneNum) != getType(twoNum) {
		fmt.Println("не поддерживаются операции между разными типами чисел")
		os.Exit(0)
	}
	numType := getType(oneNum)
	var oneNumInt, twoNumInt int64
	switch numType {
	case "dec":
		oneNumInt, _ = strconv.ParseInt(oneNum, 10, 64)
		twoNumInt, _ = strconv.ParseInt(twoNum, 10, 64)
	case "rim":
		oneNumInt = DecodeRim(oneNum)
		twoNumInt = DecodeRim(twoNum)
	}
	var result int64
	switch operator {
	case "+":
		result = oneNumInt + twoNumInt
	case "-":
		result = oneNumInt - twoNumInt
	case "*":
		result = oneNumInt * twoNumInt
	case "/":
		result = int64(oneNumInt / twoNumInt)
	}
	switch numType {
	case "dec":
		fmt.Println("Ответ:")
		fmt.Println(result)
	case "rim":
		if result < 1 {
			fmt.Println("Ответ не поддерживается римской системой счисления")
			os.Exit(0)
		}
		fmt.Println("Ответ:")
		fmt.Println(EncodeRim(result))
	}
}

func getInput() (string, string, string) { //Запрашивает информацию от пользователя и форматирует для удобства
	// получаем строку и избавляемся от всех пробелов.
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите пример:")
	userInput, _ := reader.ReadString('\n')
	userInput = strings.ReplaceAll(userInput, " ", "")

	// Поиск оператора и проверка единственный ли он.
	operators := "+-*/"
	var cursor int
	var i int
	notCursor := true
	for i = 0; i < len(userInput); i++ {
		if strings.Contains(operators, string(userInput[i])) {
			if notCursor {
				notCursor = false
				cursor = i
			} else {
				fmt.Println("Поддерживаются только выражения с одним оператором")
				os.Exit(0)
			}
		}
	}
	if notCursor {
		fmt.Println("Оператор не найден")
		os.Exit(0)
	}
	// пререведем строку в 3 отдельные переменные и вернем их.
	oneNum := string(userInput[0:cursor])
	operator := string(userInput[cursor])
	twoNum := strings.TrimSpace(string(userInput[cursor+1 : len(userInput)-1]))
	return oneNum, twoNum, operator
}
func getType(num string) string {
	alph := ""
	var result string
	if strings.Contains("IVXLCDM", string(num[0])) {
		alph = "IVXLCDM"
		result = "rim"
	} else if strings.Contains("1234567890", string(num[0])) {
		alph = "1234567890"
		result = "dec"
	} else {
		fmt.Println("Число не распознано")
		os.Exit(0)
	}
	var i int
	for i = 1; i < len(num); i++ {
		if !strings.Contains(alph, string(num[i])) {
			fmt.Println("Число не распознано")
			os.Exit(0)
		}
	}
	return result
}

func DecodeRim(roman string) int64 {
	var sum int64
	var Roman = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	for k, v := range roman {
		if k < len(roman)-1 && Roman[byte(roman[k+1])] > Roman[byte(roman[k])] {
			sum -= int64(Roman[byte(v)])
		} else {
			sum += int64(Roman[byte(v)])
		}
	}
	return sum
}

func EncodeRim(number int64) string {
	conversions := []struct {
		value int64
		digit string
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

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}
