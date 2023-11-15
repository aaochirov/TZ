package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите выражение в формате \"2 + 2\"")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	//fmt.Println(input)

	operands := strings.Split(input, " ")
	if len(operands) != 3 {
		fmt.Println("Слишком много операндов")
	}
	num1Str := operands[0]
	num2Str := operands[2]
	operator := operands[1]

	i := 0 //Индикатор

	num1, err1 := strconv.Atoi(num1Str)
	if err1 != nil {
		num1 = RomanToArabic(num1Str)
		i++
	}
	num2, err2 := strconv.Atoi(num2Str)
	if err2 != nil {
		num2 = RomanToArabic(num2Str)
		i++
	}
	switch i {
	case 1:
		fmt.Println("Разные системы счисления!")
	case 2:
		fmt.Println(ArabicToRoman(calc(num1, operator, num2)))
	default:
		fmt.Println(calc(num1, operator, num2))
	}
}

func calc(num1 int, operator string, num2 int) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Деление на ноль!")
		}
		return num1 / num2
	default:
		return 0
	}
}
func RomanToArabic(r string) int { //Тут на входе стринг с римскими числами из ввода пользователем
	r = strings.ToUpper(r)
	mapRoman := map[string]int{"I": 1, "II": 2, "III": 3, "IIII": 4, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10} //Мапа, ключ=римская, значение=арабская

	if _, ok := mapRoman[r]; !ok {
		fmt.Println("Такого значения нет!")
		return 0
	}
	return mapRoman[r]
}
func ArabicToRoman(a int) string { // Аргумент 'а' здесь результат после calc
	romanArr := [...]string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	arabicArr := [...]int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	roman := ""
	for i := 0; i < len(romanArr); i++ {
		for a >= arabicArr[i] {
			roman += romanArr[i]
			a -= arabicArr[i]
		}
	} //Внешний цикл следит за тем, что бы мы не провалились за вместимость массива romanArr, иначе дальше будем ловить ошибку.
	// Внутренний цикл находит самую большую соответствующую римску цифру, лепит её справа и идёт дальше, пока не закончатся заполненные разряды
	return roman
}
