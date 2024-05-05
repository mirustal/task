package fmtpostfix

import "fmt"

/*
Задача 1. 

Разработайте функцию, которая принимает целое число в качестве аргумента и возвращает строку, 
содержащую это число и слово "компьютер" в нужном склонении по падежам в зависимости от числа. 

Например, при вводе числа 25 функция должна возвращать "25 компьютеров", 
для числа 41 — "41 компьютер", а для числа 1048 — "1048 компьютеров".
*/

func fmtPostfix(number int) string {
	absNumber := number
	if absNumber < 0 {
		absNumber = -absNumber 
	}

	lastDigits := absNumber % 100

	if lastDigits >= 11 && lastDigits <= 14 {
		return fmt.Sprintf("%d компьютеров", number)
	}

	lastDigit := absNumber % 10
	switch lastDigit {
	case 1:
		return fmt.Sprintf("%d компьютер", number)
	case 2, 3, 4:
		return fmt.Sprintf("%d компьютера", number)
	default:
		return fmt.Sprintf("%d компьютеров", number)
	}
}

