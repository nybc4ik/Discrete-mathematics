/*
Важная информация:
Коды цифр и символов
0-48
1-49
2-50
3-51
4-52
5-53
6-54
7-55
8-56
9-57
( - 40
) - 41
* - 42
+ - 43
"-" - 45
*/
package main

import "fmt"

//Функция проверки является ли код символа числовым кодом
func number(a byte) bool {
	if (a >= 48) && (a <= 57) {
		return true
	}
	return false
}

//Функция проверки символов *+-
func step(a byte) bool {
	if a == 42 {
		return true
	}
	if a == 43 {
		return true
	}
	if a == 45 {
		return true
	}
	return false
}

//Функция перевода символов в цифры
func string_to_number(a byte) int {
	return (int(a) - 48)
}

//Функция перехода к следующему символу
func next(Symbol string) string {
	var p, u, l int
	p = 0
	u = 0
	l = len(Symbol)
	for i := 0; i < l; i++ {
		if number(Symbol[i]) {
			p++
		} else if step(Symbol[i]) {
			p = p - 2
		} else if Symbol[i] == 40 {
			u++
		} else if Symbol[i] == 41 {
			u--
			p++
		}
		if (u == 0) && (p == 1) {
			return Symbol[i+1:]
		}
	}
	return " "
}

//Функия поиска символов в строке
func searcher(Symbol string) int {
	var a byte
	a = Symbol[0]
	Symbol = Symbol[1:]
	if a == 40 {
		return searcher(Symbol)
	} else if a == 42 {
		return (searcher(Symbol) * searcher(next(Symbol)))
	} else if a == 43 {
		return (searcher(Symbol) + searcher(next(Symbol)))
	} else if a == 45 {
		return (searcher(Symbol) - searcher(next(Symbol)))
	} else if number(a) {
		return string_to_number(a)
	}
	return 0
}

func main() {
	/*
		//определение кодов символов
		var n string
		var num byte
		fmt.Scan(&n)
		num = n[0]
		fmt.Print(num)
	*/
	var Symbol, Symbol1 string //строка(Symbol1) и строка побольше(Symbol)
	for i := 0; i < 2; i++ {
		fmt.Scan(&Symbol1)
		if Symbol1 == "" {
			i = 3
		} else {
			i = -1
			Symbol = Symbol + Symbol1
			Symbol1 = ""
		}
	}
	fmt.Println(searcher(Symbol))
}